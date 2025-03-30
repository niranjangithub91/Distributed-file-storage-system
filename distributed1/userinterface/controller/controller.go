package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
	"userinterface/helper"
	"userinterface/model"
	"userinterface/userpb"

	"github.com/gorilla/mux"
	"github.com/klauspost/reedsolomon"
	"google.golang.org/grpc"
)

func getFileSize(file multipart.File) (int64, error) {
	// Check if the file supports seeking
	if seeker, ok := file.(io.Seeker); ok {
		currentOffset, _ := seeker.Seek(0, io.SeekCurrent) // Save current position
		size, err := seeker.Seek(0, io.SeekEnd)            // Move to end and get size
		seeker.Seek(currentOffset, io.SeekStart)           // Restore position
		return size, err
	}
	return 0, fmt.Errorf("file does not support seeking")
}
func Upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*1024) // Limit request size to 1GB

	fmt.Println("Sending data")
	err := r.ParseMultipartForm(10 << 20) // 10MB max memory
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Retrieve file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get file size
	fileSize, err := getFileSize(file)
	if err != nil {
		http.Error(w, "Error getting file size", http.StatusInternalServerError)
		return
	}

	// Ensure at least 3 chunks
	if fileSize < 3 {
		http.Error(w, "File too small to split into 3 chunks", http.StatusBadRequest)
		return
	}

	chunkSize := fileSize / 3 // Divide into 3 chunks
	remainder := fileSize % 3 // Handle last chunk size

	data := make([][]byte, 5)
	count := 0
	encoder, err := reedsolomon.New(3, 2)
	if err != nil {
		log.Fatal(err)
	}
	for partNum := 0; partNum < 3; partNum++ {
		fmt.Println("Entered loop")
		// Adjust the last chunk to include any remainder bytes
		actualChunkSize := chunkSize
		if partNum == 2 {
			actualChunkSize += remainder
		}

		// Read data into buffer
		buffer := make([]byte, actualChunkSize)
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		chunk := buffer[:bytesRead]
		data[count] = make([]byte, chunkSize)
		copy(data[count], chunk)
		count++
	}
	for j := 3; j < 5; j++ {
		data[j] = make([]byte, chunkSize)

	}
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal("Encoding error:", err)
	}

	partNum := 0

	var y []model.Chunk

	for i := 3001; i < 3006; i++ {

		finalDial := fmt.Sprintf("serve%d:%d", i-3000, i)

		conn, err := grpc.Dial(finalDial, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		client := userpb.NewSendDataClient(conn)
		fmt.Println(finalDial)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		final_name := fmt.Sprintf("chunk%d%s", partNum, handler.Filename)
		req := &userpb.DataSend{
			Data: data[partNum],
			Save: final_name,
		}
		res, err := client.Send(ctx, req)
		if err != nil {
			log.Fatalf("Error calling send: %v", err)
		}

		if !res.Status {
			log.Fatal(http.StatusBadRequest)
		}
		partNum++
		var y1 model.Chunk
		y1.Chunk_name = final_name
		y1.Number = int64(partNum)
		y = append(y, y1)
	}
	var y2 model.Meta
	y2.Name = handler.Filename
	y2.Chunk_size = chunkSize
	y2.Chunk_detail = y
	helper.Insert_data(y2)
}

func Download(w http.ResponseWriter, r *http.Request) {
	encoder, err := reedsolomon.New(3, 2)
	if err != nil {
		log.Fatal(err)
	}
	params := mux.Vars(r)
	t := params["name"]
	z, x := helper.Get_data(t)
	if !x {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := make([][]byte, 5)
	b := z.Chunk_detail
	count := 3001
	size := 5
	arr := make([]bool, size) // Creates a slice of size 10 with all elements set to false
	for i := 0; i < 5; i++ {
		finalDial := fmt.Sprintf("serve%d:%d", i+1, i+count)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, finalDial,
			grpc.WithInsecure(),
			grpc.WithBlock(), // Ensures it waits and fails if the server is off
		)
		if err != nil {
			fmt.Println("The server is down")
			data[i] = nil
			arr[i] = false
			continue

		}
		fmt.Println("Aiyo")
		defer conn.Close()
		client := userpb.NewSendDataClient(conn)
		fmt.Println(finalDial)
		req := &userpb.GetDataSend{
			FileName: b[i].Chunk_name,
		}
		res, err := client.Get(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		l := res.Data
		data[i] = l
		arr[i] = true
		if !res.Status {
			fmt.Println("error is coming")
		}
	}
	err = encoder.Reconstruct(data)
	if err != nil {
		log.Fatal("Reconstruction error:", err)
	}
	var display []byte
	for i := 0; i < 3; i++ {
		display = append(display, data[i]...)
	}

	w.Header().Set("Content-Type", "text/plain") // Ensure Postman displays it as text
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(display)))
	_, err1 := w.Write(display)
	if err1 != nil {
		log.Fatal(err)
	}

}
