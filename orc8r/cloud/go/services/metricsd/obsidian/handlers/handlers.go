/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package handlers

import (
	"os"

	"magma/orc8r/cloud/go/obsidian/handlers"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

// GetObsidianHandlers returns all obsidian handlers for metricsd
func GetObsidianHandlers() []handlers.Handler {
	client, err := api.NewClient(api.Config{Address: os.Getenv(prometheusAddressEnv)})
	if err != nil {
		return []handlers.Handler{
			{Path: queryURL, Methods: handlers.GET, HandlerFunc: getInitErrorHandler(err)},
			{Path: queryRangeURL, Methods: handlers.GET, HandlerFunc: getInitErrorHandler(err)},
		}
	}
	pAPI := v1.NewAPI(client)
	return []handlers.Handler{
		{Path: queryURL, Methods: handlers.GET, HandlerFunc: getPrometheusQueryHandler(pAPI)},
		{Path: queryRangeURL, Methods: handlers.GET, HandlerFunc: getPrometheusQueryRangeHandler(pAPI)},
	}
}
