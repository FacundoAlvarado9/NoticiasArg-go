/**
	This script retrieves Google News' headlines
	related to Argentina. Using newsapi.org
**/
package main

import (
	"github.com/joho/godotenv"
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Article model
type Article struct {
	Author      string
	Title       string
	Description string
	Url         string
	publishedAt string
}

//Output of the Json parse model
type Entrada struct {
	Status       string
	TotalResults int
	Articles     []Article
}

// Parses the news and returns them as an array of Articles
func parseNews() []Article {
	//Getting access to API key from .env file
	err := godotenv.Load()
	if(err != nil){
		log.Fatal("Error loading env variables")
	}

	newsApi_key := os.Getenv("NEWSAPI_KEY")

	//GET request to the API site
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=google-news-ar&apiKey=" + newsApi_key)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//Parsing request body and storing it in a string
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	newsJsonString := string(body)

	//Getting array of news from body string
	var entries Entrada

	json.Unmarshal([]byte(newsJsonString), &entries)
	return entries.Articles
}

func handler(w http.ResponseWriter, r *http.Request){
	var articles []Article
	articles = parseNews()

	//Prints title, description and url on console
	for i := 0; i < len(articles); i++ {
		// Taking the apparently common <li> </li> <ol> from the descriptions
		if (strings.Contains(articles[i].Description, "<li>")){
			articles[i].Description = strings.ReplaceAll(articles[i].Description, "<li>", "")
		}
		if (strings.Contains(articles[i].Description, "</li>")){
			articles[i].Description = strings.ReplaceAll(articles[i].Description, "</li>", "")
		}
		if (strings.Contains(articles[i].Description, "<ol>")){
			articles[i].Description = strings.ReplaceAll(articles[i].Description, "<ol>", "")
		}

		// Printing them on screen
		linea := "Titulo: " + articles[i].Title + "\n " + "Desc: " + articles[i].Description + "\n Link: <a href=\"" + articles[i].Url + "\"> link </a> \n \n"
		fmt.Fprintf(w, "<html>" + linea + "</html>")
	}
}

func main() {
	http.HandleFunc("/noticias", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))	
}
