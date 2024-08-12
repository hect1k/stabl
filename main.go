package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Config struct {
	Port    int
	Servers []string
}

var config Config
var currentServer int

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	url, err := getServer()
	if err != nil {
		http.Error(w, "All servers are down", http.StatusServiceUnavailable)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	log.Println("Forwarding request to:", url)
	proxy.ServeHTTP(w, r)
}

func getServer() (*url.URL, error) {
	for i := 0; i < len(config.Servers); i++ {
		currentServer = (currentServer + 1) % len(config.Servers)
		url, err := url.Parse(config.Servers[currentServer])
		if err != nil {
			log.Println("Error parsing URL:", err)
			continue
		} else if !checkHealth(url) {
			log.Println("Server is down:", url)
			continue
		} else {
			return url, nil
		}
	}

	log.Println("All servers are down")
	return nil, fmt.Errorf("all servers are down")
}

func checkHealth(url *url.URL) bool {
	_, err := http.Get(url.String())
	if err != nil {
		return false
	}
	return true
}
