package Core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type StoicRequest struct {
	*http.Request
}

type StoicResponse struct {
	http.ResponseWriter
}

func (response *StoicResponse) SetError(msg string) {
	response.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(response, "%s", msg)
}

func readRequestBody(r *StoicRequest) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic("")
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return body
}

func (r *StoicRequest) GetParamMap() map[string]interface{} {
	paramMap := make(map[string]interface{})

	if r.Method == "GET" {
		for key, values := range r.URL.Query() {
			if len(values) > 0 {
				paramMap[key] = values[0]
			}
		}
		return paramMap
	}

	if r.Method == "POST" {
		contentType := r.Header.Get("Content-Type")
		body := readRequestBody(r)

		if contentType == "application/json" {
			var requestBody map[string]interface{}
			if err := json.Unmarshal(body, &requestBody); err != nil {
				panic(fmt.Sprintf("error parsing JSON body: %s", err))
			}
			for key, value := range requestBody {
				paramMap[key] = value
			}

		} else if contentType == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err != nil {
				panic(fmt.Sprintf("error parsing form: %s", err))
			}
			for key, values := range r.PostForm {
				if len(values) > 0 {
					paramMap[key] = values[0]
				}
			}
		} else {
			panic(fmt.Sprintf("unsupported content type: %s", contentType))
		}

		return paramMap
	}

	panic(fmt.Sprintf("unsupported HTTP method: %s", r.Method))
}

func (r *StoicRequest) Has(name string) bool {
	params := r.GetParamMap()
	_, exists := params[name]
	return exists
}

func (r *StoicRequest) PrintRequestData() {
	params := r.GetParamMap()

	fmt.Println("Request Parameters:")
	for key, value := range params {
		fmt.Printf("  %s: %v\n", key, value)
	}
}

func (r *StoicRequest) HasAll(args ...string) bool {
	params := r.GetParamMap()
	for _, name := range args {
		if _, exists := params[name]; !exists {
			return false
		}
	}
	return true
}

func (r *StoicRequest) GetStringParam(name string) string {
	return r.GetParamMap()[name].(string)
}

func (r *StoicRequest) GetIntParam(name string) int {
	raw := r.GetStringParam(name)
	value, err := strconv.Atoi(raw)
	if err != nil {
		panic(fmt.Sprintf("invalid int value for parameter %s: %s", name, raw))
	}
	return value
}

func (r *StoicRequest) GetBoolParam(name string) bool {
	raw := r.GetStringParam(name)
	value, err := strconv.ParseBool(raw)
	if err != nil {
		panic(fmt.Sprintf("invalid bool value for parameter %s: %s", name, raw))
	}
	return value
}

func (r *StoicRequest) GetFloatParam(name string) float64 {
	raw := r.GetStringParam(name)
	value, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid float value for parameter %s: %s", name, raw))
	}
	return value
}
