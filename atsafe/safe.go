package atsafe

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

func SafeHandle(hf http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				fmt.Println("ok", ok, e)
				var result = make(map[string]interface{})
				result["code"] = 300
				result["msg"] = e.Error()
				data, _ := json.Marshal(result)
				//http.Error(w,string(data),http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				io.WriteString(w, string(data))
				//log.Println(string(debug.Stack()))
				fmt.Println("WARN:panic in %v-%v", hf, e)
			}
		}()
		hf(w, r)
	}
}
