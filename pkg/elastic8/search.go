package elastic8

import (
	"bytes"
	"context"
	"emperror.dev/errors"
	dsl "github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
	"io"
)

func (es8 *Client) Search(cfg *elastic.SearchConfig) ([]map[string][]string, map[string]any, int64, elastic.ResultFacet, error) {
	q := dsl.NewQueryObject()
	matchAll := dsl.NewMatchAllQuery()
	data, err := q(q.WithQuery(matchAll(matchAll.WithBoost(0.5))))
	if err != nil {
		return nil, nil, 0, nil, errors.Wrapf(err, "cannot marshal query %v", q)
	}
	res, err := es8.es.Search(
		es8.es.Search.WithContext(context.Background()),
		es8.es.Search.WithIndex(cfg.Index),
		es8.es.Search.WithTrackTotalHits(true),
		es8.es.Search.WithPretty(),
		es8.es.Search.WithBody(bytes.NewBuffer(data)),
	)
	defer res.Body.Close()
	resultData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, 0, nil, errors.Wrap(err, "cannot get result body")
	}
	_ = resultData
	return nil, nil, 0, nil, nil
}
