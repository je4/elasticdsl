package elastic

import "net/http"

func NewFakeElasticRoundTripper(originalTransport http.RoundTripper) FakeElasticRoundTripper {
	return FakeElasticRoundTripper{
		originalTransport: originalTransport,
	}
}

// FakeElasticRoundTripper is used for Elasticsearch 7.x instances which do not send
// the required "X-Elastic-Product" header. This Problem seems to be solved on Elasticsearch 8.x
type FakeElasticRoundTripper struct {
	originalTransport http.RoundTripper
}

func (t FakeElasticRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := t.originalTransport.RoundTrip(r)
	resp.Header.Set("X-Elastic-Product", "Elasticsearch")
	return resp, err
}
