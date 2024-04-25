package routes

import (
	"orderfaz-test-go/handlers"
	"orderfaz-test-go/pkg/middleware"
	"orderfaz-test-go/pkg/mysql"
	"orderfaz-test-go/repositories"

	"github.com/gorilla/mux"
)

func LogisticRoutes(r *mux.Router) {
	logisticRepository := repositories.RepositoryLogistic(mysql.DB)
	h := handlers.HandlerLogistic(logisticRepository)

	r.HandleFunc("/logistics", middleware.Auth(h.FindLogistics)).Methods("GET")
	r.HandleFunc("/logistic/{id}", middleware.Auth(h.GetLogistic)).Methods("GET")
	r.HandleFunc("/getlogistic", middleware.Auth(h.GetLogisticBody)).Methods("GET")
	r.HandleFunc("/logistic", middleware.Auth(h.CreateLogistic)).Methods("POST")
	r.HandleFunc("/logistic/{id}", middleware.Auth(h.UpdateLogistic)).Methods("PATCH")
	r.HandleFunc("/logistic/{id}", middleware.Auth(h.DeleteLogistic)).Methods("DELETE")
}
