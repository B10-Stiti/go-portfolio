package data

import "portfolio/models"

var dataAbout = models.AboutData{
	AboutIntro: "Currently studying Computer Science at ISAMM.",
	About: "Passionate about building things with code. I love solving problems and learning by doing — whether it's web apps, mobile tools, or anything in between. Right now, I'm focused on leveling up through personal projects and real-world experience.",
	Education:"Finished my second year of a Computer Science degree at ISAMM. Currently working on side projects to sharpen my skills and prepare for my final year project.",
	Experience: []string{
		"Built a full-stack Node.js app with MongoDB and Mongoose. Includes full CRUD, user auth with JWT, and bcrypt security.",
		"Developed Pixamm, a mobile pixel art app using Java in Android Studio.",
	},
	Skills: []string{"React", "Node.js", "Golang", "Git", "GitHub","Communication","Problem Solving"},
}

func GetAboutData() models.AboutData {
	return dataAbout
}
