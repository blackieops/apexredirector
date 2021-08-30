package main

import (
	"os"
	"strings"
)

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

func getSubdomain() string {
	if value, isSet := os.LookupEnv("SUBDOMAIN"); isSet {
		return value
	} else {
		return "www"
	}
}

func getHostOverride() string {
	if value, isSet := os.LookupEnv("HOST_OVERRIDE"); isSet {
		return value
	} else {
		return ""
	}
}
