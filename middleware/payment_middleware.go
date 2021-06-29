package middleware

import (
	"RedSaludAtv/atv/entites/payments"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var clientPayU payments.ClientPayU
	_ = json.NewDecoder(r.Body).Decode(&clientPayU)

	requestByte, _ := json.Marshal(clientPayU)
	requestReader := bytes.NewReader(requestByte)

	/* var clientRequest payments.ClientPayU
	clientRequest.Email =  validationDni.DniValidation
	*/

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

		log.Fatal("Error", err)

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			log.Fatal(err)

		} else {

			var responseClient payments.ClientPayUResponse
			json.Unmarshal(responseData, &responseClient)

			// fmt.Println(responseEnterprise)
			fmt.Println(responseClient.Id)
			// fmt.Println("UBIGEO",responseEnterprise.Data.Ubigeo)

			// utils.RespondWithSuccess(responseClient, w)
			// CreateCard(responseClient.Id)

		}
	}
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	reqBody := ioutil.NopCloser(r.Body)

	reqURL, _ := url.Parse("https://api.payulatam.com/payments-api/rest/v4.9/customers/8cfe88q78xgl/creditCards")
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

		log.Fatal("Error", err)

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			log.Fatal(err)

		} else {
			var responseCardPayU payments.CardPayUResponse
			json.Unmarshal(responseData, &responseCardPayU)
			fmt.Println(responseCardPayU)

		}
	}

}

func CreateRecurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	reqBody := ioutil.NopCloser(r.Body)

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

		log.Fatal("Error", err)

	} else {
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {

			log.Fatal(err)

		} else {

			var responseRecurrence payments.RecurrencePayUResponse
			json.Unmarshal(responseData, &responseRecurrence)
			fmt.Println(responseRecurrence)

		}
	}
}
