package Server

import (
	"fmt"
	"github.com/superg3m/stoic-go/cmd/src/Utility"
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
