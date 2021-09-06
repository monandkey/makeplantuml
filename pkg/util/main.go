package util

import (
	"os"
)

func FileRead(fileName string) string {
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

func (l Location) ValidateLocation() bool {
	_, err := os.Stat(l.Path)
	return err != nil
}
