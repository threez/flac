# flac [![GoDoc](https://godoc.org/github.com/threez/flac?status.svg)](https://godoc.org/github.com/threez/flac)

Package flac implements a simple **command line** based flac **encoder** and **decoder** pipeline.

Encoding (without err handling):

    source, _ := os.Open("test.wav")
	r, _ := NewEncoder(source)
	encoded, _ := ioutil.ReadAll(r)

Decoding (without err handling):

	source, _ := os.Open("test.flac")
	r, _ := NewDecoder(source)
	decoded, _ := ioutil.ReadAll(r)