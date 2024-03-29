package main

/*
	Author : Iordanis Paschalidis
	Date   : 03/12/2021

*/

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	config "github.com/junkd0g/covid-vaccine/internal/config"
	"github.com/junkd0g/covid-vaccine/internal/controller"
)

//Service object that contains the Port and Router of the application
type Service struct {
	Port   string
	Router *mux.Router
}

/*
   Running the service in port 8888 (getting the value from ./assets/config/production.json )

       Endpoints:
		GET:
			api/data/{country}
		POST:
*/
func (s Service) run() {

	configData, err := config.GetAppConfig("./config.yaml")
	if err != nil {
		panic(fmt.Errorf("creating_config %w", err))

	}

	s.Port = configData.Server.Port
	country, err := controller.NewCountry()
	if err != nil {
		panic(fmt.Errorf("creating_mail_controller %w", err))
	}
	s.Router.HandleFunc("/api/data/{country}", country.Middleware).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*", "Authorization"},
	})

	handler := c.Handler(s.Router)

	fmt.Println("server running at port " + s.Port)
	http.ListenAndServe(s.Port, handler)
}

func main() {
	service := Service{Router: mux.NewRouter().StrictSlash(true)}
	service.run()
}
