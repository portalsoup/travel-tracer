package main
import (
	"jcleary/traveltracer/src/db"
	"jcleary/traveltracer/src/routes"
	"flag"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"strconv"
)

// Config stores the active configuration used by this instance of the server.
type Config struct {
	dbHost string
	port int
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
	router := initRouter()

	// Begin
	log.Println("Listening on port " + strconv.Itoa(config.port) + "... ")
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.port), router))
}

// Init initializes the httprouter.Router and map the routes.
func initRouter() (router *httprouter.Router) {
	log.Print("Initializing handlers... ")
	router = httprouter.New()



	router.GET("/map/point", routes.PointHandler)
	router.GET("/map/route", routes.RouteHandler)
	router.GET("/coordinate", routes.RawCoordinateHandler)
	router.GET("/coordinates/:id", routes.GetCoordinate)
	router.GET("/coordinates", routes.FindAllCoordinates)
	router.GET("/route", routes.RawRouteHandler)
	router.PUT("/db/saveCoordinate", routes.StoreCoordinate)

	return router;
}

// initFlags initializes the command line flags and creates a new Config
func initFlags() Config {
	var dbHost = flag.String("db.host", "localhost", "define the database hostname")
	var appPort = flag.Int("app.port", 8080, "define the port to run this app on")
	flag.Parse()

	return Config{dbHost: *dbHost, port: *appPort}
}