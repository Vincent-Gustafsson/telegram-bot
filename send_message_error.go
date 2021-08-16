package bot

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type sendMessageError struct {
	Description string `json:"description"`
}

func getError(body io.ReadCloser) (string, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	errDescription := sendMessageError{}

	err = json.Unmarshal(bodyBytes, &errDescription)
	if err != nil {
		return "", err
	}
	return errDescription.Description, nil
}
