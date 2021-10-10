package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type FakeReader struct {
	payload  []byte
	position int
}

func (f *FakeReader) Read(p []byte) (n int, err error) {
	p = f.payload
	f.position += len(f.payload)
	if f.position > len(f.payload) {
		return 0, io.EOF
	}
	return len(f.payload), nil
}
func (f *FakeReader) SetFakeBytes(p []byte) {
	f.position = 0
	f.payload = p
}
func ReadAllTheBytes(reader io.Reader) []byte {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func main() {
	fakeReader := &FakeReader{}
	want := []byte("when called, return this data")
	fakeReader.SetFakeBytes(want)
	got := ReadAllTheBytes(fakeReader)
	fmt.Printf("%d/%d bytes read.\n", len(got), len(want))
	fmt.Printf("got:%d\t\n want: %d \n", got, want)
}
