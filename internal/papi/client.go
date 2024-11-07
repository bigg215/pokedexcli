package papi

import (
	"net/http"
	"time"

	"github.com/bigg215/pokedexcli/internal/papicache"
)

type Client struct {
	cache      papicache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: papicache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
