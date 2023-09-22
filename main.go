package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	// "html/template"
)

const serverPort = 8080

type Response struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func GetApi() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	// fmt.Println(response.Body)
	responseData, err := ioutil.ReadAll(response.Body)

	// var bands []BandInfo = string(responseData)
	var responseObjects []Response

	re := regexp.MustCompile(`{.+?}`)
	str_arr := re.FindAllString(string(responseData), -1)
	for _, str := range str_arr {
		var responseObject Response
		json.Unmarshal([]byte(str), &responseObject)
		responseObjects = append(responseObjects, responseObject)
	}
	fmt.Println(responseObjects)
}

func main() {
	GetApi()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
