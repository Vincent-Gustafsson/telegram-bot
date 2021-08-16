package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Bot interface {
	SendMessage()
}

type APIBot struct {
	Token   string
	BaseURL string
}

func NewAPIBot(token string) *APIBot {
	return &APIBot{
		Token:   token,
		BaseURL: "https://api.telegram.org/bot" + token,
	}
}

func (b *APIBot) SendMessage(msg *Message) error {
	url := b.BaseURL + "/sendMessage"

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(msgJson),
	)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		errMessage, err := getError(res.Body)
		if err != nil {
			return err
		}

		return errors.New(errMessage)
	}
	return nil
}
