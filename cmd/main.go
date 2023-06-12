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
	var text string

	version, err := getVersion("projects/wflow-main-app/package.json")
	if err != nil {
		text = fmt.Sprintf("Stage deployed: \n - Datetime: *%s*\n", currentTime)
		ErrorLog().Println(err)
	} else {
		text = fmt.Sprintf("Stage deployed: \n - Datetime: *%s*\n - Version: *%s*", currentTime, version)
	}

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

func getVersion(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var packageJSON PackageJSON
	err = json.Unmarshal(data, &packageJSON)
	if err != nil {
		return "", err
	}

	version := packageJSON.Version

	return version, nil
}

func currentTime() string {
	loc, err := time.LoadLocation("CET")
	if err != nil {
		ErrorLog().Fatal("Can't load location: ", err)
	}

	now := time.Now()
	czechia := now.In(loc).Format("02.01.2006 15:04:05")

	return czechia
}
