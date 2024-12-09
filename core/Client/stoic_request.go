package Client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/superg3m/stoic-go/core/Utility"
	"io"
	"net/http"
)

type StoicRequest struct {
	*http.Request
}

func readRequestBody(r *StoicRequest) []byte {
	body, err := io.ReadAll(r.Body)
	Utility.AssertOnError(err)

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
	if r.Request.Method == "OPTIONS" {
		return true
	}

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
	return Utility.CastAny[int](r.GetStringParam(name))
}

func (r *StoicRequest) GetBoolParam(name string) bool {
	return Utility.CastAny[bool](r.GetStringParam(name))
}

func (r *StoicRequest) GetFloatParam(name string) float64 {
	return Utility.CastAny[float64](r.GetStringParam(name))
}

func (r *StoicRequest) GetJsonParam(name string, target any) {
	paramMap := r.GetParamMap()

	value, exists := paramMap[name]
	Utility.Assert(exists)

	jsonData, err := json.Marshal(value)
	Utility.AssertOnErrorMsg(err, fmt.Sprintf("failed to marshal parameter %q: %s", name, err))

	err2 := json.Unmarshal(jsonData, target)
	Utility.AssertOnErrorMsg(err2, fmt.Sprintf("failed to unmarshal parameter %q into target type: %s", name, err2))
}
