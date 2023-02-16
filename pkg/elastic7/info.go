package elastic7

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

func (es8 *Client) Info() (*elastic.ResultInfo, error) {
	// 1. Get cluster info
	//
	var r = &elastic.ResultInfo{}
	res, err := es8.es.Info()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting info  response")
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		return nil, errors.Errorf("info response has error %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.Wrap(err, "cannot pars info result body")
	}
	return r, nil
}
