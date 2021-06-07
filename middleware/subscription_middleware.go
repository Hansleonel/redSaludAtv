package middleware

import (
	"RedSaludAtv/atv/dao"
	"RedSaludAtv/atv/entites"
	"RedSaludAtv/config"
	"RedSaludAtv/utils"
	"crypto/tls"
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

		if subscriptionStepOne.TypeDoc == "" || subscriptionStepOne.NumDoc == "" || subscriptionStepOne.FechaNacimiento == "" || subscriptionStepOne.CelNumber == "" || subscriptionStepOne.TcDatos == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/stepOne",
				Title:  "Error 400",
				Detail: "Bad Request, values required",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		err := subscriptionDao.CreateStepOne(&subscriptionStepOne)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(subscriptionStepOne)
		}

	}
}

func CreateQuestionDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var questionDetail entites.SubsQuestions
	_ = json.NewDecoder(r.Body).Decode(&questionDetail)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if questionDetail.Question == "" || questionDetail.Description == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/questionDetail",
				Title:  "Error 400",
				Detail: "Bad Request, values required",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		err := subscriptionDao.CreateQuestionDetail(&questionDetail)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(questionDetail)
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

		if subscriptionStepTwo.Ape1 == "" || subscriptionStepTwo.Ape2 == "" || subscriptionStepTwo.Nom1 == "" || subscriptionStepTwo.Nom2 == "" || subscriptionStepTwo.Mail == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/stepTwo",
				Title:  "Error 400",
				Detail: "Bad Request, Values Required",
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
				w.WriteHeader(http.StatusOK)
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

func PostDataPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", "Bearer 2300ffa8c8403056fe54a11a4ce463845c47b9d156d1642ed8b1311fe9f6f577")

	var dataQuery entites.DataQuery
	_ = json.NewDecoder(r.Body).Decode(&dataQuery)

	// response, err := http.Get("https://api.dniruc.com/api/search/dni/" + dataQuery.DniQuery + "/Test_4b425b9c9776b4e9896fac7b3e29829e6366f693")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	response, err := http.Get("https://consulta.apiperu.pe/api/dni/" + dataQuery.DniQuery)
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

func PostValidationPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var validationDni entites.ValidationPerson
	_ = json.NewDecoder(r.Body).Decode(&validationDni)

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if len(validationDni.DniValidation) != 8 {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/validationPerson",
				Title:  "Error 400",
				Detail: "Bad Request, Length of Dni",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		validate, err := subscriptionDao.ValidationDni(validationDni)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)

			if validate == 0 {
				validationDni.Validate = true
			} else {
				validationDni.Validate = false
			}

			json.NewEncoder(w).Encode(validationDni)
		}
	}
}

func UpdateDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var subscriptionStepThree entites.SubscriptionStepThree
	_ = json.NewDecoder(r.Body).Decode(&subscriptionStepThree)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if subscriptionStepThree.DecJur == "" || subscriptionStepThree.QuestionFirst == "" || subscriptionStepThree.QuestionSecond == "" || subscriptionStepThree.QuestionThird == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/stepThree",
				Title:  "Error 400",
				Detail: "Bad Request,Values Required",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		rows, err := subscriptionDao.UpdateDeclaration(subscriptionStepThree)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			if rows > 0 {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(subscriptionStepThree)
			}
		}

	}
}
