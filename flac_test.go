package flac

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	// load reference file
	expected, err := ioutil.ReadFile("testfiles/original.wav.flac")
	assert.NoError(t, err)

	// open source file
	source, err := os.Open("testfiles/original.wav")
	assert.NoError(t, err)

	// start encoder
	r, err := NewEncoderContext(context.Background(), source)
	assert.NoError(t, err)

	// encode while reading
	encoded, err := ioutil.ReadAll(r)
	assert.NoError(t, err)

	// compare results
	assert.EqualValues(t, expected, encoded)
}

func TestDecoding(t *testing.T) {
	// load reference file
	expected, err := ioutil.ReadFile("testfiles/decoded.original.wav.flac")
	assert.NoError(t, err)

	// open source file
	source, err := os.Open("testfiles/original.wav.flac")
	assert.NoError(t, err)

	// start encoder
	r, err := NewDecoderContext(context.Background(), source)
	assert.NoError(t, err)

	// decode while reading
	decoded, err := ioutil.ReadAll(r)
	assert.NoError(t, err)

	// compare results
	assert.EqualValues(t, expected, decoded)
}
