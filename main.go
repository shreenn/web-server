package main

import(
	"fmt"
	"log"
	"net/http"
)
//2
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!= "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method!="GET"{
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

//3
func formHandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w,"POST request successful")
	name:= r.FormValue("name")
	address:= r.FormValue("address")

	fmt.Fprintf(w,"Name=%s\n",name)
	fmt.Fprintf(w,"Address=%s\n",address)
}


//1
func main(){
	fileeeeServer := http. FileServer(http.Dir("./static"))
	http.Handle("/", fileeeeServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")

	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}