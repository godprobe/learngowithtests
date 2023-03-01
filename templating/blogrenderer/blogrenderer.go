package blogrenderer

import "bytes"

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(buf *bytes.Buffer, post Post) error {
	return nil
}
