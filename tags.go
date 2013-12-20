package instago

import (
	"fmt"
)

//Gets all (16) recent photos with the given tag
//
//tag: The tag (don't include the # hash) that you want to fetch
//
//before: (optional - use "") find photos posted before this ID (use Image.ID)
//
//after: (optional - use "") find photos posted after this ID (use Image.ID)
func (api InstagramAPI) TagRecent(tag, before, after string, max int) ([]Image, Pagination) {
	// return api.GenericImageListRequest("tags/"+tag+"/media/recent", before, after, 0)
	// var max int
	params := getEmptyMap()
	if max > 0 {
		params["count"] = fmt.Sprintf("%d", max)
	}
	if before != "" {
		params["max_tag_id"] = before
	}
	if after != "" {
		params["min_tag_id"] = after
	}

	results := api.DoRequest("tags/"+tag+"/media/recent", params)
	data := results.ObjectArray("data")
	images := make([]Image, 0)
	for _, image := range data {
		images = append(images, ImageFromAPI(image))
	}

	// log.New(os.Stderr, "ERROR ", results)

	// pagination := Pagination{}
	pagination := PaginationFromAPI(results.Object("pagination"))
	// fmt.Println("Pagination:", pagination)

	return images, pagination
	// return images
}

//Gets the total number of images on Instagram with a given tag
//
//tag: a string that represents the tag that you want to search for
func (api InstagramAPI) TagInfo(tag string) Tag {
	params := getEmptyMap()
	result := api.DoRequest("tags/"+tag, params)
	return tagObject(result.Object("data"))
}

//Will fetch the tag along with similar tags from Instagram so you can see the number of
//images with that tag
//
//tag: a string that represents the tag you want to search for
func (api InstagramAPI) TagSearch(tag string) []Tag {
	params := getEmptyMap()
	params["q"] = tag
	result := api.DoRequest("tags/search", params)
	tags := make([]Tag, 0)
	for _, tag := range result.ObjectArray("data") {
		tags = append(tags, tagObject(tag))
	}
	return tags
}

//Both TagInfo and TagSearch need to create Tag objects
func tagObject(json JSON) Tag {
	return Tag{Tag: json.String("name"), MediaCount: json.Int("media_count")}
}
