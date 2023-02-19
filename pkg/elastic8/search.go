package elastic8

import (
	"bytes"
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
	"io"
)

func (es8 *Client) Search(index string, srch any) (*elastic.SearchResult, error) {
	/*
		dsl := es8.API
		search := dsl.Search(
			dsl.Search.WithQuery(
				dsl.Query(
					dsl.BoolQuery(
						dsl.BoolQuery.WithMust(
							dsl.MatchAllQuery(
								dsl.MatchAllQuery.WithBoost(0.5),
							),
						),
						dsl.BoolQuery.WithMinimumShouldMatch(
							&dsl.MinimumShouldMatch{IntValue: 1},
						),
				),
				),
			),
			dsl.Search.WithAggs(
				dsl.Aggs(),
			),
			dsl.Search.WithIndicesBoost(
				dsl.IndicesBoost(
					dsl.IndicesBoost.AppendIndex(cfg.Index, 2.1),
				),
			),
		)
	*/
	data, err := json.Marshal(srch)
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal query")
	}
	res, err := es8.es.Search(
		es8.es.Search.WithContext(context.Background()),
		es8.es.Search.WithIndex(index),
		es8.es.Search.WithTrackTotalHits(true),
		//		es8.es.Search.WithPretty(),
		es8.es.Search.WithBody(bytes.NewBuffer(data)),
	)
	defer res.Body.Close()
	resultData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get result body")
	}
	var result = &elastic.SearchResult{}
	if err := json.Unmarshal(resultData, result); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal result - %s", string(resultData))
	}
	return result, nil
}
