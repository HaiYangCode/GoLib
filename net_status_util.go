package GoLib

import (
	"encoding/json"
	"io"
	"net/http"
)

var result = make(map[string]interface{})

func RStatus(w http.ResponseWriter, statusCode int, statusMsg string, resultData interface{}) {
	result["code"] = statusCode
	result["msg"] = statusMsg
	result["data"] = resultData
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(result)
	writeString, err := io.WriteString(w, string(data))
	if err != nil {
		println(writeString)
		return
	}
}
