package twitter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kynrai/sentiment-cli/pkg/twitter/models"
	"go.opencensus.io/plugin/ochttp"
)

type repo struct {
	httpClient    *http.Client
	twitterKey    string
	twitterSecret string
}

type Option func(*repo)

func TwitterKey(key string) Option {
	return func(r *repo) {
		r.twitterKey = key
	}
}

func TwitterSecret(secret string) Option {
	return func(r *repo) {
		r.twitterSecret = secret
	}
}

func New(opts ...Option) *repo {
	r := &repo{}
	r.httpClient = &http.Client{
		Transport: &ochttp.Transport{},
		Timeout:   time.Second * 1,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *repo) token(ctx context.Context) (*models.TokenResp, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.twitter.com/oauth2/token",
		strings.NewReader(`grant_type=client_credentials`),
	)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(r.twitterKey, r.twitterSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error calling token, code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	tr := &models.TokenResp{}
	return tr, json.NewDecoder(resp.Body).Decode(tr)
}

func (r *repo) Tweets30Days(ctx context.Context, query string, maxRes int, from, to time.Time) (*models.Tweets30DayResp, error) {
	type payload struct {
		Query      string `json:"query"`
		MaxResults string `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	}
	data := payload{
		Query:      "from:Monzo lang:en",
		MaxResults: strconv.Itoa(maxRes),
		FromDate:   twitterTime(from),
		ToDate:     twitterTime(to),
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/tweets/search/30day/dev.json", &buf)
	if err != nil {
		return nil, err
	}
	token, err := r.token(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tr models.Tweets30DayResp
	return &tr, json.NewDecoder(resp.Body).Decode(&tr)
}

// converts a time.Time to a twitter time format
func twitterTime(t time.Time) string {
	return t.Format("200601021504")
}
