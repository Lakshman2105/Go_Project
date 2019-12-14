
package main

import (
    "fmt"
    "log"
    "net/http"
    //"strconv"

    "github.com/gorilla/mux"
)

var name []string
func contains(arr []string, str string) bool {
    for _, a := range arr {
       if a == str {
          return true
       }
    }
    return false
 }
  

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post clled"}`))
}

// get function for disaplay name. 
func params(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")


    if val, ok := pathParams["username"]; ok {
        
        w.Write([]byte(fmt.Sprintf(`{"uname": "%s" }`, val)))
    }
       
}

// post function for create name 
func create(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
      // var input[] string  
    if a, ok := pathParams["name"]; ok {
    name = append(name,a)
    fmt.Println(name)
    w.WriteHeader(http.StatusCreated)
    }
}

// compare the name and display in web (GET)
func comp(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")
    if val1, ok := pathParams["valid_name"]; ok { 
       // w.Write([]byte(fmt.Sprintf(`{"valid_name": "%s" }`)))
      result := contains(name,val1)
       if result {
       w.WriteHeader(http.StatusOK)
        } else
       {
           w.WriteHeader(http.StatusNotFound)      
         }
           fmt.Println(result)
    
}
 
}

func main() {
    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("", get).Methods(http.MethodGet)
    api.HandleFunc("", post).Methods(http.MethodPost)
    api.HandleFunc("/acc/{name}", create).Methods(http.MethodPost)// post comment for add name
    api.HandleFunc("/list_the_name/{valid_name}",comp).Methods(http.MethodGet)//display name from post
    log.Fatal(http.ListenAndServe(":8083", r))
}