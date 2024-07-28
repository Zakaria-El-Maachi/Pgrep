package main

import (
	"fmt"
	"os"
	"strconv"
)

var bufferSize int64 = 16384

//

func main() {

	done := make(chan int)
	defer close(done)

	k := 0

	if len(os.Args) >= 2 {
		if os.Args[1] == "-p" {
			k++
		}
	}

	if len(os.Args) < 3+k {
		fmt.Println("Please Specify both the Pattern and the Filename")
		return
	}

	if len(os.Args) == 4+k {
		temp, err := strconv.Atoi(os.Args[3+k])
		if err != nil {
			fmt.Println(err)
			return
		}
		bufferSize = int64(temp)
	}

	pattern, filename := os.Args[1+k], os.Args[2+k]

	if k == 0 {
		searchPatternInFile(filename, pattern)
	} else {
		ParaSearch(pattern, filename, done)
	}

}
