package main

import (
	"fmt"

	"io"
	"log"
	"math"
	"os"
)

func main() {

	file, err := os.Create("/tmp/alilog.txt")
	if err != nil {
		log.Fatal("Error while trying to create file", err)
	}

	defer file.Close()

	for i := 1; i < 10001; i++ {
		fmt.Fprintf(file, "Hi, this is line number %v\n", i)
	}

	// Task2: call fetchLogs
	fetchLogs("/tmp/alilog.txt", 96, 13, true)
}

func fetchLogs(logFilePath string, numberOfLines, offset int64, reverse bool) {
	// Calculate the offset position and the fetch size
	byteToRead, byteOffset := calcByteToReadAndOffset(numberOfLines, offset, reverse)
	file, err := os.Open(logFilePath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer file.Close()

	byteBuffer := make([]byte, byteToRead)

	// Reads len(byteBuffer) bytes from the File starting at byteOffset.
	fetchBytes, err := file.ReadAt(byteBuffer, byteOffset)

	if err != nil {
		if err != io.EOF {
			log.Fatal("Error while trying to fetchLogs", err.Error())
			return
		}
	}

	fmt.Println(string(byteBuffer[:fetchBytes]))
}

func calcByteToReadAndOffset(numberOfLines, offset int64, reverse bool) (int64, int64) {
	// Check reverse flag
	if reverse {
		offset = 10000 - numberOfLines - offset
	}

	// Calculate totalLines: sum of numberOfLines and offset
	totalLines := numberOfLines + offset
	byteOffset := offset * 25
	byteToRead := totalLines * 25

	// Calculate byteOffset and byteToRead by iterating the totalLines
	// And add the delta bytes
	// e.x. if offset = 13 and numberOfLines = 96
	// byteOffset will be 342 (9*26 + 4*27)
	// byteToRead will be 2602 (86*27 + 10*28)
	for i := 0; totalLines > 0; i++ {
		reduce := 9 * int64(math.Pow10(i))
		byteOffset += offset
		offset -= reduce

		byteToRead += totalLines
		totalLines -= reduce

		if offset < 0 {
			offset = 0
		}
	}

	return byteToRead - byteOffset, byteOffset
}
