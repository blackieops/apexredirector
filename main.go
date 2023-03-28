package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.yml", "Provide a path to the configuration yaml file. Defaults to config.yml in the current directory.")
	flag.Parse()

	config, err := ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	listenPort := getListenPort()

	fmt.Printf("\n" +
		"░█▀█░█▀█░█▀▀░█░█░░░█▀▄░█▀▀░█▀▄░▀█▀░█▀▄░█▀▀░█▀▀░▀█▀░█▀█░█▀▄\n" +
		"░█▀█░█▀▀░█▀▀░▄▀▄░░░█▀▄░█▀▀░█░█░░█░░█▀▄░█▀▀░█░░░░█░░█░█░█▀▄\n" +
		"░▀░▀░▀░░░▀▀▀░▀░▀░░░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀\n" +
		"\n")

	fmt.Println("Listening on HTTP " + listenPort + "...\n")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method=GET healthz=true")
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method=GET host=%s\n", r.Host)

		requestURL := *r.URL
		redirect, err := GetRedirectRule(config.Redirects, r.Host)

		if err == nil {
			http.Redirect(w, r, BuildResponseURL(&requestURL, redirect, config.Secure), 308)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(listenPort, nil)
}

func GetRedirectRule(redirects []RedirectConfig, host string) (*RedirectConfig, error) {
	for _, redirect := range redirects {
		if redirect.FromHost == host {
			return &redirect, nil
		}
	}

	return &RedirectConfig{}, errors.New("Redirect host is not configured.")
}

func BuildResponseURL(requestURL *url.URL, redirect *RedirectConfig, forceSecure bool) string {
	if forceSecure {
		requestURL.Scheme = "https"
	}

	requestURL.Host = redirect.ToHost

	return requestURL.String()
}

func getListenPort() string {
	if port, isSet := os.LookupEnv("PORT"); isSet {
		return ":" + port
	}

	return ":8080"
}
