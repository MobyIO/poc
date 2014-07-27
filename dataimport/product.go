package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "strconv"
)

type ImportOrigin interface {
    importProducts()
}

type GoogleBooksImporter struct {
    endpoint    string
}

type JsonBooks struct {
    Kind        string              `json:"kind"`
    Items       []JsonBookItems       `json:"items"`
    TotalItems  int             `json:"totalItems"`
}

type JsonBookItems struct {
    Id          string      `json:"id"`
}

type JsonVolumeInfo struct {
    Title       string      `json:"title"`
    Description string      `json:"description"`
}

type JsonBook struct {
    Id          string      `json:"id"`
    ImageLinks  ImageLink   `json:"imageLinks"`
    VolumeInfo JsonVolumeInfo  `json:"volumeInfo"`
}

type ImageLink struct {
    SmallThumbnail      string      `json:"smallThumbnail"`
    Thumbnail           string      `json:"thumbnail"`
    Small               string      `json:"small"`
    Medium              string      `json:"medium"`
    Large               string      `json:"large"`
    ExtraLarge          string      `json:"extraLarge"`
}

func (r GoogleBooksImporter) importProducts() {
    query := "?q=children&maxResults=40&startIndex="
    totalItems := getTotalItems(r.endpoint + query + "0")
    pages := totalItems/40
    if(totalItems%40 > 0){
        pages = pages +1
    }
    for p:= 0; p <= pages; p++ {
        index := p*40
        url := r.endpoint + query + strconv.Itoa(index)
        fmt.Println(url)
        items := getItems(url)
        if(len(items) > 0){
            fmt.Println(len(items))
            for i:= 0; i < len(items); i++ {
                volumeId := items[i].Id
                fmt.Println(r.endpoint + volumeId)
                item := jsonBookApiGet(r.endpoint + volumeId)
                fmt.Println(item.VolumeInfo.Title)
            }
        }
    }
}

func jsonBooksApiGet(url string)*JsonBooks {
    resp, err := http.Get(url)
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
    return data
}

func getTotalItems(url string) int {
    data := jsonBooksApiGet(url)
    return data.TotalItems
}

func getItems(url string) []JsonBookItems {
    data := jsonBooksApiGet(url)
    return data.Items
}

func jsonBookApiGet(url string) *JsonBook {
    resp, err := http.Get(url)
    if(err != nil){
        panic(err.Error())
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if(err != nil){
        panic(err.Error())
    }
    data := &JsonBook{}
    err = json.Unmarshal(body, &data)
    if(err != nil){
        panic(err)
    }
    return data
}

func main() {
    g := GoogleBooksImporter{"https://www.googleapis.com/books/v1/volumes/"}
    importAll(g)
}

func importAll(resource ImportOrigin){
    resource.importProducts();
}
