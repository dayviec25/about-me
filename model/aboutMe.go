package model

type AboutMe struct {
	Title          string
	Description    string
	Traits         []string
	Projects       []Project
	WorkExperience []WorkExperience
	Reviews        []Review
	Blogs          []Blog
	Contact        Contact
}
