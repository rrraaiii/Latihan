package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Saya ganti baru")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at localhost:9017")
http.ListenAndServe(":9017", nil)

}
