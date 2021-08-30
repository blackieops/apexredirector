package main

import (
	"fmt"
	"net/http"
	"net/url"
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

func BuildResponseURL(requestURL *url.URL, subdomain string, forceSecure bool) string {
	if forceSecure {
		requestURL.Scheme = "https"
	} else {
		requestURL.Scheme = "http"
	}

	requestURL.Host = subdomain + "." + requestURL.Host

	return requestURL.String()
}

func main() {
	listenPort := getListenPort()
	forceSecure := getSecure()
	allowedHosts := getAllowedHosts()
	subdomain := getSubdomain()
	hostOverride := getHostOverride()

	fmt.Println("\n" +
		"░█▀█░█▀█░█▀▀░█░█░░░█▀▄░█▀▀░█▀▄░▀█▀░█▀▄░█▀▀░█▀▀░▀█▀░█▀█░█▀▄\n" +
		"░█▀█░█▀▀░█▀▀░▄▀▄░░░█▀▄░█▀▀░█░█░░█░░█▀▄░█▀▀░█░░░░█░░█░█░█▀▄\n" +
		"░▀░▀░▀░░░▀▀▀░▀░▀░░░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀\n" +
		"\n")

	fmt.Println("Listening on HTTP " + listenPort + "...\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method=GET host=%s\n", r.Host)

		requestURL := *r.URL

		if hostOverride != "" {
			// If HOST_OVERRIDE is set, use that instead of the request host.
			requestURL.Host = hostOverride
		} else {
			// There is no "full URL" field in the request... So we have to
			// manually add the host from the request into our own url object.
			requestURL.Host = r.Host
		}

		if ValidateHost(requestURL.Host, &allowedHosts) {
			http.Redirect(w, r, BuildResponseURL(&requestURL, subdomain, forceSecure), 308)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(listenPort, nil)
}
