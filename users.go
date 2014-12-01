package instago

import "fmt"

//Gets basic information about a given user
//
//userID: a string representing the ID (not the username) of a given user
func (api InstagramAPI) UserDetail(userID string) (User, error) {
	params := getEmptyMap()
	result := api.DoRequest("users/"+userID, params)
	data := result.Object("data")
	return UserFromAPI(data), api.ErrorFromAPI(result)
}

//Query the users on Instagram and get a list of them back
//
//query: The description such as 'jack' or 'thomas' to search for
//
//max: (optional, default = 0) the number of users to return
func (api InstagramAPI) SearchUsers(query string, max int) ([]User, Pagination, error) {
	params := getEmptyMap()
	params["q"] = query
	if max > 0 {
		params["count"] = fmt.Sprintf("%d", max)
	}
	result := api.DoRequest("users/search", params)
	data := result.ObjectArray("data")
	users := make([]User, 0)
	for _, user := range data {
		users = append(users, UserFromAPI(user))
	}
	pagination := PaginationFromAPI(result.Object("pagination"))
	err := api.ErrorFromAPI(result)

	return users, pagination, err
}

//Will return an array of recently posted media objects by a user. Requires OAuth
//
//userId: string representing the user
//
//max: the greatest number of media objects to return
//
//before: (optional = "") posts before a certain ID
//
//after: (optional = "") posts after a certain ID
func (api InstagramAPI) RecentPostsByUser(userId string, max int, before, after string) ([]Media, Pagination, error) {
	return api.GenericMediaListRequest("users/"+userId+"/media/recent", before, after, max)
}

//Gets the current user's feed (requires OAuth)
//
//before: (optional = "") posts before a certain ID
//
//after: (optional = "") posts after a certain ID
//
//max: (optional = 0) the greatest number of media objects to return
func (api InstagramAPI) Feed(before, after string, max int) ([]Media, Pagination, error) {
	return api.GenericMediaListRequest("users/self/feed", before, after, max)
}

//Gets the posts like by the current user (requires OAuth)
//
//max: (optional = 0) the greatest number of posts to return
//
//before: (optional = 0) posts liked before a certain media ID
func (api InstagramAPI) Liked(max int, before string) ([]Media, Pagination, error) {
	return api.GenericMediaListRequest("users/self/media/liked", before, "", max)
}
