// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/SidingsMedia/dns-control/config"
	"github.com/SidingsMedia/dns-control/dnscontrol"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func init() {
	config.CliFlags.ConfigFilePath = flag.String("config", "config.yaml", "path to configuration file")
	config.CliFlags.ShowVersion = flag.Bool("version", false, "show current command version")
}

func main() {
	flag.Parse()
	buildinfo, haveBuildInfo := debug.ReadBuildInfo()

	if *config.CliFlags.ShowVersion {
		if haveBuildInfo {
			fmt.Printf("dns-config %s (%s)", buildinfo.Main.Version, buildinfo.GoVersion)
		}
		return
	}

	if haveBuildInfo {
		slog.Info("Starting dns-config", "version", buildinfo.Main.Version, "go", buildinfo.GoVersion)
	} else {
		slog.Info("Starting dns-config")
	}

	conf, err := config.ReadConfigFile(*config.CliFlags.ConfigFilePath)

	if err != nil {
		slog.Error("Failed to read configuration file", "error", err)
		return
	}

	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(sloggin.New(slog.Default()))
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	dnscontrol.NewController(
		engine,
		dnscontrol.NewService(
			dnscontrol.NewRepository(conf.Servers),
		),
	)

	// Set trusted proxies. If user has set it to * then we can just
	// ignore it as GIN trusts all by default
	if conf.TrustedProxies[0] != "*" {
		if err := engine.SetTrustedProxies(conf.TrustedProxies); err != nil {
			slog.Error("Failed to set trusted proxies", "error", err)
		}
		slog.Info("Set trusted proxies", "proxies", conf.TrustedProxies)
	} else {
		slog.Warn("You trusted all proxies, this is NOT safe. We recommend you to set a value.", "trustedProxies", conf.TrustedProxies)
	}

	slog.Info("Starting server", "bind", conf.BindAddr)
	engine.Run(conf.BindAddr)
}
