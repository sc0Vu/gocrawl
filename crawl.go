package main

import (
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "log"
    "strings"
)

type Response struct {
    body string
}

var res Response

// only return true
// if failed program will stop
func crawl(method string, url string, data io.Reader, headers map[string]string) bool {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, data)

    if err != nil {
        log.Fatal(err)
    }
    if (headers != nil) {
        for key, value := range headers {
            req.Header.Add(key, value)
        }
    }

    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
    res.body = string(body[:])
    return true
}

func main() {
    headers := map[string]string{ "Origin":"https://www.google.com" }

    body := strings.NewReader("")

    isGet := crawl("GET", "https://www.google.com", sendbody, headers)

    if isGet {
        fmt.Println(res)
    }
}
