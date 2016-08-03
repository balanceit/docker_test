package main

import(
  "log"
  "os"
  "fmt"
  "encoding/json"
  "net/http"
  "runtime"
  "github.com/rubenv/sql-migrate"
	"database/sql"

	_ "github.com/lib/pq"
)

type Response struct {
  Status    string
  Runtime   string
  Arch      string
  Tables    []string
}

func listTables(c *sql.DB) ([]string, error){
  var tables []string

  rows, err := c.Query(`select table_name from information_schema.tables where table_schema= 'public';`)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var table string
    err = rows.Scan(&table)
    if err != nil {
      return nil, err
    }
    tables = append(tables, table)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }
  return tables, nil
}

func migrations() migrate.MigrationSource {
        return &migrate.AssetMigrationSource{
                Asset:    Asset,
                AssetDir: AssetDir,
                Dir:      "db/migrations",
        }
}

func indexHandler(client *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
    log.Println("indexHandler")

    tables, err := listTables(client)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    log.Println(tables)
    res := Response{"OK", runtime.GOOS, runtime.GOARCH, tables }
    // res := Response{"OK", runtime.GOOS, runtime.GOARCH, []string{"testing"} }

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
  client, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
  if err != nil {
    log.Fatal("cannot connect to database: ", err)
  }
  log.Println("connected to database")

  log.Println("before ping")
  if err2 := client.Ping(); err2 != nil {
   log.Println("Failed to keep connection alive", err2)
  }
  log.Println("after ping")

  log.Println("before running migrations")
  n, err := migrate.Exec(client, "postgres", migrations(), migrate.Up)
  if err != nil {
      log.Fatal("db migrations failed: ", err)
  }
  log.Println(n, "migrations run")

  // log.SetPrefix("web_server:")
  log.Println("before root handler")
  http.HandleFunc("/", indexHandler(client))
  log.Println("before listen and serve")
  http.ListenAndServe(":8080",nil)
  //log.Fatal(http.ListenAndServe(":8080", nil))
  log.Println("running on port 8080")

}
