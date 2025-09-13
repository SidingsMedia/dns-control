package dnscontrol

import (
	"github.com/SidingsMedia/dns-control/config"
	"github.com/SidingsMedia/dns-control/dnscontrol/domain"
)

type Repository interface {
	GetServers() []domain.Server
}

type repository struct {
	servers []config.Server
}

func (r *repository) GetServers() []domain.Server {
	servers := make([]domain.Server, len(r.servers))
	for i, server := range r.servers {
		servers[i] = domain.Server{Name: server.Name, Target: server.Target}
	}

	return servers
}

func NewRepository(servers []config.Server) Repository {
	repository := &repository{
		servers: servers,
	}
	return repository
}
