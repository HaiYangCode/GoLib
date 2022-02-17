package atlib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
				w.Header().Set("Content-Type", "application/json")
				io.WriteString(w, string(data))

			}
		}()
		hf(w, r)
	}
}
