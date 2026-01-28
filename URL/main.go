package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("I want to learn concepts of URL in Golang")

	// ---------------- URL PARSING ----------------
	myurl := "https://www.example.com:8080/path/to/resource?query=golang#section1"

	parsedURL, err := url.ParseRequestURI(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	fmt.Println("Hostname:", parsedURL.Hostname())
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Is Absolute:", parsedURL.IsAbs())

	// ---------------- QUERY PARAMETERS ----------------
	queryParams := parsedURL.Query()
	for key, values := range queryParams {
		for _, value := range values {
			fmt.Printf("Query Param: %s = %s\n", key, value)
		}
	}

	// Modify query
	newQuery := url.Values{}
	newQuery.Add("search", "golang")
	newQuery.Add("page", "1")
	parsedURL.RawQuery = newQuery.Encode()

	fmt.Println("Modified URL:", parsedURL.String())

	// ---------------- BUILD CUSTOM URL ----------------
	customURL := &url.URL{
		Scheme:   "https",
		Host:     "api.example.com",
		Path:     "/v1/resources",
		RawQuery: "type=example&limit=10",
	}

	fmt.Println("Custom URL:", customURL.String())

	// ---------------- HTTP REQUEST ----------------
	client := &http.Client{}

	req, err := http.NewRequest("GET", customURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add custom header
	req.Header.Add("Custom-Header", "HeaderValue")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status Code:", resp.StatusCode)
	fmt.Println("Response Body:")
	fmt.Println(string(body))

	fmt.Println("File deleted successfully:")

	// Making a POST request
	postData := url.Values{}
	postData.Set("username", "testuser")
	postData.Set("password", "testpass")
	postResponse, err := http.PostForm(customURL.String(), postData)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer postResponse.Body.Close()

	postBody, err := io.ReadAll(postResponse.Body)
	if err != nil {
		fmt.Println("Error reading POST response body:", err)
		return
	}
	fmt.Println("POST Response Body:")
	fmt.Println(string(postBody))
	fmt.Println("POST Status Code:", postResponse.StatusCode)

	// Handling query parameters
	reqWithParams, err := http.NewRequest("GET", customURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating request with params:", err)
		return
	}
	q := reqWithParams.URL.Query()
	q.Add("userId", "1")
	reqWithParams.URL.RawQuery = q.Encode()
	clientWithParams := &http.Client{}
	respWithParams, err := clientWithParams.Do(reqWithParams)
	if err != nil {
		fmt.Println("Error making request with params:", err)
		return
	}
	defer respWithParams.Body.Close()

	bodyWithParams, err := io.ReadAll(respWithParams.Body)
	if err != nil {
		fmt.Println("Error reading response body with params:", err)
		return
	}
	fmt.Println("Response Body with Params:")
	fmt.Println(string(bodyWithParams))
}