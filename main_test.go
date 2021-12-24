package main

import (
	"net/url"
	"os"
	"testing"
)

func TestGetListenPortWhenIsNotSet(t *testing.T) {
	os.Unsetenv("PORT")

	actual := getListenPort()

	if actual != ":8080" {
		t.Errorf("getListenPort had unexpected default '%s'", actual)
	}
}

func TestGetListenPortWhenIsSet(t *testing.T) {
	os.Setenv("PORT", "9876")

	actual := getListenPort()

	if actual != ":9876" {
		t.Errorf("getListenPort had incorrect port '%s'", actual)
	}
}

func TestGetRedirectRuleWhenHostMatches(t *testing.T) {
	redirects := []RedirectConfig{
		{FromHost: "example.com", ToHost: "www.example.com"},
		{FromHost: "example.net", ToHost: "www.example.net"},
	}
	result, err := GetRedirectRule(redirects, "example.com")

	if err != nil {
		t.Errorf("GetRedirectRule did not find a host, despite being configured.")
	}

	if result.ToHost != "www.example.com" {
		t.Errorf("GetRedirectRule found the wrong redirect for the host.")
	}
}

func TestGetRedirectRuleWhenHostIsUnknown(t *testing.T) {
	redirects := []RedirectConfig{
		{FromHost: "example.com", ToHost: "www.example.com"},
	}
	_, err := GetRedirectRule(redirects, "example.net")

	if err == nil {
		t.Errorf("GetRedirectRule returned a redirect when it should not have.")
	}
}

func TestBuildResponseURLAddsWWWAndProtocol(t *testing.T) {
	source, _ := url.Parse("https://example.com")
	redirect := &RedirectConfig{FromHost: "example.com", ToHost: "www.example.com"}
	result := BuildResponseURL(source, redirect, true)

	if result != "https://www.example.com" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLAddsWWWAndProtocolInsecure(t *testing.T) {
	source, _ := url.Parse("http://example.com")
	redirect := &RedirectConfig{FromHost: "example.com", ToHost: "www.example.com"}
	result := BuildResponseURL(source, redirect, false)

	if result != "http://www.example.com" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLPreservesPath(t *testing.T) {
	source, _ := url.Parse("https://example.com/asdf/one/2")
	redirect := &RedirectConfig{FromHost: "example.com", ToHost: "www.example.com"}
	result := BuildResponseURL(source, redirect, true)

	if result != "https://www.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLPreservesPathInsecure(t *testing.T) {
	source, _ := url.Parse("http://example.com/asdf/one/2")
	redirect := &RedirectConfig{FromHost: "example.com", ToHost: "www.example.com"}
	result := BuildResponseURL(source, redirect, false)

	if result != "http://www.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}

func TestBuildResponseURLWithSomethingOtherThanWWW(t *testing.T) {
	source, _ := url.Parse("http://example.com/asdf/one/2")
	redirect := &RedirectConfig{FromHost: "example.com", ToHost: "blog.example.com"}
	result := BuildResponseURL(source, redirect, false)

	if result != "http://blog.example.com/asdf/one/2" {
		t.Errorf("BuildResponseURL generated incorrect URL: %s", result)
	}
}
