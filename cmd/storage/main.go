package main

import (
	"fmt"
	"log"
)
import "github.com/aleksey-syrov/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded", file)
}
