package util

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"sync"
)

const Google_News_URL = "https://newsapi.org/v2/everything"

func AsyncHTTP(queries []string) (map[string]string, error) {
	ch := make(chan map[string]string)
	var wg sync.WaitGroup
	for _, query := range queries {
		wg.Add(1)
		go executeHttpRequest(query, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	responses := make(map[string]string)
	for mapString := range ch {
		for k, v := range mapString {
			responses[k] = v
			fmt.Println(mapString)
		}
	}
	return responses, nil
}

func executeHttpRequest(query string, ch chan<- map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	params := map[string]string{
		"q":      query,
		"from":   "2020-1-15",
		"sortBy": "publishedAt",
		"apiKey": "e4d1a5d882eb439ea2471a6d9948ac1c"}
	resp, err := GetResponse(Google_News_URL, params)
	if err != nil {
		log.Println("Error occured")
	}
	m := make(map[string]string)
	articles := gjson.Get(string(resp), `articles`)
	m[query] = articles.String()
	ch <- m

	// metricValue := gjson.Get(string(resp), `data.result.0.value.1`)
	// m := make(map[string]string)
	// m[metricName.String()] = metricValue.String()
	// ch <- m
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
