package elastic8

import (
	"emperror.dev/errors"
	"encoding/json"
	"time"
)

type ResultInfo struct {
	Name        string `json:"name,omitempty"`
	ClusterName string `json:"cluster_name,omitempty"`
	ClusterUUID string `json:"cluster_uuid,omitempty"`
	Version     struct {
		Number                           string    `json:"number,omitempty"`
		BuildFlavor                      string    `json:"build_flavor,omitempty"`
		BuildType                        string    `json:"build_type,omitempty"`
		BuildHash                        string    `json:"build_hash,omitempty"`
		BuildDate                        time.Time `json:"build_date,omitempty"`
		BuildSnapshot                    bool      `json:"build_snapshot,omitempty"`
		LuceneVersion                    string    `json:"lucene_version,omitempty"`
		MinimumWireCompatibilityVersion  string    `json:"minimum_wire_compatibility_version,omitempty"`
		MinimumIndexCompatibilityVersion string    `json:"minimum_index_compatibility_version,omitempty"`
	} `json:"version,omitempty"`
	Tagline string `json:"tagline,omitempty"`
}

func (es8 *Client) Info() (*ResultInfo, error) {
	// 1. Get cluster info
	//
	var r = &ResultInfo{}
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
