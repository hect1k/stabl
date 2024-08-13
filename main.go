package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

type Config struct {
	Port    int
	Servers []string
	checkAfter int
}

type Server struct {
	URL string
	Proxy *httputil.ReverseProxy
	Healthy bool
}

type ServerPool struct {
	Servers []Server
}

var config Config
var currentServer int
var serverPool ServerPool
var checkAfter int

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

	if config.Port == 0 {
		fmt.Println("No port specified in config.json")
		return
	}

	if len(config.Servers) == 0 {
		fmt.Println("No servers specified in config.json")
		return
	}

	serverPool, err = createPool()
	if err != nil {
		fmt.Println("Error creating server pool:", err)
		return
	}

	currentServer = 0

	go healthCheck()

	http.HandleFunc("/", forwardRequest)
	log.Println("Starting load balancer on port:", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port),nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server, err := getServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Forwarding request to:", server.URL)
	server.Proxy.ServeHTTP(w, r)
	return
}

func createPool() (ServerPool, error) {
	var servers []Server
	for _, serverUrl := range config.Servers {
		proxyURL, err := url.Parse(serverUrl)
		if err != nil {
			log.Println("Error parsing server URL:", err)
		}
		proxy := httputil.NewSingleHostReverseProxy(proxyURL)
		servers = append(servers, Server{proxyURL.String(), proxy, true})
	}
	return ServerPool{servers}, nil
}

func getServer() (Server, error) {
	for i := 0; i < len(serverPool.Servers); i++ {
		server := serverPool.Servers[currentServer]
		currentServer = (currentServer + 1) % len(serverPool.Servers)
		if server.Healthy {
			return server, nil
		}
	}
	log.Println("All servers are down!")
	return Server{}, fmt.Errorf("All servers are down :/")
}

func healthCheck() {

	if config.checkAfter <= 0 {
		checkAfter = 5
	} else {
		checkAfter = config.checkAfter
	}

	for {
		for i, server := range serverPool.Servers {
			_, err := http.Get(server.URL)
			if err != nil {
				serverPool.Servers[i].Healthy = false
				log.Println("Server unhealthy:", server.URL)
			} else {
				serverPool.Servers[i].Healthy = true
			}
		}
		time.Sleep(time.Duration(checkAfter) * time.Minute)
	}
}
