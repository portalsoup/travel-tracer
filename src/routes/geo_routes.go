package routes
import (
	"net/http"
	"log"
	"fmt"
	"jcleary/traveltracer/src/geo"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"html/template"
	"strconv"
)

var (
	route = template.Must(template.ParseFiles(
		"../views/route-map.html",
	))

	point = template.Must(template.ParseFiles(
		"../views/point-map.html",
	))
)

// pointHandler handles the /map/point endpoint and renders the appropriate template.
func PointHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered the point handler")
	coord, err := getLatLngParams(r)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Unexpected query parameters.", http.StatusInternalServerError)
		return
	}

	err = point.ExecuteTemplate(w, "point-map.html", coord)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// routeHandler handles the /map/route endpoint and renders the appropriate template.
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered the route handler")

	coords := []geo.Coordinate{
		{Latitude: 44.053, Longitude: -123.091, Id: bson.NewObjectId()},
		{Latitude: 37.772, Longitude: -122.214, Id: bson.NewObjectId()},
		{Latitude: 21.291, Longitude: -157.821, Id: bson.NewObjectId()},
		{Latitude: -18.14, Longitude: 178.431, Id: bson.NewObjectId()},
		{Latitude: -29.467, Longitude: 153.027, Id: bson.NewObjectId()},
		{Latitude: 23.467, Longitude: 74.848, Id: bson.NewObjectId()},
	}

	fmt.Println(coords)

	err := route.ExecuteTemplate(w, "route-map.html", coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RawCoordinateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered the raw coordinate handler")

	coord := &geo.Coordinate{Latitude: 44.053, Longitude: -123.091, Id: bson.NewObjectId()}

	b, err := json.Marshal(coord)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(b))

}

func RawRouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered the raw route handler")

	coords := []geo.Coordinate{
		{Latitude: 44.053, Longitude: -123.091, Id: bson.NewObjectId()},
		{Latitude: 37.772, Longitude: -122.214, Id: bson.NewObjectId()},
		{Latitude: 21.291, Longitude: -157.821, Id: bson.NewObjectId()},
		{Latitude: -18.14, Longitude: 178.431, Id: bson.NewObjectId()},
		{Latitude: -29.467, Longitude: 153.027, Id: bson.NewObjectId()},
		{Latitude: 23.467, Longitude: 74.848, Id: bson.NewObjectId()},
	}

	route := geo.Route{
		Id: bson.NewObjectId(),
		Coordinates: coords,
	}


	b, err := json.Marshal(route)

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(b))
}


// getLatLngParams will parse out the coordinate query values from the incoming request.
func getLatLngParams(r *http.Request) (c *geo.Coordinate, err error) {
	log.Println("Entered the latlng extractor")

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

	c = &geo.Coordinate{Latitude: lat, Longitude: lng, Id: bson.NewObjectId()}

	fmt.Printf("Lat: %f Lng: %f\n", c.Latitude, c.Longitude)
	return c, nil

}