package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	
	n, err = r.r.Read(b)
	if err != nil {
		return n, err
	}
	
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] <= 'Z' : 
			b[i] = (b[i] - 'A' + 13) % 26 + 'A'
		case b[i] >= 'a' && b[i] <= 'z' : 
			b[i] = (b[i] - 'a' + 13) % 26 + 'a'
		}
	}
		
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
