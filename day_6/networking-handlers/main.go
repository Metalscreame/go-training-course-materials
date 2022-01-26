package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/middlware"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/repository/inmemory"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/repository/postgres"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/server"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Main function
func main() {
	logger, _ := zap.NewProduction()
	repo := inmemory.NewRepository()
	postgres.MigrateUp("migrations", "postgres://postgres:secret@localhost:5432/test?sslmode=disable")
	srv := server.NewServer(repo, logger)
	md := &middlware.Middleware{Logger: logger}
	// Init router
	r := mux.NewRouter()

	// we can also use middleware
	r.Use(md.LoggingMiddleware)

	// Route handles & endpoints
	r.HandleFunc("/books", srv.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", srv.GetBook).Methods("GET")
	r.HandleFunc("/books", srv.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", srv.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", srv.DeleteBook).Methods("DELETE")
	logger.Info("starting web server")

	// pprof handlers
	// if you're using standart listener - just use blank import
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))

	// Start server
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}












// go tool pprof -seconds 30 http://localhost:8080/debug/pprof/heap
// top top5 -cum (cumulative flag)
// go to http://localhost:8080/debug/pprof/
// download heap, do go tool pprof heap
// show https://blog.golang.org/pprof