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
	t.Run("basic slow server vs fast server test", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error, but didn't get one")
		}
	})
}

func makeDelayedServer(delayAmount time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayAmount)
		w.WriteHeader(http.StatusOK)
	}))
}
