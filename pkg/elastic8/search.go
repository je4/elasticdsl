package elastic8

import (
	"context"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

func (es8 *Client) Search(cfg *elastic.SearchConfig) ([]map[string][]string, map[string]any, int64, elastic.ResultFacet, error) {
	res, err := es8.es.Search(
		es8.es.Search.WithContext(context.Background()),
		es8.es.Search.WithIndex(cfg.Index),
		es8.es.Search.WithTrackTotalHits(true),
		es8.es.Search.WithPretty(),
	)
	_ = res
	_ = err
	return nil, nil, 0, nil, nil
}
