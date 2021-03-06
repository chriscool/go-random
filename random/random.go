package main

import (
	randcrypto "crypto/rand"
	randmath "math/rand"
	"fmt"
	"io"
	"os"
	"strconv"
	"bytes"
)

func main() {
	l := len(os.Args)
	if l != 2 && l != 3 {
		usageError()
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usageError()
	}

	if l == 2 {
		err = writeRandomBytes(num, os.Stdout)
	} else {
		seed, err2 := strconv.Atoi(os.Args[2])
		if err2 != nil {
			usageError()
		}
		err = writePseudoRandomBytes(num, os.Stdout, seed)
	}

	if err != nil {
		die(err)
	}
}

func usageError() {
	fmt.Fprintf(os.Stderr, "Usage: %s <count> [<seed>]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "If <seed> is given, output <count> pseudo random bytes made from <seed> (from Go's math/rand)\n")
	fmt.Fprintf(os.Stderr, "Otherwise, output <count> random bytes (from Go's crypto/rand)\n")
	os.Exit(-1)
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v", err)
	os.Exit(-1)
}

func writeRandomBytes(num int, w io.Writer) error {
	r := &io.LimitedReader{R: randcrypto.Reader, N: int64(num)}
	_, err := io.Copy(w, r)
	return err
}

func writePseudoRandomBytes(num int, w io.Writer, seed int) error {
	randmath.Seed(int64(seed))
	b := make([]byte, num)
	for i := 0 ; i < num ; i++ {
		b[i] = byte(randmath.Intn(256))
	}
	r := bytes.NewReader(b)
	_, err := io.Copy(w, r)
	return err
}

