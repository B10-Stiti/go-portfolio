package data

import "portfolio/models"

var dataContact = models.ContactData {
	ContactIntro: "Got a project, opportunity, or question?, Let’s connect — I’d love to hear from you.",
	SocialLinks: map[string]string{
    "GitHub": "https://github.com/B10-Stiti",
    "LinkedIn": "https://www.linkedin.com/in/be10",
	"Instagram": "https://www.instagram.com/stiti_badis",
},
}

func GetContactData() models.ContactData {
	return dataContact
}