package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)
func main() {
	fmt.Println("i want to learn consepts of web request in golang:")
	url := "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:")
	fmt.Println(string(body))

	// Print status code and headers
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Headers:", response.Header)

	// Making a POST request
	postURL := "https://jsonplaceholder.typicode.com/posts"
	postData := `{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`
	postResponse, err := http.Post(postURL, "application/json", ioutil.NopCloser(strings.NewReader(postData)))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer postResponse.Body.Close()

	postBody, err := ioutil.ReadAll(postResponse.Body)
	if err != nil {
		fmt.Println("Error reading POST response body:", err)
		return
	}
	fmt.Println("POST Response Body:")
	fmt.Println(string(postBody))

	fmt.Println("POST Status Code:", postResponse.StatusCode)

	// Custom headers in request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating new request:", err)
		return
	}
	req.Header.Set("Custom-Header", "CustomValue")
	customResp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making custom request:", err)
		return
	}
	defer customResp.Body.Close()

	customBody, err := ioutil.ReadAll(customResp.Body)
	if err != nil {
		fmt.Println("Error reading custom response body:", err)
		return
	}
	fmt.Println("Custom Response Body:")
	fmt.Println(string(customBody))
	fmt.Println("Custom Status Code:", customResp.StatusCode)

	// Handling query parameters
	reqWithParams, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request with params:", err)
		return
	}
	q := reqWithParams.URL.Query()
	q.Add("userId", "1")
	reqWithParams.URL.RawQuery = q.Encode()




	respWithParams, err := client.Do(reqWithParams)
	if err != nil {
		fmt.Println("Error making request with params:", err)
		return
	}
	defer respWithParams.Body.Close()

	bodyWithParams, err := ioutil.ReadAll(respWithParams.Body)
	if err != nil {
		fmt.Println("Error reading response body with params:", err)
		return
	}
	fmt.Println("Response Body with Params:")
	fmt.Println(string(bodyWithParams))


	fmt.Println("Status Code with Params:", respWithParams.StatusCode)
    
	// Note: For more advanced usage, consider using third-party libraries like "gorilla/mux" for routing and "resty" for simplified HTTP requests.
	// This example covers basic web requests using the standard library.
	// Always handle errors and edge cases in production code.
	
}