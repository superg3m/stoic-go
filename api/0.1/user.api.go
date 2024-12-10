package __1

import (
	"github.com/superg3m/stoic-go/Core/Client"
	"github.com/superg3m/stoic-go/Core/Server"
)

func sendUserMetrics(r *Client.StoicRequest, w Server.StoicResponse) {
	if !r.Has("data") {
		w.SetError("No user data")
	}

	data := r.GetParamMap()
	_ = data

	// w.SetData(data)
}

func init() {
	Server.RegisterApiEndpoint("User/Metric", sendUserMetrics, "POST")
}
