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
	AboutIntro: "Currently studying Computer Science at ISAMM.",
	About:      "Passionate about building things with code. I love solving problems and learning by doing — whether it's web apps, mobile tools, or anything in between. Right now, I'm focused on leveling up through personal projects and real-world experience.",
	Experience: []models.Experience {experience1, experience1,},
	Skills:     []string{"React", "Node.js", "Go", "Java", "Git", "GitHub", "Communication", "Problem Solving"},
	ImageURL : "../static/images/profil.jpeg",
}

func GetAboutData() models.AboutData {
	return dataAbout
}
