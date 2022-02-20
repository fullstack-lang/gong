package walk

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

func CaptureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

// func main() {
// 	re := captureOutput(func() {
// 		fmt.Println("test fmt 1")
// 		os.Stderr.WriteString("test stderr 1\n")
// 		fmt.Println("test fmt 2")
// 		fmt.Println("test fmt 3")
// 		log.Println("test log 1")
// 	})
// 	fmt.Printf(re)
// 	fmt.Println("exit")
// }
