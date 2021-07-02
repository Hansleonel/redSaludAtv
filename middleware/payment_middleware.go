package middleware

import (
	"RedSaludAtv/atv/entites"
	"RedSaludAtv/atv/entites/payments"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var paymentPayU payments.PaymentPayU
	_ = json.NewDecoder(r.Body).Decode(&paymentPayU)

	var clientPayU payments.ClientPayU
	clientPayU.Email = paymentPayU.Mail
	clientPayU.FullName = paymentPayU.CompleteName

	fmt.Println(clientPayU.Email)
	requestByte, _ := json.Marshal(clientPayU)
	requestReader := bytes.NewReader(requestByte)

	/* var clientRequest payments.ClientPayU
	clientRequest.Email =  validationDni.DniValidation
	*/
	if paymentPayU.CompleteName == "" || paymentPayU.Document == "" || paymentPayU.CardNumber == "" || paymentPayU.ExpMonth == "" || paymentPayU.ExpYear == "" || paymentPayU.Type == "" || paymentPayU.Phone == "" || paymentPayU.Mail == "" {
		errorRequestValidation(w, "Value Required")
		// log.Fatal("error")
		return
	}

	reqBody := ioutil.NopCloser(requestReader)

	reqURL, _ := url.Parse("https://api.payulatam.com/payments-api/rest/v4.9/customers")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
			"Authorization": {"Basic STlkR1VzUzh2TGdsNFB0Om0xcmdaWDI1OE0yZnNNNEpBMGhGSzJhanJa"},
		},
		Body: reqBody,
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		errorRequest(w, err)
		// log.Fatal("error", err.Error())
		return

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			errorRequest(w, err)
			// log.Fatal("error," + err.Error())
			return

		} else {

			var responseClient payments.ClientPayUResponse
			json.Unmarshal(responseData, &responseClient)

			// fmt.Println(responseEnterprise)
			fmt.Println(responseClient.Id)
			// fmt.Println("UBIGEO",responseEnterprise.Data.Ubigeo)

			// utils.RespondWithSuccess(responseClient, w)
			CreateCard(w, paymentPayU, responseClient.Id)

		}
	}
}

func CreateCard(w http.ResponseWriter, paymentPayU payments.PaymentPayU, idClient string) {
	// w.Header().Add("Content-type", "application/json")

	// reqBody := ioutil.NopCloser(r.Body)

	var cardPayU payments.CardPayU
	cardPayU.Name = paymentPayU.CompleteName
	cardPayU.Document = paymentPayU.Document
	cardPayU.Number = paymentPayU.CardNumber
	cardPayU.ExpMonth = paymentPayU.ExpMonth
	cardPayU.ExpYear = paymentPayU.ExpYear
	cardPayU.Type = paymentPayU.Type
	cardPayU.Address.Line2 = "Address Name"
	cardPayU.Address.Line2 = "17 25"
	cardPayU.Address.Line3 = "of 301"
	cardPayU.Address.PostalCode = "15021"
	cardPayU.Address.City = "Lima"
	cardPayU.Address.State = "Lima"
	cardPayU.Address.Country = "PE"
	cardPayU.Address.Phone = paymentPayU.Phone

	fmt.Println("CardPayU Number")
	fmt.Println(cardPayU.Number)

	requestByte, _ := json.Marshal(cardPayU)
	requestReader := bytes.NewReader(requestByte)

	reqBody := ioutil.NopCloser(requestReader)

	reqURL, _ := url.Parse("https://api.payulatam.com/payments-api/rest/v4.9/customers/" + idClient + "/creditCards")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
			"Authorization": {"Basic STlkR1VzUzh2TGdsNFB0Om0xcmdaWDI1OE0yZnNNNEpBMGhGSzJhanJa"},
		},
		Body: reqBody,
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		errorRequest(w, err)
		// log.Fatal("error", err.Error())
		return

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			errorRequest(w, err)
			// log.Fatal("error", err.Error())
			return

		} else {
			var responseCardPayU payments.CardPayUResponse
			json.Unmarshal(responseData, &responseCardPayU)
			fmt.Println(responseCardPayU.Token)

			CreateRecurrence(w, idClient, responseCardPayU.Token)
		}
	}

}

func CreateRecurrence(w http.ResponseWriter, idClient string, tokenPayU string) {

	var recurrencePayU payments.RecurrencePayU
	recurrencePayU.Quantity = "1"
	recurrencePayU.Installments = "1"
	recurrencePayU.TrialDays = "0"
	recurrencePayU.NotifyUrl = "https://mercadopago-devzamse.herokuapp.com/webcheckout"
	recurrencePayU.Customer.Id = idClient
	recurrencePayU.Customer.CreditCards = []payments.CreditCards{payments.CreditCards{Token: tokenPayU}}
	recurrencePayU.Plan.PlanCode = "plan-day-oficial-1"

	requestByte, _ := json.Marshal(recurrencePayU)
	requestReader := bytes.NewReader(requestByte)

	reqBody := ioutil.NopCloser(requestReader)

	reqURL, _ := url.Parse("https://api.payulatam.com/payments-api/rest/v4.9/subscriptions/")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
			"Authorization": {"Basic STlkR1VzUzh2TGdsNFB0Om0xcmdaWDI1OE0yZnNNNEpBMGhGSzJhanJa"},
		},
		Body: reqBody,
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		errorRequest(w, err)
		// log.Fatal("error", err.Error())
		return

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			errorRequest(w, err)
			// log.Fatal("error", err.Error())
			return

		} else {

			var responseRecurrence payments.RecurrencePayUResponse
			json.Unmarshal(responseData, &responseRecurrence)
			fmt.Println("el ID es")
			fmt.Println(responseRecurrence.Id)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(responseRecurrence)
		}
	}
}

func errorRequest(w http.ResponseWriter, err error) {
	requestError := entites.SubsError{
		Type:   "/api/atv/payment",
		Title:  "Error 400",
		Detail: "Bad Request," + err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(requestError)
	return
}

func errorRequestValidation(w http.ResponseWriter, err string) {
	requestError := entites.SubsError{
		Type:   "/api/atv/payment",
		Title:  "Error 400",
		Detail: "Bad Request," + err,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(requestError)
	return
}

/* func CreateMultiPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	var paymentGroupPayU payments.PaymentGroupPayU
	_ = json.NewDecoder(r.Body).Decode(&paymentGroupPayU)

	fmt.Println(paymentGroupPayU)
	var paymentPayU payments.PaymentPayU
	_ = json.NewDecoder(r.Body).Decode(&paymentPayU)

	var paymentClient payments.PaymentClient
	paymentClient.IdClient = paymentPayU.IdContratante

	db, err := config.GetMySQLDB()
	if err != nil {

	} else {
		subscriptionDao := dao.SubscriptionDao{
			Db: db,
		}

		clients, err := subscriptionDao.FindClients(paymentClient.IdClient)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println("Size")
			fmt.Println(len(clients))

			for i := 0; i < len(clients); i++ {
				go CreatePlan(w, r, clients[0].IdClient)
			}
		}
	}

}*/
