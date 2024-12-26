package dbutils

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}

	// Create train table
	_, statementError := statement.Exec()

	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()

	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()

	log.Println("All tables created successfully!")
}

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register adds paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	var requestBody = request.Request.Body
	log.Println(requestBody)
	decoder := json.NewDecoder(requestBody)

	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)

	// Error handling is obvious here. So omitting...
	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)

	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}
