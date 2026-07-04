package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// OmikujiResult represents the fortune result
type OmikujiResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

// HostnameResult represents the hostname response
type HostnameResult struct {
	Hostname string `json:"hostname"`
}

// HealthResult represents the health check response
type HealthResult struct {
	Status string `json:"status"`
}

// StressResult represents the stress test response
type StressResult struct {
	Message string `json:"message"`
}

var omikujiList = []OmikujiResult{
	{Result: "aws大吉", Message: "すべてがうまくいく最高のawsです！"},
	{Result: "aws中吉", Message: "良いことが起こりそうなawsです。"},
	{Result: "aws小吉", Message: "小さなawsが訪れるでしょう。"},
	{Result: "aws吉", Message: "穏やかなawsになりそうです。"},
	{Result: "aws末吉", Message: "努力が実を結ぶawsがあります。"},
	{Result: "aws凶", Message: "慎重にawsすることをお勧めします。"},
	{Result: "aws大凶", Message: "今日は控えめにaws。"},
}

func main() {
	http.HandleFunc("/omikuji", omikujiHandler)
	http.HandleFunc("/hostname", hostnameHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/stress", stressHandler)

	fmt.Println("Starting Omikuji API server on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// omikujiHandler returns a random fortune
func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	result := omikujiList[rand.Intn(len(omikujiList))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// hostnameHandler returns the instance hostname/ID
func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("INSTANCE_ID")
	if hostname == "" {
		hostname, _ = os.Hostname()
	}

	result := HostnameResult{Hostname: hostname}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// healthHandler returns health status
func healthHandler(w http.ResponseWriter, r *http.Request) {
	result := HealthResult{Status: "healthy"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// stressHandler triggers CPU stress test for Auto Scaling verification
func stressHandler(w http.ResponseWriter, r *http.Request) {
	// Run CPU stress for 60 seconds in background
	go func() {
		cmd := exec.Command("stress-ng", "--cpu", "1", "--timeout", "60s")
		cmd.Run()
	}()

	result := StressResult{Message: "CPU stress test started for 60 seconds"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
