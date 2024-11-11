package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	var (
		token  = flag.String("token", "", "[*] Telegram bot token acquired form BotFather")
		chatID = flag.String("chat-id", "", "[*] ChatID to post to")
		msg    = flag.String("msg", "", "[*] Message to post to telegram")

		timeout = flag.Duration("timeout", 5*time.Second, "Request timeout")
	)
	flag.Parse()
	if token == nil || *token == "" {
		slog.Error("token is missing")
		return
	}
	if chatID == nil || *chatID == "" {
		slog.Error("chat-id is missing")
		return
	}
	if msg == nil || *msg == "" {
		slog.Error("msg is missing")
		return
	}

	if timeout == nil {
		slog.Error("timeout is nil")
		return
	}

	notify(
		*token, *chatID, *msg,
		*timeout,
	)
}

func notify(
	token, chatID, msg string,
	timeout time.Duration,
) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	body, err := json.Marshal(&sendMessageBody{
		ChatID: chatID,
		Text:   msg,
	})
	if err != nil {
		log.Panic(fmt.Errorf("unmarshal body: %w", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		log.Panic(fmt.Errorf("create request: %w", err))
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(fmt.Errorf("do request: %w", err))
	}

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		log.Panic(fmt.Errorf("unexpected response status: %d: %s for request: %s",
			resp.StatusCode,
			string(respBody),
			string(body),
		))
	}
}

type sendMessageBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
