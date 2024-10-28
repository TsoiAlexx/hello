package main

import "github.com/TsoiAllex/hello"

func main() {
	hello.Main()
	//fmt.Println("Listening on http://localhost:9001")
	//http.ListenAndServe(":9001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	p := hello.NewPrinter()
	//	p.Output = w
	//	p.Print()
	//}))

}
