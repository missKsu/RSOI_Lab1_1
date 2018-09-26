package main

import (
	"net/http"
    "log"
    "fmt"
	"github.com/gorilla/mux"
	"strconv")

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	router := mux.NewRouter().StrictSlash(true)

    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/",Welcome).Methods("GET")
    api.HandleFunc("/add",Add).Methods("POST")
    api.HandleFunc("/inc",Inc).Methods("POST")

    log.Fatal(http.ListenAndServe(":3000", router))
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
    res := CheckErr(err)
    if res!="Continue" {
    	fmt.Fprint(w,res)
    }else{
    	sum:=r.Form.Get("sum")
    	CheckErr(err)
    	
    	element:=r.Form.Get("element")
    	CheckErr(err)

    	if sum != "" && element != ""{
    		var1,err := strconv.Atoi(sum)
    		CheckErr(err)
    		var2,err := strconv.Atoi(element)
    		CheckErr(err)
    		fmt.Fprint(w,var1+var2)
    	}else{
    		fmt.Fprint(w, "Incorrect request")	
    	}
    }
}

func Inc(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
    res := CheckErr(err)
    if res!="Continue" {
    	fmt.Fprint(w,res)
    }else{
    	element:=r.Form.Get("element")
    	CheckErr(err)

    	if element != ""{
    		variable,err := strconv.Atoi(element)
    		CheckErr(err)
    		fmt.Fprint(w,variable+1)
    	}else{
    		fmt.Fprint(w, "Incorrect request")	
    	}
    }
}

func CheckErr(err error) string {
	if err != nil {
		return "Error occured: "+err.Error()
	}else{
		return "Continue"
	}
}