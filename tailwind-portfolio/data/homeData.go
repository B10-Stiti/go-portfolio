package data

import "portfolio/models"

var dataHome = models.HomeData{
	Name:      "Mohamed Badis Stiti",
	Tagline:   "Computer Science Student",
	Intro:     "CS student who genuinely enjoys coding, solving real problems, and building things that make a difference.",
	TechStack: []string{"Go", "React", "Node.js"},
	ImageURL : "../static/images/profil.jpeg",
}

func GetHomeData() models.HomeData {
	return dataHome
}
