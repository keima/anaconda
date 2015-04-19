package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetUsersLookup(usernames string, v url.Values) (u []User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", usernames)
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/lookup.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUsersLookupByIds(ids []int64, v url.Values) (u []User, err error) {
	var pids string
	for w, i := range ids {
		//pids += strconv.Itoa(i)
		pids += strconv.FormatInt(i, 10)
		if w != len(ids)-1 {
			pids += ","
		}
	}
	v = cleanValues(v)
	v.Set("user_id", pids)
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/lookup.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUsersShow(username string, v url.Values) (u User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", username)
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/show.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUsersShowById(id int64, v url.Values) (u User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/show.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUserSearch(searchTerm string, v url.Values) (u []User, err error) {
	v = cleanValues(v)
	v.Set("q", searchTerm)
	// Set other values before calling this method:
	// page, count, include_entities
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/search.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) ReportSpamUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return a.ReportSpam(v)
}

func (a TwitterApi) ReportSpamUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return a.ReportSpam(v)
}

func (a TwitterApi) ReportSpam(v url.Values) (user User, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/users/report_spam.json", v, &user, _POST, response_ch}
	return user, (<-response_ch).err
}
