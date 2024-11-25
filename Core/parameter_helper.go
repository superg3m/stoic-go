package Core

import (
	"net/http"
	"strconv"
)

type StoicRequest struct {
	*http.Request
}

type StoicResponse struct {
	http.ResponseWriter
}

func (r *StoicRequest) Has(name string) bool {
	switch r.Method {
	case "GET":
		values := r.URL.Query()
		_, exists := values[name]
		return exists
	case "POST":
		if err := r.ParseForm(); err != nil {
			return false
		}
		data := r.FormValue(name)
		return data != ""
	default:
		return false
	}
}

// GetParam retrieves the parameter and attempts to cast it to the appropriate type
func (r *StoicRequest) GetParam(name string) interface{} {
	var strValue string

	switch r.Method {
	case "GET":
		strValue = r.URL.Query().Get(name)
	case "POST":
		if err := r.ParseForm(); err != nil {
			return nil
			// Error
		}

		strValue = r.FormValue(name)
	default:
		return nil
	}

	if strValue == "" {
		return nil
		// Error!
	}

	if intValue, err := strconv.Atoi(strValue); err == nil {
		return intValue
	}

	if boolValue, err := strconv.ParseBool(strValue); err == nil {
		return boolValue
	}

	return strValue
}
