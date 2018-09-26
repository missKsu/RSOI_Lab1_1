package main

import (
    "testing"
    "net/url"
    "net/http"
    "io/ioutil"
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

func TestInc(t *testing.T) {
    tests:=[]IncData{{"4","5"},{"2","3"},{"","Incorrect request"}}
    for _,val := range tests{
        urlData := url.Values{}
        urlData.Set("element", val.element)
        res := check("Post","http://whispering-badlands-32857.herokuapp.com/api/inc",urlData)

        if res != val.result {
            t.Errorf("%s", res)
        }
    }
}

func TestAdd(t *testing.T) {
    tests:=[]AddData{{"4","5","9"},{"2","3","5"},{"","3","Incorrect request"}}
    for _,val := range tests{
        urlData := url.Values{}
        urlData.Set("sum", val.sum)
        urlData.Set("element", val.element)
        res := check("Post","http://whispering-badlands-32857.herokuapp.com/api/add",urlData)

        if res != val.result {
            t.Errorf("%s", res)
        }
    }
}

func check(method string, url string, params url.Values) string {
    var response *http.Response
    var err error
    if method=="Get"{
        response, err = http.Get(url)
    }else if method=="Post"{
        response, err = http.PostForm(url,params)
    }
    checkErr(err)
    defer response.Body.Close()

    if response.StatusCode!=200{
        return response.Status
    }else{
        contents, err := ioutil.ReadAll(response.Body)
        checkErr(err)
        return string(contents)
    }    
    return ""
}