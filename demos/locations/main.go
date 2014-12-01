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

	fmt.Println("   INSTAGO  DEMO   ")
	fmt.Println("===================")
	fmt.Println("Posts at Instagram:")

	//Instagram HQ
	imagesInstagram, pagination, err := api.LocationPosts("514276", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range imagesInstagram {
		fmt.Println("User:", image.User, "Filter:", image.Filter, "Likes:", image.Likes)
	}
	fmt.Println(pagination)
	fmt.Println("===============================")
	fmt.Println("Locations near the Eiffel Tower")

	//Locations near the Eiffel Tower
	locationsInParis, pagination, err := api.LocationsNear(48.858844, 2.294351, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, loc := range locationsInParis {
		fmt.Println("Name:", loc.Name, "Coords:", loc.Latitude, loc.Longitude)
	}
	fmt.Println(pagination)
}
