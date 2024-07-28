package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func searchPatternInFile(pattern, filename string) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	size := fileInfo.Size()
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	lps := preprocess(pattern)
	var fileOffset int64 = 0

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("could not read file: %v", err)
			}
			break
		}
		if bytesRead == 0 {
			break
		}

		indices := kmp(pattern, buffer[:bytesRead], lps)
		for _, index := range indices {
			fmt.Println(fileOffset + index)
		}

		fileOffset += int64(bytesRead)

		if fileOffset >= size {
			break
		}

		overlapSize := int64(len(pattern)) - 1
		if fileOffset-overlapSize >= 0 {
			fileOffset -= overlapSize
			if _, err := file.Seek(fileOffset, os.SEEK_SET); err != nil {
				log.Fatalf("could not seek in file: %v", err)
			}
		}
	}
}
