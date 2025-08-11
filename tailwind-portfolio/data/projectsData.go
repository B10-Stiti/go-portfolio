package data

import "portfolio/models"

var project1 = models.Project{
	Title:       "Mapping App",
	Description: "a Node.js application using Mongoose for MongoDB, featuring CRUD operations, user management, and secure authentication with JWT and bcrypt.",
	GithubRepo:  "https://github.com/B10-Stiti/Mapping-APP",
	ImageUrl:    "../static/images/proj_node.jpg",
}
var project2 = models.Project{
	Title:       "PIXAM",
	Description: "Pixel Art app built in Java, offering intuitive drawing tools, a customizable color palette, and reliable local storage powered by Room database",
	GithubRepo:  "https://github.com/Kefrov/PIXSAM",
	ImageUrl:    "../static/images/proj_pixam.png",
}
var project3 = models.Project{
	Title:       "Go portfolio",
	Description: "A personal portfolio website built with Go for the backend, utilizing HTML templates for server-side rendering and tailwind CSS for styling.",
	GithubRepo:  "https://github.com/B10-Stiti/go-portfolio",
	ImageUrl:    "../static/images/proj_portfolio.png",
}

var projects = []models.Project{project1, project2,project3}

func GetProjectsData() []models.Project {
	return projects
}
