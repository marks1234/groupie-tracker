package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	// "html/template"
)

type LocationDate struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
} //trying to access the relation api

type Response struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var (
	tmpl        = template.Must(template.ParseFiles("index.html"))
	tmpl_gotcha = template.Must(template.ParseFiles("gotcha.html"))
)

func GetApi() []Response {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	// fmt.Println(response.Body)

	// var bands []BandInfo = string(responseData)
	var responseObjects []Response
	if err := json.NewDecoder(response.Body).Decode(&responseObjects); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// fmt.Println(responseObjects)
	return responseObjects
}

// this is where I stopped, I wanted to try to put both calls in the same function :/
func GetApiRelation() []LocationDate {
	relation, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer relation.Body.Close()
	// fmt.Println(response.Body)

	// var bands []BandInfo = string(responseData)
	var relationObjects []LocationDate
	if err := json.NewDecoder(relation.Body).Decode(&relationObjects); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// fmt.Println(relationObjects)
	return relationObjects
}

func BandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call the GetApi function to fetch the data
	bands := GetApi()

	if err := tmpl.ExecuteTemplate(w, "index.html", bands); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func PrivateHandler(w http.ResponseWriter, r *http.Request, res Response) {
	if err := tmpl_gotcha.ExecuteTemplate(w, "gotcha.html", []Response{res}); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func main() {
	// http.HandleFunc("/artists", BandsHandler)

	// fileServer := http.FileServer(http.Dir("./static"))
	// fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", BandsHandler)
	for _, res := range GetApi() {
		res := res
		http.HandleFunc(fmt.Sprint("/", res.Name), func(w http.ResponseWriter, r *http.Request) {
			PrivateHandler(w, r, res)
		})
	}

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
