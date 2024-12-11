package Router

import (
	"encoding/json"
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
	"net/http"
	"reflect"
)

type StoicResponse struct {
	http.ResponseWriter
}

func (response *StoicResponse) SetError(msg string) {
	response.WriteHeader(http.StatusInternalServerError)
	_, err := fmt.Fprintf(response, "%s", msg)
	if err != nil {
		Utility.Assert(false)
	}
}

func (response *StoicResponse) SetData(data any) {
	contentType := "text/plain"

	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		contentType = "application/json"
		response.Header().Set("Content-Type", contentType)

		jsonData, err := json.Marshal(data)
		if err != nil {
			Utility.AssertOnErrorMsg(err, "failed to marshal data to JSON")
		}
		_, err = response.Write(jsonData)
		if err != nil {
			Utility.Assert(false)
		}
		return
	default:
		contentType = "text/plain"
		response.Header().Set("Content-Type", contentType)

		_, err := fmt.Fprintf(response, "%+v", data)
		if err != nil {
			Utility.Assert(false)
		}
	}
}
