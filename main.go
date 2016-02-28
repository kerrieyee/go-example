package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

type book struct {
    Title  string `json:"title"`
    Author string `json:"author name"`
    Age    int `json:"age,omitempty"`
}

func main() {
    http.HandleFunc("/", showBooks)
    http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
    b := book{"Building Web Apps with Go", "Jeremy Saenz", 5}
    b2 := book{ "adfad", "J", 0}
    w.Header().Set("Content-Type", "application/json")
    
    enc := json.NewEncoder(w)
    if err := enc.Encode(b); err != nil {
        fmt.Println("Error Encoding", err)
    }
    
    if err := enc.Encode(b2); err != nil {
        fmt.Println("Error Encoding", err)
    }
    

    // js, err := json.Marshal(b)
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }

    // w.Header().Set("Content-Type", "application/json")
    // w.Write(js)
}