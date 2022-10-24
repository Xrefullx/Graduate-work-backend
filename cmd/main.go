package main

import (
	"GraduateWork"
	"log"
)

func main() {
	srv := new(GraduateWork.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
