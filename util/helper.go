package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

const NewsURL = "https://newsapi.org/v2/everything"

func AsyncHTTP(queries []string) (map[string][]string, error) {
	ch := make(chan map[string][]string)
	var wg sync.WaitGroup
	for _, query := range queries {
		wg.Add(1)
		go executeHttpRequest(query, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	responses := make(map[string][]string)
	for mapString := range ch {
		for k, v := range mapString {
			responses[k] = v
			fmt.Println(mapString)
		}
	}
	return responses, nil
}

func executeHttpRequest(query string, ch chan<- map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	params := map[string]string{
		"q":      query,
		"from":   "2020-1-15",
		"sortBy": "publishedAt",
		"apiKey": "e4d1a5d882eb439ea2471a6d9948ac1c"}
	resp, err := GetResponse(NewsURL, params)
	if err != nil {
		log.Println("Error occured")
	}
	var articleTitles []string
	m := make(map[string][]string)
	articles := gjson.Get(string(resp), `articles.#.title`)
	for _, name := range articles.Array() {
		articleTitles = append(articleTitles, name.String())
	}
	m[query] = articleTitles
	ch <- m
}

func GetResponse(url string, params map[string]string) ([]byte, error) {
	client := resty.New()
	client.SetQueryParams(params)
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp.Body(), nil
}

func GetURLResponse(query string) ([]byte, error) {
	u, _ := url.Parse(NewsURL)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("q", query)
	q.Add("from", "2020-1-20")
	q.Add("sortBy", "publishedAt")
	q.Add("apiKey", "e4d1a5d882eb439ea2471a6d9948ac1c")
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return b, nil
}
