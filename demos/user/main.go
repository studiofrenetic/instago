package main

import (
	"fmt"
	"github.com/mbelousov/instago"
	"io/ioutil"
)

func main() {
	//Load the Client ID from a file called config.txt
	api := instago.InstagramAPI{}
	clientId, _ := ioutil.ReadFile("config.txt")
	api.ClientID = string(clientId)

	fmt.Println("INSTAGO  DEMO")
	fmt.Println("=============")
	fmt.Println("Enter a user:")

	var query string
	fmt.Scan(&query)

	//Search the users
	users, pagination, err := api.SearchUsers(query, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(users) <= 0 {
		fmt.Println("No results")
		return
	}
	for _, user := range users {
		fmt.Println("Username:", user.Username, "Full Name:", user.FullName)
	}
	fmt.Println(pagination)
	//Present basic inforamtion about the user
	fmt.Println("More detail on @" + users[0].Username)
	user, err := api.UserDetail(users[0].ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:", user.ID)
	fmt.Println("Username:", user.Username)
	fmt.Println("Full name:", user.FullName)
	fmt.Println("Bio:", user.Bio)
	fmt.Println("Website:", user.Website)
	fmt.Println("Follows:", user.TotalFollows)
	fmt.Println("Followers:", user.TotalFollowers)
	fmt.Println("Media:", user.TotalMedia)
}
