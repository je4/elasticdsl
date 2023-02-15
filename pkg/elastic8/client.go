package elastic8

import (
	"emperror.dev/errors"
	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/op/go-logging"
	"go.elastic.co/apm/module/apmelasticsearch"
	"net/http"
	"os"
	"time"
)

type Client struct {
	index string
	es    *elasticsearch.Client
	log   *logging.Logger
}

func NewClient(address string, index string, apikey string, logger *logging.Logger) (*Client, error) {
	var err error
	idx := &Client{
		index: index,
		es:    nil,
		log:   logger,
	}
	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//
	// Use a third-party package for implementing the backoff function
	//
	retryBackoff := backoff.NewExponentialBackOff()

	cfg := elasticsearch.Config{
		APIKey: apikey,
		Addresses: []string{
			address,
		},
		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		// ... using the "apmelasticsearch" wrapper for instrumentation
		Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},
		// Configure the backoff function
		//
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},

		// Retry up to 5 attempts
		//
		MaxRetries: 5,

		Logger: &elastictransport.ColorLogger{Output: os.Stdout},
	}

	idx.es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create client")
	}
	return idx, nil
}
