package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// hit two sites with an HTTP GET
// return the URL that returned first
// if neither return after 10 seconds, return error

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(delayAmount time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayAmount)
		w.WriteHeader(http.StatusOK)
	}))
}
