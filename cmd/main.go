package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PackageJSON struct {
	Version string `json:"version"`
}

type Message struct {
	Text string `json:"text"`
}

func main() {
	webhookKey := flag.String("webhookKey", "webhook", "Webhook Key to send POST request to bot. \n Format: XXXXXXXXX/YYYYYYYY/XXXXXXXXX")
	flag.Parse()

	contentType := "application/json"

	currentTime := currentTime()
	version := getVersion("projects/wflow-main-app/package.json")

	text := fmt.Sprintf("Stage deployed: \n - Datetime: *%s*\n - Version: *%s*", currentTime, version)
	webhookPath := fmt.Sprintf("https://hooks.slack.com/services/%s", *webhookKey)

	message := Message{
		Text: text,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		ErrorLog().Fatal(err)
		return
	}

	resp, err := http.Post(webhookPath, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		ErrorLog().Fatal(err)
		return
	}

	defer resp.Body.Close()

	InfoLog().Printf("Message `%s` was sent successfully", text)
}

func getVersion(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		ErrorLog().Fatal(err)
	}

	var packageJSON PackageJSON
	err = json.Unmarshal(data, &packageJSON)
	if err != nil {
		ErrorLog().Fatal(err)
	}

	version := packageJSON.Version

	return version
}

func currentTime() string {
	now := time.Now().Local().Add(time.Hour * 2)
	formattedTime := now.Format("02.01.2006 15:04:05")

	return formattedTime
}
