package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type InputData struct {
	webhookPath string
}

type Message struct {
	Text string `json:"text"`
}

func existWithError(err error) {
	fmt.Fprintf(os.Stderr, "Slack notification error: %s \n", err)
}

func main() {
	tempHook := "T059NH4TER0/B05A17WKJDP/ULWETY3DAgzi4llUEkoFMxGF"
	// envKey := flag.String("envKey", "webhook", "Set env key from which cli has to get Webhook Key")
	webhookPath := fmt.Sprintf("https://hooks.slack.com/services/%s", tempHook)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	currentTime := time.Now().Format("2022-04-30 11:30:42")
	text := fmt.Sprintf("Stage deployed at: %s", currentTime)

	message := Message{
		Text: text,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	// client := http.Client{}

	// fmt.Println(webhookPath)
	// request, err := http.NewRequest(http.MethodPost, webhookPath, string(jsonData))
	// if err != nil {
	// 	errorLog.Fatal(err)
	// }
	resp, err := http.Post(webhookPath, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	defer resp.Body.Close()

	infoLog.Printf("Message `%s` was sent successfully", text)
}
