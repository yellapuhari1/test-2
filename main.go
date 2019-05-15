package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func main() {
    var router = mux.NewRouter()
    router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

    headersOk := handlers.AllowedHeaders([]string{"Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    fmt.Println("Running server!")
    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
   
   // fetch last commit sha
    var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"log", "--pretty=format:'%h'", "-n", "1"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	
    type commit_detail struct {
        Version string `json:"version"`
        Description string `json:"description"`
        Lastcommitsha string `json:"lastcommitsha"`
    }
    
    type commits = []commit_detail
    
    response := commits{commit_detail{Version: "1.0", Description: "pre-interview technical test", Lastcommitsha: sha}}


    type status_1 struct {
        MyApplication commits `json:"myapplication"`
    }
    
    finalresponse := status_1{MyApplication: response}
    
    json.NewEncoder(w).Encode(finalresponse)
}