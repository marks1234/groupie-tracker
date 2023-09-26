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
	DatesLocations map[string][]string `json:"datesLocations"`
} // trying to access the relation api

type Response struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
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

// I rewrote the function getLocation to a struct method
func (r *Response) getLocation() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/relation/" + fmt.Sprint(r.Id))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	var loca LocationDate
	if err := json.NewDecoder(response.Body).Decode(&loca); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	r.DatesLocations = loca.DatesLocations
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
	res.getLocation()
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
