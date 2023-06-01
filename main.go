package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	webhookKey := flag.String("webhookKey", "webhook", "Webhook Key to send POST request to bot. \n Format: XXXXXXXXX/YYYYYYYY/XXXXXXXXX")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	contentType := "application/json"

	currentTime := currentTime()

	text := fmt.Sprintf("Stage deployed at: %s", currentTime)
	webhookPath := fmt.Sprintf("https://hooks.slack.com/services/%s", *webhookKey)

	message := Message{
		Text: text,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	resp, err := http.Post(webhookPath, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	defer resp.Body.Close()

	infoLog.Printf("Message `%s` was sent successfully", text)
}

func currentTime() string {
	now := time.Now().Local().Add(time.Hour * 2)
	formattedTime := now.Format("02.01.2006 15:04:05")

	return formattedTime
}
