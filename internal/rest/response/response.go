package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
)

// WriteError provides writing of error
func WriteError(w http.ResponseWriter, r *http.Request, statusCode int, msg restModel.Error) {
	w.WriteHeader(statusCode)
	out, _ := json.Marshal(msg)
	render.JSON(w, r, restModel.Response{
		Data: (json.RawMessage)(out),
	})
}
