package instago

import "fmt"

//Gets details for media with the given ID
//
//mediaId: A string representing the ID of the media to get info on
func (api InstagramAPI) Media(mediaId string) Media {
	params := getEmptyMap()
	response := api.DoRequest("media/"+mediaId, params)
	return MediaFromAPI(response.Object("data"))
}

//Gets a list of popular media at the moment
func (api InstagramAPI) Popular() []Media {
	media, _, _ := api.GenericMediaListRequest("media/popular", "", "", 0)
	return media
}

//Gets a list of media posted from a certain location recently.
//N.B. This seems a bit unreliable...
//
//lat: The latitude to search near
//
//long: The longitude to search near
//
//distance: (optional = 0) The number of meters to search within
func (api InstagramAPI) LocationSearch(lat, lng, distance float64) []Media {
	//Unfortunately I couldn't use GenericMediaListRequest because it takes in location
	params := getEmptyMap()
	if distance > 0 {
		params["distance"] = fmt.Sprintf("%f", distance)
	}
	params["lat"] = fmt.Sprintf("%f", lat)
	params["lng"] = fmt.Sprintf("%f", lng)
	results := api.DoRequest("media/search", params)
	data := results.ObjectArray("data")
	media_objects := make([]Media, 0)
	for _, media := range data {
		media_objects = append(media_objects, MediaFromAPI(media))
	}
	return media_objects
}
