package main
import(
    "net/http"
    "fmt"
)
func main() {
    url := "https://groupietrackers.herokuapp.com/api/artists"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Print(err.Error())
    }
	res, err := http.DefaultClient.Do(req)
	if err != nil {
        fmt.Print(err.Error())
    }
	defer res.Body.Close()
}