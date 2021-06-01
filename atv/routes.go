package atv

import (
	"RedSaludAtv/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/atv/subscription", middleware.GetSubscriptions).Methods(http.MethodGet)

	r.HandleFunc("/api/atv/subscription", middleware.CreateSubscription).Methods(http.MethodPost)
	r.HandleFunc("/api/atv/subscription/stepOne", middleware.CreateSubscriptionStepOne).Methods(http.MethodPost)

	r.HandleFunc("/api/atv/subscription/stepTwo", middleware.UpdateSubscription).Methods(http.MethodPut)
	r.HandleFunc("/api/atv/subscription/validationSubscription", middleware.GetValidationPerson).Methods(http.MethodGet)
	r.HandleFunc("/api/atv/subscription/stepThree", middleware.UpdateDeclaration).Methods(http.MethodPut)

	r.HandleFunc("/api/atv/person", middleware.GetDataPerson).Methods(http.MethodGet)
}
