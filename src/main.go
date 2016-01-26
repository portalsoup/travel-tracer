package main
import (
	"net/http"
	"log"
	"jcleary/traveltracer/src/db"
	"jcleary/traveltracer/src/routes"
	"flag"
)

type Config struct {
	dbHost string
}

func main() {

	config := initFlags()

	dbStore := new(db.DataStore)

	err :=  dbStore.OpenSession(config.dbHost)
	if err != nil {
		panic(err)
	}

	defer db.Mgo.CloseSession()

	// Init the endpoints
	initHandlers()

	// Begin
	log.Println("About to listen and serve... ")
	http.ListenAndServe(":8284", nil)
	log.Println("About to terminate")
}

func initHandlers() {
	log.Print("Initializing handlers... ")
	http.HandleFunc("/map/point", routes.PointHandler)
	http.HandleFunc("/map/route", routes.RouteHandler)
	http.HandleFunc("/raw/coordinate", routes.RawCoordinateHandler)
	http.HandleFunc("/raw/route", routes.RawRouteHandler)
	log.Println("complete")
}

func initFlags() Config {
	var host = flag.String("db.host", "localhost", "define the database hostname")

	log.Println(*host)
	flag.Parse()

	return Config{dbHost: *host}
}