package main

import (
	"fmt"
	"github.com/mbelousov/instago"
	"io/ioutil"
)

func displayUserList(users []instago.User) {
	for _, user := range users {
		fmt.Println("Username:", user.Username, "Full Name:", user.FullName)
	}
}
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
	displayUserList(users)
	fmt.Println(pagination)
	//Present basic inforamtion about the user
	fmt.Println("@" + users[0].Username + " followers: ")
	followers, pagination, err := api.UserFollowers(users[0].ID, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	displayUserList(followers)
	fmt.Println("@" + users[0].Username + " follows: ")
	follows, pagination, err := api.UserFollows(users[0].ID, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	displayUserList(follows)

}
