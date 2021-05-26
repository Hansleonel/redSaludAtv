package middleware

import (
	"RedSaludAtv/atv/dao"
	"RedSaludAtv/atv/entites"
	"RedSaludAtv/config"
	"RedSaludAtv/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var subscription entites.Subscription
	fmt.Println("Body")
	fmt.Println(r.Body)
	_ = json.NewDecoder(r.Body).Decode(&subscription)

	fmt.Println("Subs")
	fmt.Println(subscription)
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if subscription.FechaNacimiento == "" || subscription.CelNumber == "" || subscription.TcDatos == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription",
				Title:  "Error 400",
				Detail: "Bad Request",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		err := subscriptionDao.Create(&subscription)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(subscription)
		}
	}
}

func CreateSubscriptionStepOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var subscriptionStepOne entites.SubscriptionStepOne
	_ = json.NewDecoder(r.Body).Decode(&subscriptionStepOne)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if subscriptionStepOne.FechaNacimiento == "" || subscriptionStepOne.CelNumber == "" || subscriptionStepOne.TcDatos == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/stepOne",
				Title:  "Error 400",
				Detail: "Bad Request",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		err := subscriptionDao.CreateStepOne(&subscriptionStepOne)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(subscriptionStepOne)
		}

	}
}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var subscriptionStepTwo entites.SubscriptionStepTwo
	_ = json.NewDecoder(r.Body).Decode(&subscriptionStepTwo)

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if subscriptionStepTwo.NumDoc == "" || subscriptionStepTwo.Ape1 == "" || subscriptionStepTwo.Ape2 == "" || subscriptionStepTwo.Nom1 == "" || subscriptionStepTwo.Nom2 == "" || subscriptionStepTwo.Mail == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/stepTwo",
				Title:  "Error 400",
				Detail: "Bad Request",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		rows, err := subscriptionDao.Update(subscriptionStepTwo)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			if rows>0{
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(subscriptionStepTwo)
			}
		}
	}

}

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := config.GetMySQLDB()
	database := dao.SubscriptionDao{
		Db: db,
	}
	subscriptions, err := database.FindAll()

	if err == nil {
		utils.RespondWithSuccess(subscriptions, w)
	} else {
		utils.RespondWithError(err, w)
	}
}
