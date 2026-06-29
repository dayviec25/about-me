package model

import "html/template"

type Blog struct {
	Slug        string
	Title       string
	Date        string
	Author      string
	Description string
	// Content holds the rendered HTML body of the post. It is typed as
	// template.HTML so the template engine renders it without escaping.
	Content template.HTML
}
