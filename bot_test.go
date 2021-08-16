package bot

import "testing"

const (
	TOKEN   = ""
	CHAT_ID = 0
)

func TestSendMessage(t *testing.T) {
	bot := NewAPIBot(TOKEN)

	msg := &Message{
		ChatID:                CHAT_ID,
		Text:                  "test",
		ParseMode:             "",
		DisableWebPagePreview: false,
		DisableNotification:   false,
		ReplyToMessageID:      0,
	}

	err := bot.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}
