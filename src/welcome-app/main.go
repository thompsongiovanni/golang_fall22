package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

/* type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct {
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	JsonNested JsonNested `json:"JsonNested`
}

type JsonNested struct {
	NestedValue1 string `json:"nestedKey1"`
	NestedValue2 string `json:"nestedKey2"`
} */

type Welcome struct {
	Name string
	Time string
}

type ReturnResp struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	ReturnNested ReturnNested `json:"ReturnNested"`
}

type ReturnNested struct {
	NestedStreet string `json:"street"`
	NestedCity string `json:"city"`
	NestedEmail string `json:email`
	NestedPhone string `json:phone`
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	/*nested := JsonNested{
		NestedValue1: "first nested value",
		NestedValue2: "second nested value",
	}

	jsonResp := JsonResponse{
		Value1: "some Data",
		Value2: "other Data",
		JsonNested: nested,
	} */

	nested := ReturnNested{
		NestedStreet: "123 Vice City",
		NestedCity: "Las Vegas, Nevada",
		NestedEmail: "johndoe123@gmail.com",
		NestedPhone: "(718) 555-9037",
	}

	returnResp := ReturnResp{
		FirstName: "John",
		LastName: "Doe",
		ReturnNested: nested,
	}
	
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	/*http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})*/

	http.HandleFunc("/returnResp", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(returnResp)
	})

	// third path, get/fetch, return an json object like an API , 2 nested objects {firstname:"", lastname:"", address:{street:"", city...}, contactInfo:{email:"", phone:""}}
	// 3 new structs, 2 new values

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
