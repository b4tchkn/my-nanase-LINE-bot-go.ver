package main

import (
	//"fmt"
	"encoding/json"
	"net/url"
	"os"
	//"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

const ENDPOINT = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"

type Results struct {
	Perplexity float64 `json:"perplexity"`
	Reply string `json:"reply"`
}

type Responce struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Results []Results `json:"results"`
}

func A3rt(query string) (string, error) {
	apikey := os.Getenv("APIKEY")
	values := url.Values{}
	values.Add("apikey", apikey)
	values.Add("query", query)


	// http.PostForm(ENDPOINT, values)の中身
	// https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk?apikey=apikey&query=query
	res, err := http.PostForm(ENDPOINT, values)

	if err != nil {
		log.Fatal("*PostFrom*\n", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("*ReadAll*\n", err)
	}
	jsonBytes := ([]byte)(body)
	responce := new(Responce)
	err = json.Unmarshal(jsonBytes, responce)

	if err := json.Unmarshal(jsonBytes, responce); err != nil {
		log.Fatal("*json.Unmarshal*\n", err)
	}
	return responce.Results[0].Reply, err
}