package elastic

import "time"

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
