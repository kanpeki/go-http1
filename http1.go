package main

import (
        "encoding/json"
        "net/http"
        "time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        enc :=  json.NewEncoder(w)
        v := map[string]string {
                "now" : time.Now().Format(time.RFC3339Nano),
        }
        if err := enc.Encode(v); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                enc.Encode(map[string]string {
                        "error" : err.Error(),
                })
                return
        }
}

func main() {
        http.HandleFunc("/time", getTime)
        http.ListenAndServe(":8080", nil)
}