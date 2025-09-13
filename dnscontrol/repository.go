// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/SidingsMedia/dns-control/config"
	"github.com/SidingsMedia/dns-control/dnscontrol/domain"
	"github.com/jinzhu/copier"
)

type Repository interface {
	GetServers() []domain.Server
	GetCache(domain string, servers []string) (map[string]domain.CacheResult, error)
}

type repository struct {
	servers   []config.Server
	serverMap map[string]config.Server
}

// Return a list of all configured servers
func (r *repository) GetServers() []domain.Server {
	servers := make([]domain.Server, len(r.servers))
	copier.Copy(&servers, &r.servers)
	return servers
}

func (r *repository) formatListCacheUrl(servers []string, searchDomain string) (urls []string, err error) {
	for _, id := range servers {
		server, exists := r.serverMap[id]
		if !exists {
			return nil, ErrServerNotFound
		}

		urls = append(urls, server.Target+"/api/cache/list?token="+server.Token+"&domain="+searchDomain)
	}
	return urls, nil
}

// Process the HTTP response from the server and return the resulting CacheResult.
func processCacheResponse(response *http.Response) (*domain.CacheResult, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		slog.Error("Response from server was not 200 OK", "requestUrl", response.Request.URL, "code", response.StatusCode, "body", body)
		return nil, ErrStatusNotOk
	}

	var result domain.CacheResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	slog.Info("Processed response", "result", result)

	return &result, nil
}

// Get the cached results for a set of servers
func (r *repository) GetCache(searchDomain string, servers []string) (map[string]domain.CacheResult, error) {
	urls, err := r.formatListCacheUrl(servers, searchDomain)

	if err != nil {
		return nil, err
	}

	results := make(chan (struct {
		id       string
		response *http.Response
		err      error
	}), len(urls))

	for i, url := range urls {
		go func(url string, index int) {
			res, err := http.Get(url)
			results <- struct {
				id       string
				response *http.Response
				err      error
			}{
				id: servers[index], response: res, err: err,
			}
		}(url, i)
	}

	cache := make(map[string]domain.CacheResult)

	for range urls {
		result := <-results
		if result.err != nil {
			slog.Error("Failed to make request", "server", result.id, "error", result.err)
			return nil, result.err
		}

		response, err := processCacheResponse(result.response)
		if err != nil {
			return nil, err
		}

		cache[result.id] = *response

		if response.Status != "ok" {
			slog.Error(
				"Got an error from Technetium DNS",
				"error", response.ErrorMessage,
				"trace", response.StackTrace,
				"innerMessage", response.InnerErrorMessage,
			)

			return nil, errors.New(response.ErrorMessage)
		}
	}

	return cache, nil
}

func NewRepository(servers []config.Server) Repository {
	repository := &repository{
		servers:   servers,
		serverMap: make(map[string]config.Server),
	}

	for i := range repository.servers {
		repository.serverMap[servers[i].Id] = servers[i]
	}

	return repository
}
