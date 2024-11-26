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

func readRequestBody(r *StoicRequest) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func (r *StoicRequest) Has(name string) bool {
	body, err := readRequestBody(r)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return false
	}

	switch r.Method {
	case "GET":
		values := r.URL.Query()
		_, exists := values[name]
		return exists

	case "POST":
		contentType := r.Header.Get("Content-Type")

		if contentType == "application/json" {
			var requestBody map[string]interface{}
			err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&requestBody)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return false
			}
			_, exists := requestBody[name]
			return exists

		} else if contentType == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err != nil {
				fmt.Println("Error parsing form:", err)
				return false
			}
			_, exists := r.PostForm[name]
			return exists
		}

	default:
		fmt.Println("Unsupported method or content type.")
		return false
	}

	return true
}

func To[T any](value interface{}) (T, error) {
	var zero T // Default zero value of type T
	switch v := value.(type) {
	case T:
		return v, nil
	default:
		return zero, fmt.Errorf("unable to cast %v to %T", value, zero)
	}
}

func (r *StoicRequest) PrintRequestData() {
	body, err := readRequestBody(r)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Printf("Request Body: %s\n", string(body))

	if r.Method == "GET" {
		fmt.Println("Query Parameters:")
		for key, values := range r.URL.Query() {
			fmt.Printf("  %s: %v\n", key, values)
		}
	}

	if r.Method == "POST" {
		contentType := r.Header.Get("Content-Type")

		if contentType == "application/json" {
			var requestBody map[string]interface{}
			err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&requestBody)
			if err == nil {
				fmt.Println("JSON Body:")
				for key, value := range requestBody {
					fmt.Printf("  %s: %v\n", key, value)
				}
			} else {
				fmt.Println("Error parsing JSON body:", err)
			}

		} else if contentType == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err == nil {
				fmt.Println("Form Data:")
				for key, values := range r.PostForm {
					fmt.Printf("  %s: %v\n", key, values)
				}
			} else {
				fmt.Println("Error parsing form data:", err)
			}
		}
	}
}

func (r *StoicRequest) HasAll(args ...string) bool {
	for _, name := range args {
		if !r.Has(name) {
			return false
		}
	}

	return true
}

func (r *StoicRequest) GetStringParam(name string) (string, error) {
	body, err := readRequestBody(r)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return "", err
	}

	var strValue string
	switch r.Method {
	case "GET":
		strValue = r.URL.Query().Get(name)
	case "POST":
		var requestBody map[string]interface{}
		err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&requestBody)
		if err != nil {
			return "", err
		}

		strValue2, exists := requestBody[name]
		if !exists {
			return "", err
		}

		strValue, _ = To[string](strValue2)

	default:
		return "", fmt.Errorf("unsupported HTTP method")
	}

	if strValue == "" {
		return "", fmt.Errorf("parameter %s not found", name)
	}

	return strValue, nil
}

func (r *StoicRequest) GetIntParam(name string) (int, error) {
	strValue, err := r.GetStringParam(name)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %s to int", strValue)
	}

	return intValue, nil
}

func (r *StoicRequest) GetBoolParam(name string) (bool, error) {
	strValue, err := r.GetStringParam(name)
	if err != nil {
		return false, err
	}

	boolValue, err := strconv.ParseBool(strValue)
	if err != nil {
		return false, fmt.Errorf("cannot convert %s to bool", strValue)
	}

	return boolValue, nil
}
