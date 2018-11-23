

default: build


build:
    go build -ldflags "-s -w" -o bioweb cmd/main.go
