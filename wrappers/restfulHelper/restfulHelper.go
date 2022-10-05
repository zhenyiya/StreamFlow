package restfulHelper

import (
	"encoding/json"
	"github.com/zhenyiya/artifacts/restful"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/utils"
	"io"
	"net/http"
)

func SendErrorWith(w http.ResponseWriter, errPayload restful.ErrorPayload, status int) error {
	mal, err := json.Marshal(errPayload)
	if err != nil {
		return err
	}
	utils.AdaptHTTPWithHeader(w, constants.HeaderContentTypeJSON)
	utils.AdaptHTTPWithStatus(w, status)
	io.WriteString(w, string(mal))
	return nil
}
