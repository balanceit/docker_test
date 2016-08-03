package main

import(
  "log"
  "os"
  "fmt"
  "encoding/json"
  "net/http"
  "runtime"
  //"github.com/rubenv/sql-migrate"
//	"database/sql"

//	_ "github.com/lib/pq"
)

type Response struct {
  Status    string
  Runtime   string
  Arch      string
  Tables    []string
}




func indexHandler() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
    log.Println("indexHandler")

    tables, err := listTables(client)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    log.Println(tables)
    //res := Response{"OK", runtime.GOOS, runtime.GOARCH, tables }
    res := Response{"OK", runtime.GOOS, runtime.GOARCH, []string{"testing"} }

    js, err := json.Marshal(res)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    // w.WriteHeader(400)
    w.Write(js)
  }
}

func main(){

  fmt.Println("connection string: ", os.Getenv("DB_CONNECTION_STRING"))
  //client, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
  //if err != nil {
  //  log.Fatal("cannot connect to database: ", err)
  //}
  //log.Println("connected to database")

  //if err2 := client.Ping(); err2 != nil {
  // log.Println("Failed to keep connection alive", err2)
  //}

  log.Println("before running migrations")
  //n, err := migrate.Exec(client, "postgres", migrations(), migrate.Up)
  //if err != nil {
  //    log.Fatal("db migrations failed: ", err)
  //}
  //log.Println(n, "migrations run")

  // log.SetPrefix("web_server:")
  log.Println("before root handler")
  http.HandleFunc("/", indexHandler())
  log.Println("before listen and serve")
  //http.ListenAndServe(":8080",nil)
  log.Fatal(http.ListenAndServe(":8080", nil))
  log.Println("running on port 8080")

}
