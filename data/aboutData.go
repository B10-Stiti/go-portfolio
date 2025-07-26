package data

import "portfolio/models"

var dataAbout = models.AboutData{
	AboutIntro: "Currently studying Computer Science at ISAMM.",
	About:      "Passionate about building things with code. I love solving problems and learning by doing — whether it's web apps, mobile tools, or anything in between. Right now, I'm focused on leveling up through personal projects and real-world experience.",
	Experience: []string{
		"Built a full-stack Node.js app with MongoDB and Mongoose. Includes full CRUD, user auth with JWT, and bcrypt security.",
		"Developed Pixamm, a mobile pixel art app using Java in Android Studio.",
	},
	Skills: []string{"React", "Node.js", "Go", "java", "Git", "GitHub", "Communication", "Problem Solving"},
}

func GetAboutData() models.AboutData {
	return dataAbout
}
