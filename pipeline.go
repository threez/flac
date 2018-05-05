package flac

import (
	"context"
	"io"
	"os"
	"os/exec"
)

// Cmd is the command line executable (can be an absolute path)
var Cmd = "flac"

// Stderr flac error output can be redirected somewhere else if os.Stderr
// is not wanted
var Stderr = os.Stderr

// EncodingArguments is the default configuration for the encoding cmd
var EncodingArguments = []string{
	"--best",   // Synonymous with -l 12 -b 4096 -m -r 6 -A tukey(0.5) -A partial_tukey(2) -A punchout_tukey(3)
	"--stdout", // Write output to stdout
	"--silent", // Do not write runtime encode/decode statistics
	"-",        // read from stdin
}

// DecodingArguments is the default configuration for the decoding cmd
var DecodingArguments = []string{
	"--decode", // Decode (the default behavior is to encode)
	"--silent", // Do not write runtime encode/decode statistics
	"-",        // read from stdin
}

// NewDecoderContext creates a new decoder pipeline that reads flac input
// via the reader and tranforms it to wav encoding with context
func NewDecoderContext(ctx context.Context, r io.Reader) (io.Reader, error) {
	return newCmdPipe(ctx, Cmd, DecodingArguments, r)
}

// NewDecoder creates a new decoder pipeline that reads flac input
// via the reader and tranforms it to wav encoding
func NewDecoder(r io.Reader) (io.Reader, error) {
	return NewDecoderContext(context.Background(), r)
}

// NewEncoderContext creates a new encoder pipeline that reads wav input via the reader
// and transforms it to flac encoding with context
func NewEncoderContext(ctx context.Context, r io.Reader) (io.Reader, error) {
	return newCmdPipe(ctx, Cmd, EncodingArguments, r)
}

// NewEncoder creates a new encoder pipeline that reads wav input via the reader
// and transforms it to flac encoding
func NewEncoder(r io.Reader) (io.Reader, error) {
	return NewEncoderContext(context.Background(), r)
}

func newCmdPipe(ctx context.Context, command string, args []string, r io.Reader) (io.Reader, error) {
	// pipe for the output
	rp, wp := io.Pipe()
	cmd := exec.CommandContext(ctx, command, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	// copy all data into the pipe
	go func() {
		_, err := io.Copy(stdin, r)
		if err != nil {
			wp.CloseWithError(err)
		}

		err = stdin.Close()
		if err != nil {
			wp.CloseWithError(err)
		}

		err = cmd.Wait()
		if err != nil {
			wp.CloseWithError(err)
		}

		err = wp.Close()
		if err != nil {
			wp.CloseWithError(err)
		}
	}()

	cmd.Stdout = wp
	cmd.Stderr = Stderr
	return rp, cmd.Start()
}
