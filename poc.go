package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "flag"
)

type Product struct {
    Title           string  `json:"title"`
    Description     string  `json:"description"`
}

func listProducts( w http.ResponseWriter, r *http.Request){
    response, err := getProducts()
    if err != nil{
        panic(err)
    }

    fmt.Fprintf(w, string(response))
}

func main() {

    // command line flags
    dir := flag.String("directory", "app/", "angular-famous app")
    flag.Parse()

    // handle all requests by serving a file of the same name
    fs := http.Dir(*dir)
    fileHandler := http.FileServer(fs)
    http.Handle("/", fileHandler)

    // handle app api
    http.HandleFunc("/api/products/", listProducts)

    err := http.ListenAndServe(":8080", nil)
    fmt.Println(err.Error())
}

func getProducts()([]byte, error) {
    var products = make([]Product ,0)
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast"})
    products = append(products, Product{"Moby Dick","The narrator, Ishmael, is an observant young man setting out from Manhattan Island who has experience in the merchant marine but has recently decided his next voyage will be on a whaling ship. On a cold, gloomy night in December, he arrives at the Spouter-Inn in New Bedford, Massachusetts, and agrees to share a bed with a then-absent stranger. When his bunk mate, a heavily tattooed Polynesian harpooner named Queequeg, returns very late and discovers Ishmael beneath his covers, both men are alarmed, but the two quickly become close friends and decide to sail together from Nantucket, Massachusetts, on a whaling voyage."})
    return json.MarshalIndent(products, "", "  ")
}
