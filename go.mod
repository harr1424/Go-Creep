module github.com/harr1424/Go-Creep

go 1.21.1

require github.com/rs/cors v1.11.0

require (
	github.com/gorilla/mux v1.8.1 // indirect
	golang.org/x/time v0.5.0 // indirect
)

require github.com/harr1424/Go-Creep/gocreep v0.0.0

replace github.com/harr1424/Go-Creep/gocreep => ./gocreep
