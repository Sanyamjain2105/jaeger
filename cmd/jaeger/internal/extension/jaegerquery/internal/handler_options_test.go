// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore/mocks"
)

func TestHandlerOptions(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		logger := zap.NewNop()
		handler := &APIHandler{}
		opt := HandlerOptions.Logger(logger)
		opt(handler)
		assert.Equal(t, logger, handler.logger)
	})

	t.Run("BasePath", func(t *testing.T) {
		handler := &APIHandler{}
		opt := HandlerOptions.BasePath("/api")
		opt(handler)
		assert.Equal(t, "/api", handler.basePath)
	})

	t.Run("Prefix", func(t *testing.T) {
		handler := &APIHandler{}
		opt := HandlerOptions.Prefix("/v1")
		opt(handler)
		assert.Equal(t, "/v1", handler.apiPrefix)
	})

	t.Run("QueryLookbackDuration", func(t *testing.T) {
		handler := &APIHandler{queryParser: queryParser{}}
		duration := 48 * time.Hour
		opt := HandlerOptions.QueryLookbackDuration(duration)
		opt(handler)
		assert.Equal(t, duration, handler.queryParser.traceQueryLookbackDuration)
	})

	t.Run("Tracer", func(t *testing.T) {
		handler := &APIHandler{}
		tracerProvider := noop.NewTracerProvider()
		opt := HandlerOptions.Tracer(tracerProvider)
		opt(handler)
		assert.Equal(t, tracerProvider, handler.tracer)
	})

	t.Run("MetricsQueryService", func(t *testing.T) {
		handler := &APIHandler{}
		metricsReader := &mocks.Reader{}
		opt := HandlerOptions.MetricsQueryService(metricsReader)
		opt(handler)
		assert.Equal(t, metricsReader, handler.metricsQueryService)
	})
}
