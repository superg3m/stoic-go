package Server

import (
	"fmt"
	"github.com/superg3m/stoic-go/core/Utility"
	"net/http"
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
	response.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(response, "%+v", data)
	if err != nil {
		Utility.Assert(false)
	}
}
