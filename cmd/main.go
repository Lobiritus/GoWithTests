package main

import (
	dl "GoWithTests/DependencyInjection"
	"net/http"
)

func main() {

	//log.Fatal(http.ListenAndServe("localhost:8080", http.HandlerFunc(MyGreetHandler)))

}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	dl.Greet(w, "world")
}
