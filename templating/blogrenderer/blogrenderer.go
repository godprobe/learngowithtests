package blogrenderer

import (
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	// write data to w
	w.Write([]byte("<h1>hello world</h1>"))
	return nil
}
