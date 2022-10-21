package uploader

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func UploadFile(file *os.File, fileName string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("asset", fileName)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	writer.Close()

	res, err := http.Post("https://transfer.sh/", writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}

	urlResult, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(urlResult), "\n"), nil
}
