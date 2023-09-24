package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	// "html/template"
)

type Response struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func GetApi() []Response {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	// fmt.Println(response.Body)
	responseData, _ := ioutil.ReadAll(response.Body)

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
	return responseObjects
}

func BandsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call the GetApi function to fetch the data
	bands := GetApi()

	if err := tmpl.ExecuteTemplate(w, "index.html", bands); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func main() {
	// http.HandleFunc("/artists", BandsHandler)

	// fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", BandsHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
