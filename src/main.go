package main
import (
	"html/template"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"regexp"
	//"errors"
)

// templates is a cache for html templates.
var (
	templates = template.Must(template.ParseFiles("../views/point-map.html"))
	//validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

// Coordinate represents a single point on earth using latitude and longitude.
type Coordinate struct {
	Latitude string
	Longitude string
}

// Query represents a single url query.
type Query struct {
	Key string
	Value string
}

// getCoordinates will parse out the coordinate query values from the incoming request.
func getCoordinates(r *http.Request) (c *Coordinate, err error) {
	queries := r.URL.RawQuery
	if queries == "" {
		fmt.Println("No queries found!")
	}

	c = &Coordinate{Latitude: r.FormValue("latitude"), Longitude: r.FormValue("longitude")}

	if c.Latitude == "" || c.Longitude == "" {
		fmt.Printf("There was a problem finding the lat [%s] or lng [%s]\n", c.Latitude, c.Longitude)
		return nil, err
	}
	fmt.Printf("Lat: %s Lng: %s\n", c.Latitude, c.Longitude)
	return c, nil
}

// renderTemplate searches and substitutes arguments in a template with real values.
func renderTemplate(w http.ResponseWriter, tmpl string, coords *Coordinate) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// makeHandler to be honest I forget how this works, this was an "optimization" that a tutorial did that seems to work magic.
func makeHandler(fn func (http.ResponseWriter, *http.Request, string, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
//		m := validPath.FindStringSubmatch(r.URL.Path)
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
		fn(w, r, "", "")
	}
}

// pointHandler handles the point-map endpoint and renders the appropriate template.
func pointHandler(w http.ResponseWriter, r *http.Request, latitude string, longitude string) {
	coord, _ := getCoordinates(r)

	renderTemplate(w, "point-map", coord)
}

func main() {
	http.HandleFunc("/point-map", makeHandler(pointHandler))
	http.ListenAndServe(":8080", nil)
}