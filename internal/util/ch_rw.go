package util

import "io"

type ChatReadWriter struct {
	// input  <-chan string
	// output chan<- string

	input  chan string
	output chan string
}

// func NewChatReadWriter(reader <-chan string, writer chan<- string) *ChatReadWriter {
// 	return &ChatReadWriter{
// 		input:  reader,
// 		output: writer,
// 	}
// }

func NewChatReadWriter(reader chan string, writer chan string) *ChatReadWriter {
	return &ChatReadWriter{
		input:  reader,
		output: writer,
	}
}

func (crw *ChatReadWriter) Write(p []byte) (n int, err error) {
	data := string(p)
	crw.output <- data
	return len(p), nil
}

func (crw *ChatReadWriter) Read(p []byte) (n int, err error) {

	data, ok := <-crw.input
	if !ok {
		return 0, io.EOF
	}

	copy(p, []byte(data))
	return len(data), nil
}

func (crw *ChatReadWriter) Input() chan string {
	return crw.input
}

func (crw *ChatReadWriter) Output() chan string {
	return crw.output
}
