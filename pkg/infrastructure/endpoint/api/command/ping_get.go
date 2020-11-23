package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mlambda-net/template/pkg/application/message"
	"github.com/mlambda-net/template/pkg/infrastructure/endpoint/api/model"
	"net/http"
	"strconv"
)

// GetPing godoc
// @Summary Get the ping
// @Produce json
// @Param id path string true "search by id"
// @Success 200 {object} model.Dummy
// @Failure 500 {string} string "Internal error"
// @Router /ping/{id} [get]
func (c *control) getPing(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi( mux.Vars(r)["id"])
	//token := r.Header.Get("Authorization")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		result, e := c.template.Request(&message.Ping{
			Id: int64(id),
		}).Unwrap()

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			pong := result.(*message.Pong)
			_ = json.NewEncoder(w).Encode(&model.Dummy{Value: pong.Value})
		}
	}
}
