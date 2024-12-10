package __1

import (
	"github.com/superg3m/stoic-go/Core/Router"
)

func sendUserMetrics(r *Router.StoicRequest, w Router.StoicResponse) {
	if !r.Has("data") {
		w.SetError("No user data")
	}

	data := r.GetParamMap()
	_ = data

	// w.SetData(data)
}

func init() {
	Router.RegisterApiEndpoint("User/Metric", sendUserMetrics, "POST")
}
