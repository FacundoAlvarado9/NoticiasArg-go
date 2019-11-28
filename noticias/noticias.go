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

func main() {
	//Getting acces to API key from .env file
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

	//Parsing request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	newsJSON := string(body)

	//Sending it to an Array of news
	var entries Entrada
	var articles []Article

	json.Unmarshal([]byte(newsJSON), &entries)
	articles = entries.Articles

	//Prints title, description and url on console
	for i := 0; i < len(articles); i++ {
		linea := "Titulo: " + articles[i].Title + " \n " + "Desc: " + articles[i].Description + "\n Link: " + articles[i].Url + " \n \n"
		fmt.Printf(linea)
	}
}
