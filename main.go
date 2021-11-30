package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"background/models"

	"github.com/reujab/wallpaper"
)

func main() {

	const Secret = ""
	const ClientId = ""
	const UserAgent = ""

	//Get Authorization Token
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "")
	data.Set("password", "")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Set("User-Agent", UserAgent)
	req.SetBasicAuth(ClientId, Secret)

	fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error")
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Response OK")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		respString := string(bodyBytes)
		resp.Body.Close()
		log.Print(respString)

		var redditResp models.RedditResponse

		if err = json.Unmarshal([]byte(respString), &redditResp); err != nil {
			fmt.Println("failed to unmarshal: ", err)
		}

		//Get JSON from r/anime/stuff
		data := url.Values{}
		data.Set("after", "")
		data.Set("before", "")
		data.Set("count", "5")
		data.Set("limit", "25")

		req, err = http.NewRequest(http.MethodGet, "https://oauth.reddit.com/r/Animewallpaper/new", strings.NewReader(url.Values.Encode(data)))
		if err != nil {
			fmt.Println("Error")
		}
		req.BasicAuth()
		req.Header.Set("Authorization", redditResp.Token_type+" "+redditResp.Access_token)
		req.Header.Set("User-Agent", UserAgent)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//req.Host = "http://oauth.reddit.com"
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Error on second request")
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Response OK")
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			respString := string(bodyBytes)
			resp.Body.Close()

			var listing models.Listing

			if err = json.Unmarshal([]byte(respString), &listing); err != nil {
				fmt.Println("failed to unmarshal: ", err)
			}
			fmt.Println("success")
			if len(listing.Data.Children) > 0 {
				//SOME CHILDREN DO NOT HAVE NULL CHILDREN
				for i := 0; i < len(listing.Data.Children); i++ {
					if len(listing.Data.Children[i].Data.Preview.Images) > 0 {
						fmt.Println(listing.Data.Children[i].Data.UrlOverriddenByDest)
						err = wallpaper.SetFromURL(listing.Data.Children[i].Data.UrlOverriddenByDest)

					}
				}
			}
		} else {
			fmt.Println("Error on second request", resp.Status)
		}

	} else {
		fmt.Println(resp.StatusCode)
	}

}

//https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example
//https://medium.com/@fsufitch/deserializing-json-in-go-a-tutorial-d042412958ea

//https://www.sohamkamani.com/golang/json/
