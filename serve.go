package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
)

func ignore(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(""))
}

func banner() {
    b, err := ioutil.ReadFile("goservelogo.txt")
    if err != nil {
        panic(err)
    }
	fmt.Println(string(b))
	fmt.Println("-----------------------------------------------------------------------------------------------------------------")
}

func main() {
  banner()

  data, err := ioutil.ReadFile("goserve.conf")
  parsed := strings.Split(string(data), ";")

  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fmt.Println(parsed)

  port := parsed[0]
  srcDir := ".." + parsed[1]
  ignoreDir := ".." + parsed[2]

  http.Handle("/", http.FileServer(http.Dir(srcDir)))
  http.HandleFunc(ignoreDir, ignore)
  fmt.Println("Server is up and running, serving " + srcDir + ", ignoring " + ignoreDir + ".")
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
