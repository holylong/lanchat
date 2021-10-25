package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func DoSha256(filepath string) string {
	ha := sha256.New()
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error1!")
	}
	defer f.Close()
	if _, err := io.Copy(ha, f); err != nil {
		fmt.Println("error2")
	}

	str := hex.EncodeToString(ha.Sum(nil))
	fmt.Printf("%X", ha.Sum(nil))

	return str
}

func GetFileSize(fileurl string) int64 {
	var result int64
	filepath.Walk(fileurl, func(tfileurl string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// func IsDirectory(fileName string)(bool,error){
// 	info, err := os.Stat(fileName)
// 	if err != nil {
// 		log.Println("os.Stat err =", err)
// 		return false,nil
// 	}
// 	return info.IsDir(), nil
// }

func IsDirectory(fileName string) bool {
	info, err := os.Stat(fileName)
	if err != nil {
		log.Println("os.Stat err =", err)
		return false
	}
	return info.IsDir()
}

func GetFileModTime(fileurl string) int64 {
	f, err := os.Open(fileurl)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
