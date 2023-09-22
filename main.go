package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

const serverPort = 8080

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	GetApi()
}
func GetApi() {
    // url := "https://groupietrackers.herokuapp.com/api/artists"
    // req, err := http.NewRequest("GET", url, nil)
    // if err != nil {
    //     fmt.Print(err.Error())
    // }
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
    //     fmt.Print(err.Error())
    // }
	// defer res.Body.Close()
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}