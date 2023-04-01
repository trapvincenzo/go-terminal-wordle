package reader

import (
	"bufio"
	"os"
	"strings"
)

type Bufio struct {
	reader *bufio.Reader
}

// var stdin bytes.Buffer

// stdin.Write([]byte("hunter2\n"))

func (r *Bufio) ReadString() string {
	if r.reader == nil {
		r.reader = bufio.NewReader(os.Stdin)
	}

	text, _ := r.reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}
