package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
)

const ApiEndpoint = "/api/objects"
const ListenPort = "8000"

var store = map[string]*json.RawMessage{}

type jsonFormat struct {
    Id string
    Data *json.RawMessage
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
    log.Print("Got an API invocation")

    switch method := req.Method; method {
        case http.MethodPost:
            var data jsonFormat

            decoder := json.NewDecoder(req.Body)
            err := decoder.Decode(&data)
            if err != nil {
                log.Print("Failed to parse the data!")
                res.WriteHeader(http.StatusInternalServerError)
                return
            }

            defer req.Body.Close()

            if data.Id == "" {
                res.WriteHeader(http.StatusInternalServerError)
                log.Print("ID field was missing!")
                return
            }

            fmt.Println("Saving item:", data.Id)
            store[data.Id] = data.Data

        case http.MethodGet:
            runes := []rune(req.URL.Path)
            itemId := string(runes[len(ApiEndpoint) + 1:])

            fmt.Println("Getting item:", itemId)

            dataString, err := json.Marshal(store[itemId])
            if err != nil {
                res.WriteHeader(http.StatusInternalServerError)
                log.Print("Failed to serialize the data!")
                return
            }

            if store[itemId] == nil {
                res.WriteHeader(http.StatusNotFound)
                log.Print("Could not find the item!")
                return
            }

            fmt.Fprintf(res, "%s", dataString)

        case http.MethodDelete:
            runes := []rune(req.URL.Path)
            itemId := string(runes[len(ApiEndpoint) + 1:])
            fmt.Println("Deleting item:", itemId)

            if store[itemId] == nil {
                res.WriteHeader(http.StatusNotFound)
                log.Print("Could not find the item to delete!")
                return
            }

            delete(store, itemId)
    }
}

func main() {
    http.HandleFunc(ApiEndpoint, apiHandler)
    http.HandleFunc(ApiEndpoint + "/", apiHandler)

    log.Printf("Staring to listen on port %s...", ListenPort)
    err := http.ListenAndServe(":" + ListenPort, nil)
    if err != nil {
        log.Fatal("Failed to start the server: ", err)
    }
}
