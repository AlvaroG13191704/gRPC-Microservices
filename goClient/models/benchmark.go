package models

type Benchmark struct {
	Http Http
	Grpc Grpc
}

type Http struct {
	Time   []float64
	Weight []float64
}

type Grpc struct {
	Time   []float64
	Weight []float64
}
