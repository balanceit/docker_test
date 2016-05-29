package main

import(
  "fmt"
  "log"
  "net/http"
  "runtime"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
  log.Println("indexHandler")
  fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS,runtime.GOARCH)
}

func main(){
	log.SetPrefix("web_server:")
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":8080",nil)
}
