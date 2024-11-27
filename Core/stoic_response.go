package Core

import (
	"fmt"
	"net/http"
)

type StoicResponse struct {
	http.ResponseWriter
}

func (response *StoicResponse) SetError(msg string) {
	response.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(response, "%s", msg)
}
