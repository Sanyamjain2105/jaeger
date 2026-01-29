// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/proto-gen/api_v2/metrics"
)

func TestDefaultParameters(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected interface{}
	}{
		{
			name:     "defaultDependencyLookbackDuration",
			value:    defaultDependencyLookbackDuration,
			expected: time.Hour * 24,
		},
		{
			name:     "defaultTraceQueryLookbackDuration",
			value:    defaultTraceQueryLookbackDuration,
			expected: time.Hour * 24 * 2,
		},
		{
			name:     "defaultMetricsQueryLookbackDuration",
			value:    defaultMetricsQueryLookbackDuration,
			expected: time.Hour,
		},
		{
			name:     "defaultMetricsQueryStepDuration",
			value:    defaultMetricsQueryStepDuration,
			expected: 5 * time.Second,
		},
		{
			name:     "defaultMetricsQueryRateDuration",
			value:    defaultMetricsQueryRateDuration,
			expected: 10 * time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value,
				"%s should have the correct default value", tt.name)
		})
	}
}

func TestDefaultMetricsSpanKinds(t *testing.T) {
	assert.NotNil(t, defaultMetricsSpanKinds, "defaultMetricsSpanKinds should not be nil")
	assert.Len(t, defaultMetricsSpanKinds, 1, "defaultMetricsSpanKinds should have exactly one element")
	assert.Equal(t, metrics.SpanKind_SPAN_KIND_SERVER.String(), defaultMetricsSpanKinds[0],
		"defaultMetricsSpanKinds should contain SPAN_KIND_SERVER")
}
