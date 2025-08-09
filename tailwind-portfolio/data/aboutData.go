package data

import (
	"portfolio/models"
)

var experience1 = models.Experience{
	Title:       "Go Backend Intern",
	Company:     "Holmona",
	StartDate:   "July",
	EndDate:     "August 2025",
	Description: "Working on a personal portfolio website using Go, HTML templates, and CSS as part of a backend development learning internship.",
}
var dataAbout = models.AboutData{
	AboutIntro: "Computer Science student at ISAMM who loves creating things with code.",
	About : "I enjoy turning ideas into real projects, learning new skills, and finding better ways to solve problems. Whether it’s a website, a small tool, or something just for fun, I like building and improving until it works well. Right now, I’m focusing on personal projects and real-world practice to grow as a developer.",
	Experience: []models.Experience {experience1, experience1,},
	Skills:     []string{"React", "Node.js", "Go", "Java", "Git", "GitHub", "Communication", "Problem Solving"},
}

func GetAboutData() models.AboutData {
	return dataAbout
}
