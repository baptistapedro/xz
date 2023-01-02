package myfuzz

import (
    "bytes"
    "io"
    "github.com/ulikunitz/xz"
)

func Fuzz(data []byte) int {
    var buf bytes.Buffer
    // compress text
    w, err := xz.NewWriter(&buf)
    if err != nil {
	return 1
    }
    if _, err := io.WriteString(w, string(data)); err != nil {
        return 1
    }
    if err := w.Close(); err != nil {
       return 1
    }
    return 0
}
nasa@nasa
