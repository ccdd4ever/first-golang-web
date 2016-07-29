package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"html/template"
	"regexp"
	"time"
	"crypto/md5"
	"io"
	"strconv"
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

func login(res http.ResponseWriter,req *http.Request) {
	fmt.Println("method:",req.Method)
	if req.Method=="GET" {
		curtime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(curtime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		fmt.Println(token)

		t,_:=template.ParseFiles("front/src/templates/login.gtpl")
		t.Execute(res,token)
	}else {
		req.ParseForm()//显式调用解析表单数据
		fmt.Println("username:",req.Form["username"])
		fmt.Println("password:",req.Form["password"])
		//token校验,防止重复提交
		token:=req.Form.Get("token")
		if token!="" {
			//do some token check
		}else {
			//token is empty,return an error
		}
		//验证表单输入
		if len(req.Form["username"][0])==0{
			fmt.Fprint(res,"username must be input")
		}
		//正则判断
		if m,_:=regexp.MatchString(("^[0-9]+$"),req.Form.Get("password"));!m{
			fmt.Fprint(res,"invalid password")
		}else {
			fmt.Fprint(res,"welcome")
		}
	}
	
}

func main() {
	http.HandleFunc("/",sayHello)//绑定请求路径
	http.HandleFunc("/login",login)
	err :=http.ListenAndServe(":9090",nil)//监听端口
	if err !=nil{
		log.Fatalf("ListenAndServe",err)
	}
}