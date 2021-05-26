package atv

import (
	"RedSaludAtv/middleware"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/atv/subscription", middleware.GetSubscriptions).Methods("GET")

	r.HandleFunc("/api/atv/subscription", middleware.CreateSubscription).Methods("POST")
	r.HandleFunc("/api/atv/subscription/stepOne", middleware.CreateSubscriptionStepOne).Methods("POST")

	r.HandleFunc("/api/atv/subscription/stepTwo", middleware.UpdateSubscription).Methods("PUT")

	r.HandleFunc("/api/atv/person", middleware.GetDataPerson).Methods("GET")
}
