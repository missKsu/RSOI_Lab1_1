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

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {
	a.Router = a.Router.PathPrefix("/api").Subrouter()
	a.Router.HandleFunc("/",a.Welcome).Methods("GET")
	a.Router.HandleFunc("/add",a.Add).Methods("POST")
	a.Router.HandleFunc("/inc",a.Inc).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main(){
	var router App
	router.Initialize()
	//router.InitializeRoutes()
	
    api := router.Router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/",router.Welcome).Methods("GET")
    api.HandleFunc("/add",router.Add).Methods("POST")
    api.HandleFunc("/inc",router.Inc).Methods("POST")

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router.Router))
    
}

func (a *App) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func (a *App) Add(w http.ResponseWriter, r *http.Request) {
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

func (a *App) Inc(w http.ResponseWriter, r *http.Request) {
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