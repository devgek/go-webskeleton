package helper

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//GoPaths returns the GOPATH as an array of paths
func GoPaths() []string {
	return strings.Split(os.Getenv("GOPATH"), ":")
}

// ExitOnError logs error message in fatal mode.
func ExitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s\n%s", msg, err.Error())
	}
}

// PanicOnError logs error message in fatal mode.
func PanicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//IsWindows detect os windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// RecursiveSearchReplaceFiles find and replace various strings defined in replacers.
func RecursiveSearchReplaceFiles(fullpath string, replacers map[string]string) error {
	fileOrDirList := []string{}
	err := filepath.Walk(fullpath, func(path string, f os.FileInfo, err error) error {
		fileOrDirList = append(fileOrDirList, path)
		return nil
	})

	if err != nil {
		return err
	}

	for _, fileOrDir := range fileOrDirList {
		fileInfo, _ := os.Stat(fileOrDir)
		if !fileInfo.IsDir() {
			for oldString, newString := range replacers {
				contentBytes, _ := ioutil.ReadFile(fileOrDir)
				newContentBytes := bytes.Replace(contentBytes, []byte(oldString), []byte(newString), -1)

				err := ioutil.WriteFile(fileOrDir, newContentBytes, fileInfo.Mode())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//EncryptPassword create hashed password
func EncryptPassword(password string) string {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	PanicOnError(err)

	return string(encryptedPass)
}

//ComparePassword compare hashed password and possible plaintext equivalent
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
