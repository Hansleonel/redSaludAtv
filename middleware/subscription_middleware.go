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
	"net/url"
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

		if subscriptionStepOne.Importe <= 0.0 || subscriptionStepOne.NumDoc == "" || subscriptionStepOne.CelNumber == "" || subscriptionStepOne.TcDatos == "" {
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

func CreateSubscriptionFamiliar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var subscriptionFamiliar entites.SubsFamiliar
	_ = json.NewDecoder(r.Body).Decode(&subscriptionFamiliar)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println()
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		if subscriptionFamiliar.NumDoc == "" || subscriptionFamiliar.TcDatos == "" || subscriptionFamiliar.FechaNacimiento == "" || subscriptionFamiliar.Ape1 == "" || subscriptionFamiliar.Ape2 == "" || subscriptionFamiliar.Nom1 == "" || subscriptionFamiliar.Nom2 == "" || subscriptionFamiliar.IdContratante == "" {
			requestError := entites.SubsError{
				Type:   "/api/atv/subscription/createFamiliar",
				Title:  "Error 400",
				Detail: "Bad Request, values required",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(requestError)
			return
		}

		err := subscriptionDao.CreateFamiliar(&subscriptionFamiliar)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(subscriptionFamiliar)
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

func CreateDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var questions entites.RequestSubsDeclaration
	err := json.NewDecoder(r.Body).Decode(&questions)

	db, err2 := config.GetMySQLDB()

	if err != nil && err2 != nil {
		utils.RespondWithError(err2, w)
	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		err := subscriptionDao.CreateDeclaration(questions.Data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(questions)
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

		if subscriptionStepTwo.FechaNacimiento == "" || subscriptionStepTwo.Ape1 == "" || subscriptionStepTwo.Ape2 == "" || subscriptionStepTwo.Nom1 == "" || subscriptionStepTwo.Nom2 == "" || subscriptionStepTwo.Mail == "" {
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
	TOKEN := "2300ffa8c8403056fe54a11a4ce463845c47b9d156d1642ed8b1311fe9f6f577"

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Authorization", "Bearer "+TOKEN)

	// r.Header.Add("Content-Type","")

	fmt.Println(TOKEN)
	fmt.Println(w.Header())

	var dataQuery entites.DataQuery
	_ = json.NewDecoder(r.Body).Decode(&dataQuery)

	// response, err := http.Get("https://api.dniruc.com/api/search/dni/" + dataQuery.DniQuery + "/Test_4b425b9c9776b4e9896fac7b3e29829e6366f693")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	response, err := http.Get("https://consulta.apiperu.pe/api/dni/" + dataQuery.DniQuery)

	fmt.Println(response)

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

func GetDataPerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var dataQuery entites.DataQuery
	_ = json.NewDecoder(r.Body).Decode(&dataQuery)

	if len(dataQuery.DniQuery) < 8 {
		requestError := entites.SubsError{
			Type:   "/api/atv/dni",
			Title:  "Error 400",
			Detail: "Bad Request, length of query must have just 8 digits",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestError)
		return
	}

	reqURL, _ := url.Parse("https://consulta.apiperu.pe/api/dni/" + dataQuery.DniQuery)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer 2300ffa8c8403056fe54a11a4ce463845c47b9d156d1642ed8b1311fe9f6f577"},
		},
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		log.Fatal("Error", err)

	} else {

		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			log.Fatal(err)

		} else {

			var responsePerson entites.PersonReniec
			json.Unmarshal(responseData, &responsePerson)

			fmt.Println(responsePerson.Data)

			if responsePerson.Data.ApellidoPaterno == "" {
				requestError := entites.SubsError{
					Type:   "/api/atv/dni",
					Title:  "Error 404",
					Detail: "Bad Request,dni not found",
				}
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(requestError)
				return
			}

			utils.RespondWithSuccess(responsePerson.Data, w)

		}
	}
}

func GetDataEnterprise(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var dataQueryRuc entites.DataQueryRuc
	_ = json.NewDecoder(r.Body).Decode(&dataQueryRuc)

	if len(dataQueryRuc.RucQuery) < 11 {
		requestError := entites.SubsError{
			Type:   "/api/atv/ruc",
			Title:  "Error 400",
			Detail: "Bad Request, length of query must have just 11 digits",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestError)
		return
	}

	reqURL, _ := url.Parse("https://consulta.apiperu.pe/api/ruc/" + dataQueryRuc.RucQuery)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer 2300ffa8c8403056fe54a11a4ce463845c47b9d156d1642ed8b1311fe9f6f577"},
		},
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		log.Fatal("Error", err)

	} else {

		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			log.Fatal(err)

		} else {

			var responseEnterprise entites.EnterpriseSunat
			json.Unmarshal(responseData, &responseEnterprise)

			// fmt.Println(responseEnterprise)
			fmt.Println(responseEnterprise.Data)
			// fmt.Println("UBIGEO",responseEnterprise.Data.Ubigeo)

			utils.RespondWithSuccess(responseEnterprise.Data, w)
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
