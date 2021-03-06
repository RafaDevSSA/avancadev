package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	Status string
	Message string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9091", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")
	productId := r.PostFormValue("productId")

	resultCoupon := makeHttpCall("http://localhost:9092", coupon, productId)

	result := Result{Status: "declined"}

	if ccNumber == "1" {
		result.Status = "approved"
		result.Message = resultCoupon.Message
	}

	if resultCoupon.Status == "invalid" {
		result.Status = "invalid coupon"
	}


	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}


func makeHttpCall(urlMicroservice string, coupon string, productId string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("productId", productId)

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Servidor fora do ar!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}
