package main

import (
    "testing"
    "net/url"
    "net/http"
    "io/ioutil"
    "os"
    "net/http/httptest"
    "strings"
    "io"
    )

type IncData struct{
    element string
    result string
}

type AddData struct{
    sum string
    element string
    result string
}

var a App

func TestMain(m *testing.M) {
    a = App{}
    a.Initialize()

    code := m.Run()

    os.Exit(code)
}

func TestWelcome(t *testing.T) {
    res := check("GET","/api/",nil)

        if res != "Hello World!!!" {
            t.Errorf("%s", res)
        }
}

func TestInc(t *testing.T) {
    tests:=[]IncData{{"4","5"},{"2","3"},{"","Incorrect request"},}
    for _,val := range tests{
        query:="element="+val.element
        res := check("POST","/api/inc",strings.NewReader(query))

        if res != val.result {
            t.Errorf("%s %s", val.element,res)
        }
    }
}

func TestAdd(t *testing.T) {
    tests:=[]AddData{{"4","5","9"},{"2","3","5"},{"","3","Incorrect"}}
    for _,val := range tests{
        urlData := url.Values{}
        urlData.Set("sum", val.sum)
        urlData.Set("element", val.element)
        query :="sum="+val.sum+"&element="+val.element
        res := check("POST","/api/add",strings.NewReader(query))

        if res != val.result {
            t.Errorf("%s", res)
        }
    }
}

func check(method string, url string, body io.Reader) string {
    req, err := http.NewRequest(method, url, body)
    checkErr(err)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    response := executeRequest(req)
    res := response.Result()
    /*
    var response *http.Response
    var err error
    if method=="Get"{
        response, err = http.Get(url)
    }else if method=="Post"{
        response, err = http.PostForm(url,params)
    }
    checkErr(err)
    defer response.Body.Close()
    */
    if res.StatusCode!=200{
        return res.Status
    }else{
        contents, err := ioutil.ReadAll(res.Body)
        checkErr(err)
        return string(contents)
    }    
    return ""
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}