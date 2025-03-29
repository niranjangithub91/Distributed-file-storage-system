package model

type Meta struct {
	Name         string  `json:"name"`
	Chunk_size   int64   `json:"chunk_size"`
	Chunk_detail []Chunk `json:"chunk_details"`
}

type Chunk struct {
	Number     int64  `json:"number"`
	Chunk_name string `json:"chunk_name"`
}
