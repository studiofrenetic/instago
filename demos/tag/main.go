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
	fmt.Println("Enter a tag:")

	var tag string
	fmt.Scan(&tag)

	tagInfo := api.TagInfo(tag)
	fmt.Println("Tag: ", tagInfo.Tag, "Total: ", tagInfo.MediaCount)

	images, pagination, err := api.TagRecent(tag, "", "", 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, img := range images {
		fmt.Println(img.User, img.Filter)
	}
	fmt.Println(pagination)

	fmt.Println("Similar tags")
	tags, pagination, err := api.TagSearch(tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tag := range tags {
		fmt.Println("Tag: ", tag.Tag, "Total: ", tag.MediaCount)
	}

	fmt.Println("Pagination :", pagination)
}
