package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "flag"
)

type Thumbnail struct {
    Small       string      `json:"small"`
    Medium      string      `json:"medium"`
    Large       string      `json:"large"`
    XLarge      string      `json:"xLarge"`
}

type Product struct {
    Title           string  `json:"title"`
    Description     string  `json:"description"`
    Thumbnails      Thumbnail `json:"thumbnail"`
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
    products = append(products, Product{"one World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})


    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})
    products = append(products, Product{"World War Z","Book about Zombie that runs really fast",Thumbnail{"http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api","http://bks1.books.google.com/books?id=5kJOuuykL24C&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api","",""}})


    return json.MarshalIndent(products, "", "  ")
}
