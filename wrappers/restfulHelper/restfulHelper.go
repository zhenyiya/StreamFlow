package restfulHelper

import (
	"encoding/json"
	"github.com/zhenyiya/artifacts/restful"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/utils"
	"io"
	"net/http"
)

func SendErrorWith(w http.ResponseWriter, errPayload restful.ErrorPayload, header constants.Header) error {
	mal, err := json.Marshal(errPayload)
	if err != nil {
		return err
	}
	utils.AdaptHTTPWithHeader(w, header)
	io.WriteString(w, string(mal))
	return nil
}
