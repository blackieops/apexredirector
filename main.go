package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ValidateHost(host string, allowedHosts *[]string) bool {
	if len(*allowedHosts) == 0 {
		// If no allowlist has been set, just allow everything.
		return true
	}

	for _, a := range *allowedHosts {
		if a == host {
			return true
		}
	}

	return false
}

func BuildResponseURL(requestURL *url.URL, forceSecure bool) string {
	if forceSecure {
		requestURL.Scheme = "https"
	} else {
		requestURL.Scheme = "http"
	}

	requestURL.Host = "www." + requestURL.Host

	return requestURL.String()
}

func getListenPort() string {
	if port, isSet := os.LookupEnv("PORT"); isSet {
		return ":" + port
	}

	return ":8080"
}

func getSecure() bool {
	if _, isSet := os.LookupEnv("SECURE"); isSet {
		return true
	}

	return false
}

func getAllowedHosts() []string {
	if hosts, isSet := os.LookupEnv("ALLOWED_HOSTS"); isSet {
		return strings.Split(hosts, ",")
	} else {
		return []string{}
	}
}

func main() {
	listenPort := getListenPort()
	forceSecure := getSecure()
	allowedHosts := getAllowedHosts()

	fmt.Println("\n" +
		"░█▀█░█▀█░█▀▀░█░█░░░█▀▄░█▀▀░█▀▄░▀█▀░█▀▄░█▀▀░█▀▀░▀█▀░█▀█░█▀▄\n" +
		"░█▀█░█▀▀░█▀▀░▄▀▄░░░█▀▄░█▀▀░█░█░░█░░█▀▄░█▀▀░█░░░░█░░█░█░█▀▄\n" +
		"░▀░▀░▀░░░▀▀▀░▀░▀░░░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀\n" +
		"\n")

	fmt.Println("Listening on HTTP " + listenPort + "...\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method=GET host=%s\n", r.Host)

		// There is no "full URL" field in the request... So we have to
		// manually add the host from the request into our own url object.
		requestURL := *r.URL
		requestURL.Host = r.Host

		if ValidateHost(requestURL.Host, &allowedHosts) {
			http.Redirect(w, r, BuildResponseURL(&requestURL, forceSecure), 308)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(listenPort, nil)
}
