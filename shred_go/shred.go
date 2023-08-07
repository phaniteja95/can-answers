package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

const CHUNK_SIZE int64 = (64 * 1024)

func WriteRandData(f *os.File, size int64, off int64) {
	buffer := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(buffer)
	_, err := f.WriteAt(buffer, off)
	if err != nil {
		fmt.Println("ERROR: Write to file failed")
		return	
	}
}

func WriteInChunks (f *os.File, size int64) {
	var size_rem int64 = size
	var off int64 = 0

	// Write the file full of random data.
	for size_rem > 0 {
		if size_rem < CHUNK_SIZE {
			WriteRandData(f, size_rem, off)	
		} else {
			WriteRandData(f, CHUNK_SIZE, off)
		}
		size_rem -= CHUNK_SIZE
		off += CHUNK_SIZE
	}
}

func Shred(name string) {
	// Open the file
	f, err := os.OpenFile(name, os.O_WRONLY, 0)
	if err != nil {
		fmt.Println("ERROR: File open error")
		return
	}

	defer f.Close()

	fmt.Printf("INFO: %s opened\n", name)

	stats, err := os.Stat(name)
	if err != nil {
		fmt.Println("ERROR: Unable to get file info")
		return
	}

	// Get file size
	size := stats.Size()
	i := 0

	fmt.Printf("INFO: File size %d\n", size)
	fmt.Printf("INFO: Writing file now\n")

	// Write file in chunks. This avoids having a large buffer in memory
	// when writing a large file.
	for i < 3 {

		WriteInChunks(f, size)

		i++
		fmt.Printf("INFO: Finished writing file %d time\n", i)
	}

	os.Remove(name);
	fmt.Printf("INFO: Deleted file %s\n", name)

	return;
}

func main() {
	filenames := [...]string{"random.txt", "nonexistent.txt", "empty.txt", "image.jpeg", "video.mp4"}

	i := 0
	length := len(filenames)

	for i < length {
		fmt.Printf("********** Test Case %d: %s **********\n", (i+1), filenames[i])
		Shred(filenames[i])
		fmt.Printf("*************************************\n")
		i++;
	}
}