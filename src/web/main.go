package main

import(
  "log"
  "encoding/json"
  "net/http"
  "runtime"
)

type Response struct {
  Status    string
  Runtime   string
  Arch      string
}

func indexHandler( w http.ResponseWriter, r *http.Request){
  log.Println("indexHandler")
  res := Response{"OK", runtime.GOOS, runtime.GOARCH}

  js, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func main(){
	log.SetPrefix("web_server:")
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":8080",nil)
}
