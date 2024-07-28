package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func generateOffsets(filename string, done <-chan int, workers int) []chan int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return nil
	}
	offsets := make([]chan int64, workers)

	for i := range offsets {
		offsets[i] = make(chan int64)
	}

	go func() {
		defer func() {
			for i := range offsets {
				close(offsets[i])
			}
		}()

		for i := int64(0); i < fileInfo.Size()/bufferSize; i++ {
			select {
			case <-done:
				return
			default:
				offsets[i%int64(workers)] <- i * bufferSize
			}
		}
	}()

	return offsets
}

func launchWorker(pattern, filename string, offset <-chan int64, done <-chan int, lps []int64) <-chan []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil
	}

	finds := make(chan []int64)

	go func() {
		defer close(finds)
		content := make([]byte, bufferSize)
		var b int
		var off int64
		var ok bool
		var err error

		for {
			select {
			case <-done:
				return
			case off, ok = <-offset:
				if !ok {
					return
				}
				_, err = file.Seek(off, 0)
				if err != nil {
					log.Println(err)
					return
				}
				b, err = file.Read(content)
				if err != nil {
					log.Println(err)
					return
				}
				finds <- kmp(pattern, content[:b], lps)
			}
		}
	}()

	return finds
}

func divideWorkload(pattern, filename string, offsets []chan int64, done <-chan int, workers int) []<-chan []int64 {

	streams := make([]<-chan []int64, workers)
	lps := preprocess(pattern)

	for i := range workers {
		streams[i] = launchWorker(pattern, filename, offsets[i], done, lps)
	}

	return streams
}

func fanIn(inputChannels []<-chan []int64, done <-chan int) <-chan int64 {
	findings := make(chan int64)

	go func() {
		defer close(findings)
		index := int64(0)
		var indices []int64
		var ok bool
		for {
			for c := range inputChannels {
				select {
				case <-done:
					return
				default:
					indices, ok = <-inputChannels[c]
					if !ok {
						return
					}
					for i := range indices {
						findings <- index + indices[i]
					}
					index += bufferSize
				}
			}
		}
	}()

	return findings
}

func ParaSearch(pattern, filename string, done <-chan int) {
	workers := runtime.NumCPU()
	offsets := generateOffsets(filename, done, workers)
	outputs := divideWorkload(pattern, filename, offsets, done, workers)
	output := fanIn(outputs, done)

	for i := range output {
		fmt.Println(i)
	}
}
