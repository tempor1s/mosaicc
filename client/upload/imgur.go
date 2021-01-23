package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
)

// imgurResponse represents the response that we get from imgur
type imgurResponse struct {
	Data    map[string]interface{} `json:"data"`    // response data
	Success bool                   `json:"success"` // if the request was a success
	Status  int                    `json:"status"`  // http status code
}

// ToImgur will upload the given contents to imgur
func ToImgur(name string, contents []byte) (string, error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	part, err := writer.CreateFormFile("image", name)
	if err != nil {
		return "", err
	}
	// write the file contents to the buffer
	part.Write(contents)
	writer.Close()

	// create the new http request
	req, _ := http.NewRequest("POST", "https://api.imgur.com/3/upload", buf)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response from imgur
	var response imgurResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// get link from response
	link, ok := response.Data["link"].(string)
	if !ok {
		return "", errors.New("could not get link from response")
	}

	return link, nil
}
