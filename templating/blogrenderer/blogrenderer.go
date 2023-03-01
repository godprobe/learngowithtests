package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	// write data to w
	_, err := fmt.Fprintf(w, `<h1>%s</h1>
<p>%s</p>
Tags: <ul><li>%s</li><li>%s</li></ul>`, p.Title, p.Description, p.Tags[0], p.Tags[1])
	return err
}
