// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"github.com/SidingsMedia/dns-control/dnscontrol/model"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jinzhu/copier"
)

type Service interface {
	ListServers() model.List[model.Server]
	GetCache(domain string, servers []string) (*model.CacheResponse, error)
}

type service struct {
	repository Repository
}

// func (service *Service) <Handler>(<model> *model.<Model>) error {
// 	// Handler logic here
// 	return nil
// }

func (s service) ListServers() model.List[model.Server] {
	servers := s.repository.GetServers()

	var responseServers []model.Server
	copier.Copy(&responseServers, &servers)

	return model.List[model.Server]{Results: responseServers}
}

func (s service) GetCache(domain string, servers []string) (*model.CacheResponse, error) {
	cache, err := s.repository.GetCache(domain, servers)
	if err != nil {
		return nil, err
	}

	type cacheKey struct {
		name string
		typ  string
	}

	combinedCache := make(map[cacheKey]*model.CacheEntry)
	zones := mapset.NewSet[string]()

	for _, server := range servers {
		zones.Append(cache[server].Response.Zones...)

		for _, entry := range cache[server].Response.Records {
			key := cacheKey{name: entry.Name, typ: entry.Type}
			if !(key.name == "" && key.typ == "") {
				if _, exists := combinedCache[key]; exists {
					combinedCache[key].CachedResult = append(
						combinedCache[key].CachedResult,
						model.CachedResult{
							Id:    server,
							RData: entry.RData,
							Ttl:   entry.Ttl,
						},
					)
				} else {
					combinedCache[key] = &model.CacheEntry{
						Name: entry.Name,
						Type: entry.Type,
						CachedResult: []model.CachedResult{model.CachedResult{
							Id:    server,
							RData: entry.RData,
							Ttl:   entry.Ttl,
						}},
					}
				}
			}
		}
	}

	response := model.CacheResponse{
		Entries:          make([]model.CacheEntry, len(combinedCache)),
		AvailableDomains: zones.ToSlice(),
	}

	i := 0
	for _, entry := range combinedCache {
		response.Entries[i] = *entry
		i++
	}

	return &response, nil
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
