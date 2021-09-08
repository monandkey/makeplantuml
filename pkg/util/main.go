package util

import (
	"os"
    "errors"
)

func FileRead(fileName string) (string, error) {
    fp, err := os.Open(fileName)
    if err != nil {
        return "", err
    }
    defer fp.Close()

    buf := make([]byte, 102400 /* 10KB */)
    for {
        n, err := fp.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            return "", err
        }
    }
	return string(buf), nil
}

func FileRemove(fileName string) error {
    if FileExist(fileName) {
        return errors.New("The file does not exist.")
    }

    err := os.RemoveAll(fileName)
    return err
}

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err != nil
}

func (l Location) ValidateLocation() bool {
	_, err := os.Stat(l.Path)
	return err != nil
}
