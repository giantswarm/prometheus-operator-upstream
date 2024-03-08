// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

// ScrapeClassApplyConfiguration represents an declarative configuration of the ScrapeClass type for use
// with apply.
type ScrapeClassApplyConfiguration struct {
	Name             *string                       `json:"name,omitempty"`
	Default          *bool                         `json:"default,omitempty"`
	TLSConfig        *TLSConfigApplyConfiguration  `json:"tlsConfig,omitempty"`
	ExtraRelabelings []*monitoringv1.RelabelConfig `json:"extraRelabelings,omitempty"`
}

// ScrapeClassApplyConfiguration constructs an declarative configuration of the ScrapeClass type for use with
// apply.
func ScrapeClass() *ScrapeClassApplyConfiguration {
	return &ScrapeClassApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ScrapeClassApplyConfiguration) WithName(value string) *ScrapeClassApplyConfiguration {
	b.Name = &value
	return b
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *ScrapeClassApplyConfiguration) WithDefault(value bool) *ScrapeClassApplyConfiguration {
	b.Default = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *ScrapeClassApplyConfiguration) WithTLSConfig(value *TLSConfigApplyConfiguration) *ScrapeClassApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithExtraRelabelings adds the given value to the ExtraRelabelings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExtraRelabelings field.
func (b *ScrapeClassApplyConfiguration) WithExtraRelabelings(values ...**monitoringv1.RelabelConfig) *ScrapeClassApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExtraRelabelings")
		}
		b.ExtraRelabelings = append(b.ExtraRelabelings, *values[i])
	}
	return b
}
