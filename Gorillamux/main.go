package main

import (
	"context"
	"golangProjects/Microservice/Gorillamux/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	// simple http connection

	// New creates a new Logger. The out variable sets the destination to which log data will be written. The prefix appears at the beginning of each generated log line, or after the log header if the Lmsgprefix flag is provided. The flag argument defines the logging properties.
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// handlers is directory and NewHello is logger method
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoogBye(l)
	ph := handlers.NewProduct(l)

	// NewServeMux allocates and returns a new ServeMux
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)

	putrouter := sm.Methods(http.MethodPut).Subrouter()
	putrouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putrouter.Use(ph.MiddlewareProductValidation)

	postrouter := sm.Methods(http.MethodPost).Subrouter()
	postrouter.HandleFunc("/", ph.AddProducts)
	postrouter.Use(ph.MiddlewareProductValidation)

	deleterouter := sm.Methods("DELETE").Subrouter()
	deleterouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Rest APi  (reprentational state transfer)
	// json over http get post delete put (edit)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// ListenAndServe listens on the TCP network address addr and then calls Serve with

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)

	//Notify causes package signal to relay incoming signals to c.
	//If no signals are provided, all incoming signals will be
	//relayed to c. Otherwise, just the provided signals will.
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
