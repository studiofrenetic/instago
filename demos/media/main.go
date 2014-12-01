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

	fmt.Println("INSTAGO DEMO")
	fmt.Println("============")
	fmt.Println("Popular ATM:")

	images, pagination, err := api.Popular()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range images {
		fmt.Println("User:", image.User, "Filter:", image.Filter, "Likes:", image.Likes)
	}
	fmt.Println(pagination)
	fmt.Println("====================")
	fmt.Println("Posted in Manhattan:")

	//I damn hope I got these right!
	imagesNY, pagination, err := api.LocationSearch(40.7142, -74.0064, 4500)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range imagesNY {
		fmt.Println("User:", image.User, "Location:", image.Location.Name, "Coords:", image.Location.Latitude, image.Location.Longitude)
	}
	fmt.Println(pagination)
}
