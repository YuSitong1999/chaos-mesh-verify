package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func workBio() {
	log.Printf("bio begin: %s", time.Now().Format(time.RFC3339Nano))
	defer func() {
		log.Printf("bio end: %s", time.Now().Format(time.RFC3339Nano))
	}()
	file, err := os.OpenFile("/data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("open file error: %v", err)
		return
	}
	for i := 0; i < 1e4; i++ {
		_, err = file.WriteString(fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Printf("write file error: %v", err)
			return
		}
		err = file.Sync()
		if err != nil {
			log.Printf("sync file error: %v", err)
			return
		}
	}
}

func main() {
	for {
		workBio()
		log.Printf("sleep 10s")
		time.Sleep(10 * time.Second)
		log.Printf("sleep end")
	}
}
