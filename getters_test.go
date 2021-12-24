package main

import (
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

func TestGetSecureWhenIsSet(t *testing.T) {
	os.Setenv("SECURE", "true")

	actual := getSecure()

	if actual != true {
		t.Errorf("getSecure should have been true but was false.")
	}
}

func TestGetSecureWhenIsNotSet(t *testing.T) {
	os.Unsetenv("SECURE")

	actual := getSecure()

	if actual != false {
		t.Errorf("getSecure should have been false but was true.")
	}
}

func TestGetAllowedHostsWhenIsNotSet(t *testing.T) {
	os.Unsetenv("ALLOWED_HOSTS")

	actual := getAllowedHosts()

	if !testStringSliceEqual(actual, make([]string, 0)) {
		t.Errorf("getAllowedHosts should have been empty but was not.")
	}
}

func TestAllowedHostsWhenOne(t *testing.T) {
	os.Setenv("ALLOWED_HOSTS", "example.com")

	actual := getAllowedHosts()

	if !testStringSliceEqual(actual, []string{"example.com"}) {
		t.Errorf("getAllowedHosts with one host contained unexpected hosts!")
	}
}

func TestAllowedHostsWhenMany(t *testing.T) {
	os.Setenv("ALLOWED_HOSTS", "example.com,example.org")

	actual := getAllowedHosts()
	expected := []string{"example.com", "example.org"}

	if !testStringSliceEqual(actual, expected) {
		t.Errorf("getAllowedHosts with many hosts contained unexpected hosts!")
	}
}

func TestGetSubdomainWhenIsNotSet(t *testing.T) {
	os.Unsetenv("SUBDOMAIN")

	actual := getSubdomain();

	if actual != "www" {
		t.Errorf("getSubdomain returned non-default value: '%s'.", actual)
	}
}

func TestGetSubdomainWhenIsSet(t *testing.T) {
	os.Setenv("SUBDOMAIN", "www3")

	actual := getSubdomain();

	if actual != "www3" {
		t.Errorf("getSubdomain returned incorrect value: '%s'.", actual)
	}
}

func TestHostOverrideWhenIsNotSet(t *testing.T) {
	os.Unsetenv("HOST_OVERRIDE")

	actual := getHostOverride()

	if actual != "" {
		t.Errorf("getHostOverride returned unexpected value: '%s'", actual)
	}
}

func TestHostOverrideWhenIsSet(t *testing.T) {
	os.Setenv("HOST_OVERRIDE", "example.com")

	actual := getHostOverride()

	if actual != "example.com" {
		t.Errorf("getHostOverride returned unexpected value: '%s'", actual)
	}
}

// "rip stdlib" section:

func testStringSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}
