package main

import (
	"fmt"
	"github.com/viclm/maple"
)

type dongqiudi struct {
	Articles []struct {
		Title string `json:title`
	} `json:articles`
}

func main() {
	options := maple.Options{
		Url:    "https://api.dongqiudi.com/app/tabs/iphone/1.json",
		Method: "GET",
	}

	var result dongqiudi

	if err := maple.GetJSON(options, &result); err != nil {
		switch err.(type) {
		case *maple.HTTPError:
			fmt.Println(err, err.(*maple.HTTPError).Code)
		default:
			fmt.Println(err)
		}
		return
	}

	for _, article := range result.Articles {
		fmt.Println(article.Title)
	}

}
