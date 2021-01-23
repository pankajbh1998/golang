package main

import (
	"net/http"
)
func SayHello(w http.ResponseWriter,r *http.Request) {
	query,ok:=r.URL.Query()["id"]
	if ok {
		que:=query[0]
		if que=="-1" {
			w.WriteHeader(100)
			w.Write([]byte("Hello Negative"))
		} else if que =="0" {
			w.WriteHeader(100)
			w.Write([]byte("Hello"))
		} else if que =="1" {
			w.WriteHeader(100)
			w.Write([]byte("Hello Positive"))
		} else if checkNumeric(que){
			w.WriteHeader(500)
			w.Write([]byte("Not Found"))
		} else  {
			w.WriteHeader(400)
			w.Write([]byte("Invalid"))
		}
	}
}

func checkNumeric(que string) bool {
	for _,val:= range que {
		if '0'>val||'9'<val {
			return false
		}
	}
	return true
}
func main(){
	http.HandleFunc("/ping",SayHello)

}
