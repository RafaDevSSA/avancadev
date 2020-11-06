package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/hashicorp/go-retryablehttp"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
	Message string
}

var coupons Coupons

func main() {
	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	productId := r.PostFormValue("productId")
	valid := coupons.Check(coupon)

	baseUrl := "http://localhost:3333/"
	
	url := fmt.Sprintf("%s%s", baseUrl, productId)

	
	resultProductMessage := makeHttpCall(url)

	result := Result{Status: valid, Message: resultProductMessage}
	
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
	
}


func makeHttpCall(urlMicroservice string) string {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.Get(urlMicroservice)
	if err != nil {
		return "Servidor fora do ar!"
	}
	
	
	defer res.Body.Close()
	
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}
	
	return string(data)

}
