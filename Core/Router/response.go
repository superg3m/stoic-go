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
	Utility.ErrorHandler
}

func (response *StoicResponse) SetData(data any) {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		response.Header().Set("Content-Type", "application/json")

		jsonData, err := json.Marshal(data)
		Utility.AssertOnErrorMsg(err, "failed to marshal data to JSON")

		_, err = response.Write(jsonData)
		Utility.AssertOnErrorMsg(err, "failed to write data to JSON")
		return
	default:
		response.Header().Set("Content-Type", "text/plain")

		_, err := fmt.Fprintf(response, "%+v", data)
		Utility.AssertOnError(err)
	}
}