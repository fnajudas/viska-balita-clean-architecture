package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"viska/db"

	authcontroller "viska/handler/authController"
	createData "viska/handler/create"
	deleteData "viska/handler/delete"
	listData "viska/handler/read"
	updateData "viska/handler/update"
	dataToken "viska/middleware/jwt"

	"github.com/gorilla/mux"
)

const (
	writeRTO = 30
	readRTO  = 30
)

// Routes object of package 'routes'
type Routes struct {
	Router         *mux.Router
	ListDataBalita *listData.Handler
	DtRegister     *authcontroller.Auth
	DtLogin        *authcontroller.Auth
	DtToken        *dataToken.TokenGenerator
	UpdateData     *updateData.Handler
	CreateData     *createData.Handler
	DeleteData     *deleteData.Handler
}

func (r *Routes) Run(port string) {
	// Inisialisasi
	routes := mux.NewRouter()
	baseURL := os.Getenv("BASE_URL_PATH")

	if len(baseURL) > 0 && baseURL != "/" {
		routes.PathPrefix(baseURL).HandlerFunc(db.URLRewriter(routes, baseURL))
	}

	routes.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET", "OPTIONS")

	routes.HandleFunc("/register", r.DtRegister.Register).Methods("POST")
	routes.HandleFunc("/login", r.DtLogin.Login).Methods("POST")

	api := routes.PathPrefix("/api").Subrouter()
	api.Use(r.DtToken.MiddlewareJWT)
	api.HandleFunc("/read", r.ListDataBalita.GetDataBalita).Methods("GET")
	api.HandleFunc("/update", r.UpdateData.UpdateDataBalita).Methods("PUT")
	api.HandleFunc("/create", r.CreateData.CreateBalitaHandler).Methods("POST")
	api.HandleFunc("/delete", r.DeleteData.DeleteBalita).Methods("DELETE")

	r.Router = routes
	//C. Serving RESTful HTTP to Clients
	log.Print(fmt.Sprintf("[HTTP SRV] Listening on port :%s", port))
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: writeRTO * time.Second,
		ReadTimeout:  readRTO * time.Second,
	}
	log.Panic(srv.ListenAndServe())
}
