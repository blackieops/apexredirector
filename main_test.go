package main

import (
	"net/url"
	"testing"
)

func TestBuildResponseURLAddsWWWAndProtocol(t *testing.T) {
	source, _ := url.Parse("https://example.com")
	result := BuildResponseURL(source, "www", true)

	if result != "https://www.example.com" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLAddsWWWAndProtocolInsecure(t *testing.T) {
	source, _ := url.Parse("http://example.com")
	result := BuildResponseURL(source, "www", false)

	if result != "http://www.example.com" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLPreservesPath(t *testing.T) {
	source, _ := url.Parse("https://example.com/asdf/one/2")
	result := BuildResponseURL(source, "www", true)

	if result != "https://www.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLPreservesPathInsecure(t *testing.T) {
	source, _ := url.Parse("http://example.com/asdf/one/2")
	result := BuildResponseURL(source, "www", false)

	if result != "http://www.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLWithCustomSubdomain(t *testing.T) {
	source, _ := url.Parse("http://example.com/asdf/one/2")
	result := BuildResponseURL(source, "blog", false)

	if result != "http://blog.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestValidateHostWhenHostIsAllowed(t *testing.T) {
	allowedHosts := []string{"example.com"}
	result := ValidateHost("example.com", &allowedHosts)

	if result != true {
		t.Errorf("ValidateHost failed to validate host correctly.")
	}
}

func TestValidateHostWhenHostIsDisallowed(t *testing.T) {
	allowedHosts := []string{"example.com"}
	result := ValidateHost("example.biznetcom", &allowedHosts)

	if result != false {
		t.Errorf("ValidateHost allowed invalid host.")
	}
}

func TestValidateHostWhenAllowedUnset(t *testing.T) {
	allowedHosts := []string{}
	result := ValidateHost("example.com", &allowedHosts)

	if result != true {
		t.Errorf("ValidateHost blocked request when no allowed hosts are set.")
	}
}
