package main

import (
  "fmt"
  "net/http"
)

func ignore(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(""))
}

func main() {
  srcDir := "./src"
  ignoreDir := "/ignore"
  
  http.Handle("/", http.FileServer(http.Dir(srcDir)))
  http.HandleFunc(ignoreDir, ignore)
  fmt.Println("Server is up and running, serving " + srcDir + ", ignoring " + ignoreDir + ".")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
