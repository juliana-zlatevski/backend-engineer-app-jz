package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var dbFile = "employees.db"

// db instance
type Employees struct {
	db *sql.DB
}

// server instance
type httpServer struct {
	Employees *Employees
}

type Employee struct {
	ID     int    `json:"id"`
	Gender string `json:"item"`
}

func NewEmployees() (*Employees, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return &Employees{
		db: db,
	}, nil
}

func (e *Employee) Scan() []interface{} {
	return []interface{}{
		&e.ID,
		&e.Gender,
	}
}

func (server *httpServer) listAllEmployees(w http.ResponseWriter, r *http.Request) {
	var err error

	resp, err := server.Employees.getAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (db *Employees) getAllEmployees() ([]Employee, error) {
	employees := []Employee{}

	qry := `SELECT * FROM employees`

	rows, err := db.db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(employee.Scan()...)
		if err != nil {
			continue
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func NewHTTPServer(addr string) *http.Server {
	var err error

	employees, err := NewEmployees()
	if err != nil {
		log.Fatal(err)
	}

	server := &httpServer{
		Employees: employees,
	}

	r := mux.NewRouter()
	r.HandleFunc("/employees", server.listAllEmployees).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func getEnvVars(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	appPort := ":" + getEnvVars("PORT")
	server := NewHTTPServer(appPort)
	log.Fatal(server.ListenAndServe())
}
