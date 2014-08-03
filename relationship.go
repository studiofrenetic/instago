package instago

//Get the list of users this user is followed by.
//
//userID: a string representing the ID (not the username) of a given user
func (api InstagramAPI) UserFollows(userID, cursor string) ([]User, Pagination) {
  params := getEmptyMap()
  if cursor != "" {
    params["cursor"] = cursor
  }
  result := api.DoRequest("users/"+userID+"/follows", params)
  data := result.ObjectArray("data")
  users := make([]User, 0)

  for _, user := range data {
    u := UserFromAPI(user)
    users = append(users, u)
  }
  pagination := PaginationFromAPI(result.Object("pagination"))
  return users, pagination
}

//Get the list of
//
//userID: a string representing the ID (not the username) of a given user
func (api InstagramAPI) UserFollowers(userID, cursor string) ([]User, Pagination) {
  params := getEmptyMap()
  if cursor != "" {
    params["cursor"] = cursor
  }
  result := api.DoRequest("users/"+userID+"/followed-by", params)
  data := result.ObjectArray("data")
  users := make([]User, 0)

  for _, user := range data {
    u := UserFromAPI(user)
    users = append(users, u)
  }
  pagination := PaginationFromAPI(result.Object("pagination"))
  return users, pagination
}
