package atv

import (
	"RedSaludAtv/middleware"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/atv/subscription", middleware.GetSubscriptions).Methods("GET")

	r.HandleFunc("/api/atv/subscription", middleware.CreateSubscription).Methods("POST")
	r.HandleFunc("/api/atv/subscription/stepOne", middleware.CreateSubscriptionStepOne).Methods("POST")
	r.HandleFunc("/api/atv/subscription/questionDetail", middleware.CreateQuestionDetail).Methods("POST")
	r.HandleFunc("/api/atv/subscription/createFamiliar", middleware.CreateSubscriptionFamiliar).Methods("POST")

	r.HandleFunc("/api/atv/subscription/stepTwo", middleware.UpdateSubscription).Methods("PUT")
	r.HandleFunc("/api/atv/subscription/validationSubscription", middleware.PostValidationPerson).Methods("POST")
	r.HandleFunc("/api/atv/subscription/stepThree", middleware.UpdateDeclaration).Methods("PUT")

	r.HandleFunc("/api/atv/person", middleware.GetDataPerson).Methods("POST")
}
