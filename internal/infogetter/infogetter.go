package infogetter

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/fx"
)

type PlayerInfoGetter struct {
	cfg *Config

	requestQuota chan struct{}
}

type Config struct {
	fx.In

	CLI     *http.Client
	SiteURL string

	// N requests per period
	RateLimit  int
	RatePeriod time.Duration
}

func NewPlayerInfoGetter(lc fx.Lifecycle, cfg *Config) *PlayerInfoGetter {
	p := &PlayerInfoGetter{
		cfg: cfg,
	}

	rateLimitDone := make(chan struct{})
	lc.Append(fx.StartStopHook(
		func() { p.runRateLimiter(rateLimitDone) },
		func() { close(rateLimitDone) },
	))
	return p
}

func (p *PlayerInfoGetter) runRateLimiter(done <-chan struct{}) {
	tick := time.NewTicker(p.cfg.RatePeriod)
	for {
		select {
		case <-done:
			return
		case <-tick.C:
			for i := 0; i < p.cfg.RateLimit; i++ {
				select {
				case p.requestQuota <- struct{}{}:
				default:
					break
				}
			}
		}
	}
}

func (p *PlayerInfoGetter) Get(ctx context.Context, name string) (*PlayerInfo, error) {
	select {
	case <-p.requestQuota:
	case <-ctx.Done():
		return nil, errors.Wrap(ctx.Err(), "context error")
	}

	info, err := p.get(ctx, name)
	return info, errors.Wrap(err, "get")
}

func (p *PlayerInfoGetter) get(ctx context.Context, name string) (*PlayerInfo, error) {
	uri, err := url.Parse(p.cfg.SiteURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse site URL")
	}
	uri.RawQuery = url.Values{"nickname": []string{name}}.Encode()

	req := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
	}
	req = req.WithContext(ctx)

	resp, err := p.cfg.CLI.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	return &data.PlayerInfo, nil
}
