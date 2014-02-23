package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io"
)

type SearchD struct {
	Statuses        []Status
	Search_metadata interface{}
}

type Status struct {
	Entities Entity
}

type Entity struct {
	Urls []UrlData
}

type UrlData struct {
	//Url          string
	Expanded_url string
}

var TWITTER = oauth.NewConsumer(
	"",
	"",//enter your key
	oauth.ServiceProvider{
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	},
)

func Search() []Status {
  
  additionalData := map[string]string{
		oauth.SESSION_HANDLE_PARAM: "SESSION_HANDLE",
	}
  
	token := &oauth.AccessToken{"", "",additionalData}//enter yours
	resp, err := TWITTER.Get(
		"https://api.twitter.com/1.1/search/tweets.json",
		map[string]string{
			"q":     "#prfm",
			"lang":  "ja",
			"count": "20"},
		token)
	if err != nil {
		fmt.Println("Error: could not connet twitter")
		return nil
	}
	defer resp.Body.Close()

	m := SearchD{}
	dec := json.NewDecoder(resp.Body)
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	//return m.Statuses.Enti
	// for _, v := range m.Statuses {
	// 	for _, vv := range v.Entities.Urls {
	// 		result = append(result, vv.Expanded_url)
	// 		fmt.Println("url:%v", vv.Expanded_url)
	// 	}
	// }
	return m.Statuses
}
