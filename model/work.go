package model

type WorkExperience struct {
	Date        string
	Company     string
	Role        string
	Description string
}

type WorkSection struct {
	Work []WorkExperience
}
