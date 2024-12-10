package Router

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
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

	// should use reflection to determine if the header response needs to be text, json an other stuff

	response.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(response, "%+v", data)
	if err != nil {
		Utility.Assert(false)
	}
}
