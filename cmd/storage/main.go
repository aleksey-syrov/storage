package main

import "fmt"
import "github.com/aleksey-syrov/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	fmt.Println("It works", st)
}
