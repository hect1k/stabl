package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

type Config struct {
	Port       int
	Servers    []string
	checkAfter int
	LogFile    string
}

type Server struct {
	URL     string
	Proxy   *httputil.ReverseProxy
	Healthy bool
}

type ServerPool struct {
	Servers []Server
}

var config Config
var currentServer int
var serverPool ServerPool
var checkAfter int
var logFile *os.File

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {

	configFilePath := flag.String("config", "config.json", "Path to the configuration file")
	portNumber := flag.Int("port", 0, "Port number to listen on")
	logFileFlag := flag.String("logFile", "", "Path to the log file")
	flag.Parse()

	file, err := os.Open(*configFilePath)
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	if *logFileFlag != "" {
		logFile, err := os.OpenFile(*logFileFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("Error opening log file:", err)
		}
		InfoLogger = log.New(logFile, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(logFile, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(logFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else if config.LogFile != "" {
		logFile, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("Error opening log file:", err)
		}
		InfoLogger = log.New(logFile, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(logFile, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(logFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		InfoLogger = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if *portNumber != 0 {
		config.Port = *portNumber
	}

	if config.Port == 0 {
		log.Fatal("Port number not specified in config.json or as a flag")
	}

	if len(config.Servers) == 0 {
		log.Fatal("No servers specified in config.json")
	}

	serverPool, err = createPool()
	if err != nil {
		log.Fatal("Error creating server pool:", err)
	}

	currentServer = 0

	go healthCheck()

	http.HandleFunc("/", forwardRequest)
	log.Println("Starting load balancer on port:", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server, err := getServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	InfoLogger.Println("Forwarding request to:", server.URL)
	server.Proxy.ServeHTTP(w, r)
	return
}

func createPool() (ServerPool, error) {
	var servers []Server
	for _, serverUrl := range config.Servers {
		proxyURL, err := url.Parse(serverUrl)
		if err != nil {
			ErrorLogger.Println("Error parsing server URL:", err)
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
	ErrorLogger.Println("All servers are down!")
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
				WarningLogger.Println("Server unhealthy:", server.URL)
			} else {
				serverPool.Servers[i].Healthy = true
			}
		}
		time.Sleep(time.Duration(checkAfter) * time.Minute)
	}
}
