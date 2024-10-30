package main

import (
	dl "GoWithTests/DependencyInjection"
	mo "GoWithTests/Mocking"
	"net/http"
	"os"
)

func main() {

	//log.Fatal(http.ListenAndServe("localhost:8080", http.HandlerFunc(MyGreetHandler)))

	mo.CountDown(os.Stdout)

}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	dl.Greet(w, "world")
}
