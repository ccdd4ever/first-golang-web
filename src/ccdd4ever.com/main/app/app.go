package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func sayHello(res http.ResponseWriter,req *http.Request){
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("PATH",req.URL.Path)
	fmt.Println("schema",req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k,v:=range req.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprint(res,"hello golang web")


}
func main() {
	http.HandleFunc("/",sayHello)
	err :=http.ListenAndServe(":9090",nil)
	if err !=nil{
		log.Fatalf("ListenAndServe",err)
	}
}