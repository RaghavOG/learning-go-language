package main

import (
	"fmt"
	"strconv"
)

func main() {
	users := []User{
		{
			Name:   "Raghav",
			Age:    21,
			Skills: []string{"Go", "APIs", "CLI"},
			SocialLinks: map[string]string{
				"github":   "https://github.com/raghav",
				"linkedin": "https://linkedin.com/in/raghav",
			},
		},
		{
			Name:   "Aditya",
			Age:    23,
			Skills: []string{"Frontend", "React", "TypeScript"},
			SocialLinks: map[string]string{
				"github":   "https://github.com/aditya",
				"linkedin": "https://linkedin.com/in/aditya",
			},
		},
		{
			Name:   "Kabir",
			Age:    25,
			Skills: []string{"DevOps", "Docker", "CI/CD"},
			SocialLinks: map[string]string{
				"github":   "https://github.com/kabir",
				"linkedin": "https://linkedin.com/in/kabir",
			},
		},
	}

	// Add another user using AddUser
	newUser := User{
		Name:   "Meera",
		Age:    22,
		Skills: []string{"Data", "SQL", "Python"},
		SocialLinks: map[string]string{
			"github":   "https://github.com/meera",
			"linkedin": "https://linkedin.com/in/meera",
		},
	}
	users = AddUser(users, newUser)

	ListUsers(users)
}

type User struct {
	Name        string
	Age         int
	Skills      []string
	SocialLinks map[string]string
}

func (u User) Introduce() string {
	return "Hi, I'm " + u.Name + ", " + strconv.Itoa(u.Age) + " years old and my skills are " + fmt.Sprint(u.Skills)
}

// AddUser appends a user and returns the updated slice.
func AddUser(users []User, user User) []User {
	return append(users, user)
}

// ListUsers prints each user's introduction and social links.
func ListUsers(users []User) {
	for _, u := range users {
		fmt.Println(u.Introduce())
		fmt.Println("  GitHub:", u.SocialLinks["github"])
		fmt.Println("  LinkedIn:", u.SocialLinks["linkedin"])
	}
}