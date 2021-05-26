package middleware

import (
	"RedSaludAtv/atv/dao"
	"RedSaludAtv/atv/entites"
	"RedSaludAtv/config"
	"RedSaludAtv/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
			if rows > 0 {
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

func GetDataPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var dataQuery entites.DataQuery
	_ = json.NewDecoder(r.Body).Decode(&dataQuery)

	response, err := http.Get("https://api.dniruc.com/api/search/dni/" + dataQuery.DniQuery + "/Test_4b425b9c9776b4e9896fac7b3e29829e6366f693")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(responseData))

			var responsePerson entites.PersonReniec

			json.Unmarshal(responseData, &responsePerson)

			fmt.Println(responsePerson.Data)

			utils.RespondWithSuccess(responsePerson.Data, w)
		}
	}
}
