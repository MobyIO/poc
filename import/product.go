package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type ImportOrigin interface {
    importProducts()
}

type GoogleBooksImporter struct {
    endpoint string
}

type JsonBooks struct {
    Kind    string              `json:"kind"`
    Items   []JsonBookItems       `json:"items"`
}

type JsonBookItems struct {
    Id      string      `json:"id"`
}
func (r GoogleBooksImporter) importProducts() {
    query := "?q=children"
    fmt.Println("hello google books from endpoint " + r.endpoint + query)
    resp, err := http.Get(r.endpoint + query);
    if(err != nil){
        panic(err.Error())
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if(err != nil){
        panic(err.Error())
    }
    data := &JsonBooks{}
    err = json.Unmarshal(body, &data)
    if(err != nil){
        panic(err)
    }
    for i:= 0; i < len(data.Items); i++ {
        fmt.Println(data.Items[i].Id)
    }
}

type Message struct {
    Name string
    //Body string
    Time int64
}

func main() {
    g := GoogleBooksImporter{"https://www.googleapis.com/books/v1/volumes"}
    importAll(g)
}

func importAll(resource ImportOrigin){
    resource.importProducts();
}
