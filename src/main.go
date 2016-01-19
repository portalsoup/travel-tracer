package main
import (
	"html/template"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"regexp"
	//"errors"
	"strconv"
	//"encoding/json"
)

// templates is a cache for html templates.
var (
	route = template.Must(template.ParseFiles(
		"../views/route-map.html",
	))

	point = template.Must(template.ParseFiles(
		"../views/point-map.html",
	))
	//validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

// Coordinate represents a single point on earth using latitude and longitude.
type Coordinate struct {
	Latitude float64
	Longitude float64
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

	lat, err := strconv.ParseFloat(r.FormValue("latitude"), 64)

	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(r.FormValue("longitude"), 64)

	if err != nil {
		return nil, err
	}

	c = &Coordinate{Latitude: lat, Longitude: lng}

	fmt.Printf("Lat: %f Lng: %f\n", c.Latitude, c.Longitude)
	return c, nil
}

// pointHandler handles the point-map endpoint and renders the appropriate template.
func pointHandler(w http.ResponseWriter, r *http.Request) {
	coord, _ := getCoordinates(r)

	err := point.ExecuteTemplate(w, "point-map.html", coord)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeHandler(w http.ResponseWriter, r *http.Request) {

	coords := []Coordinate{
		{Latitude: 44.053, Longitude: -123.091},
		{Latitude: 37.772, Longitude: -122.214},
		{Latitude: 21.291, Longitude: -157.821},
		{Latitude: -18.14, Longitude: 178.431},
		{Latitude: -29.467, Longitude: 153.027},
		{Latitude: 23.467, Longitude: 74.848},
	}

	fmt.Println(coords)

	err := route.ExecuteTemplate(w, "route-map.html", coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/map/point", pointHandler)
	http.HandleFunc("/map/route", routeHandler)
	http.ListenAndServe(":8080", nil)
}