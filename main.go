package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", func( resp http.ResponseWriter, req *http.Request){
    resp.write([]byte("erdal was here!"))
  })
  
  http.ListenAndServe(":8080", nil)
}
