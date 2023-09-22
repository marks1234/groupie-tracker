package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	// "html/template"
)

const serverPort = 8080

type BandInfo struct {
	Id		int		  `json:"id"`
	Image 	string	  `json:"image"`
    Name    string    `json:"name"`
	Members []string  `json:"members"`
	CreationDate int  `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}



func GetApi() ([]BandInfo, error){
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
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var bands []BandInfo
    if err := json.Unmarshal(responseData, &bands); err != nil {
        return nil, err
    }

    return bands, nil
}

func main() {
	// doesn't work for now
	// bands, err := GetApi()
	// if err != nil {
    //     log.Fatal(err)
    // }

	//     // Define an HTTP handler to serve the existing HTML page with band information.
	// http.HandleFunc("/bands", func(w http.ResponseWriter, r *http.Request) {
	// 		// Parse the HTML template from the file.
	// 	tmpl, err := template.ParseFiles("index.html")
	// 		if err != nil {
	// 			log.Println(err)
	// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 			return
	// 		}
	
	// 		// Execute the template with the 'bands' data and write it to the response.
	// 	if err := tmpl.Execute(w, bands); err != nil {
	// 			log.Println(err)
	// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 			return
	// 		}
	// 	})

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}