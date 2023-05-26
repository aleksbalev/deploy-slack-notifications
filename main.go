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

	"github.com/joho/godotenv"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	
	contentType := "application/json"
	
	err := godotenv.Load()
	if err != nil {
		errorLog.Fatal(err)
		return
	}
	
	currentTime := currentTime()
	webhookKey := getWebhookKey()
	
	text := fmt.Sprintf("Stage deployed at: %s", currentTime)
	webhookPath := fmt.Sprintf("https://hooks.slack.com/services/%s", webhookKey)

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
	now := time.Now()
	formattedTime := now.Format("02.01.2006 15:04:05")

	return formattedTime
}

func getWebhookKey() string {
	envKey := flag.String("envKey", "webhook", "Set env key from which cli has to get Webhook Key")
	flag.Parse()

	webhookKey := os.Getenv(*envKey)

	return webhookKey
}
