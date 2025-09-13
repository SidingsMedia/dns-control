// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"github.com/SidingsMedia/dns-control/dnscontrol/model"
	"github.com/jinzhu/copier"
)

type Service interface {
	ListServers() model.List[model.Server]
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

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
