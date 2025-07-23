package data

import "portfolio/models"

var data = models.HomeData{
	Title:     "Badis Portfolio",
	Name:      "Mohamed Badis Stiti",
	Tagline:   "Computer Science Student & Passionate Problem Solver",
	Intro:     "As a CS student, I'm fueled by a strong passion for problem-solving.I aim to improve my programming skills and explore different areas of the tech world to hone my abilities.",
	TechStack: []string{"Go", "React", "Node.js","Java"},
	ImageURL : "../static/images/profil.jpeg",
}

func GetHomeData() models.HomeData {
	return data
}
