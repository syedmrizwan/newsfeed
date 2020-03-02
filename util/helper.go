package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// NewsURL is the the URL for newsapi
const NewsURL = "https://newsapi.org/v2/everything"

// AsyncHTTP executes http request in an asynchronous manner
// Signal With Data - No Guarantee - Buffered Channels >1
// Fan Out
func AsyncHTTP(queries []string) (map[string][]string, error) {
	ch := make(chan map[string][]string, len(queries))
	defer close(ch)
	for _, query := range queries {
		go func(query string, ch chan map[string][]string) {
			executeHTTPRequest(query, ch)
		}(query, ch)
	}

	responses := make(map[string][]string)

	for i := 0; i < len(queries); i++ {
		mapString := <-ch
		for k, v := range mapString {
			responses[k] = v
		}
	}
	return responses, nil
}

func executeHTTPRequest(query string, ch chan map[string][]string) {
	params := map[string]string{
		"q":      query,
		"from":   "2020-2-10",
		"sortBy": "publishedAt",
		"apiKey": "e4d1a5d882eb439ea2471a6d9948ac1c"}
	resp, err := getResponse(NewsURL, params)
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
	select {
	case ch <- m:
		fmt.Println("manager : send ack")
	default:
		fmt.Println("manager : drop")
	}

}

func getResponse(url string, params map[string]string) ([]byte, error) {
	client := resty.New()
	client.SetQueryParams(params)
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp.Body(), nil
}

// GetURLResponse takes a query as parameter and returns its response by using newsapi URL
func GetURLResponse(ctx context.Context, query string) ([]byte, error) {
	u, _ := url.Parse(NewsURL)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("q", query)
	q.Add("from", "2020-1-20")
	q.Add("sortBy", "publishedAt")
	q.Add("apiKey", "e4d1a5d882eb439ea2471a6d9948ac1c")
	u.RawQuery = q.Encode()

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	// If the request failed, log to STDOUT
	if err != nil {
		fmt.Println("Request failed:", err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return b, nil
}
