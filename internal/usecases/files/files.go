package files

import (
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func SaveFile(fileHeader *multipart.FileHeader, filename string) (string, error) {
	time := time.Now()
	path := filepath.Join(
		// configs.BasePath, 
		"static",
		strconv.Itoa(time.Year()), 
		strconv.Itoa(int(time.Month())),
		strconv.Itoa(time.Day()),
	)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	extensions, err := mime.ExtensionsByType(fileHeader.Header.Get("Content-Type"))
	var extension string 
	if err == nil && len(extensions) > 0 {
		extension = extensions[0]
	}
	path = filepath.Join(path, filename + extension)
	err =  os.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

//DRY
func SaveFileBase64(date string, filename string) (string, error) {
	time := time.Now()
	path := filepath.Join(
		// configs.BasePath, 
		"static",
		strconv.Itoa(time.Year()), 
		strconv.Itoa(int(time.Month())),
		strconv.Itoa(time.Day()),
	)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	buf, err := base64.StdEncoding.DecodeString(date)
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, filename)
	err =  os.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}