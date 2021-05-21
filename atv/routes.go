package atv

import (
	"RedSaludAtv/middleware"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/atv/subscription", middleware.GetSubscriptions).Methods("GET")

	r.HandleFunc("/api/atv/subscription", middleware.CreateSubscription).Methods("POST")
}
