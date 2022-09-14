package main

import "net/http"

func main() {
	//fmt.Println("hello world!")
	h := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello ravi"))
	}
	http.HandleFunc("/hello", h)

	//http.Handle("/hello", h)
	http.ListenAndServe(":8000", nil)

}
