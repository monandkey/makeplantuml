package makeplantuml

import (
	"os"
)

func fileRead(fileName string) string {
    fp, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    buf := make([]byte, 102400 /* 10KB */)
    for {
        n, err := fp.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            panic(err)
        }
    }
	return string(buf)
}

type location struct {
	path string
}

func (l location) validateLocation() bool {
	_, err := os.Stat(l.path)
	return err != nil
}
