package tplfunc

import (
	"io"
	"os"
)

func MoveUploadFile(srcFile io.Reader, dist string) (err error) {
	distFile, err := os.OpenFile(dist, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	for {
		buf := make([]byte, 524288)
		_, err := srcFile.Read(buf)

		if err != nil {
			return nil //copy has been compelete
		}
		_, err = distFile.Write(buf) //copy 512kb each time
	}
	distFile.Close()
	return nil
}

func MoveFile(src string, dist string) (err error) {
	srcFile, err := os.OpenFile(src, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	distFile, err := os.OpenFile(dist, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	for {
		_, err := io.CopyN(distFile, srcFile, int64(524288)) //copy 512kb each time
		if err != nil {
			return nil //copy has been compelete
		}
	}
	defer srcFile.Close()
	defer distFile.Close()
	defer os.Remove(src)
	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func TryMkdir(dir string) error {
	if FileExists(dir) != true {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func Pwd() string {
	dir, _ := os.Getwd()
	return dir
}
