package blogrenderer_test

import (
	"bytes"
	"fmt"
	"testing"

	"blogrenderer"
)

var (
	cases = []struct {
		Title       string
		Body        string
		Description string
		Tags        []string
	}{
		{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		},
		{
			Title:       "a new world",
			Body:        "This is another post",
			Description: "This is another description",
			Tags:        []string{"tdd", "go", "experiments"},
		},
	}
)

func TestRender(t *testing.T) {
	buf := bytes.Buffer{}

	for _, post := range cases {
		t.Run("convert a single post into HTML", func(t *testing.T) {
			buf.Reset()
			err := blogrenderer.Render(&buf, post)

			if err != nil {
				t.Fatal(err)
			}

			got := buf.String()
			want := fmt.Sprintf("<h1>%s</h1>", post.Title)
			if got != want {
				t.Errorf("got '%s' want '%s'", got, want)
			}
		})
	}
}
