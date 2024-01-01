package model

// Project represents the details of a project.
type Project struct {
	Title        string
	Logo         string
	LogoAlt      string
	Link         Link
	CompanyName  string
	Downloads    string
	ReviewRating string
	ReviewCount  string
}

// Link represents a hyperlink with a label.
type Link struct {
	Href  string
	Label string
}
