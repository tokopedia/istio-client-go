/*
Portions Copyright 2017 The Kubernetes Authors.
Portions Copyright 2018 Aspen Mesh Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DestinationRule is a Istio DestinationRule resource
type DestinationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DestinationRuleSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DestinationRuleList is a list of DestinationRule resources
type DestinationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DestinationRule `json:"items"`
}

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *DestinationRuleSpec) DeepCopyInto(out *DestinationRuleSpec) {
	*out = *in
}

type DestinationRuleSpec struct {
	Host          string         `json:"host,omitempty"`
	TrafficPolicy *TrafficPolicy `json:"trafficPolicy,omitempty"`
	Subsets       []*Subset      `json:"subsets,omitempty"`
}

type TrafficPolicy struct {
	LoadBalancer      *LoadBalancerSettings             `json:"loadBalancer,omitempty"`
	ConnectionPool    *ConnectionPoolSettings           `json:"connectionPool,omitempty"`
	OutlierDetection  *OutlierDetection                 `json:"outlierDetection,omitempty"`
	Tls               *TLSSettings                      `json:"tls,omitempty"`
	PortLevelSettings []*TrafficPolicyPortTrafficPolicy `json:"portLevelSettings,omitempty"`
}

// Traffic policies that apply to specific ports of the service
type TrafficPolicyPortTrafficPolicy struct {
	Port             *PortSelector           `json:"port,omitempty"`
	LoadBalancer     *LoadBalancerSettings   `json:"loadBalancer,omitempty"`
	ConnectionPool   *ConnectionPoolSettings `json:"connectionPool,omitempty"`
	OutlierDetection *OutlierDetection       `json:"outlierDetection,omitempty"`
	Tls              *TLSSettings            `json:"tls,omitempty"`
}

type Subset struct {
	Name          string            `json:"name,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	TrafficPolicy *TrafficPolicy    `json:"trafficPolicy,omitempty"`
}

type LoadBalancerSettings struct {
	Simple         *string                               `json:"simple,omitempty"`
	ConsistentHash *LoadBalancerSettingsConsistentHashLB `json:"consistentHash,omitempty"`
}

type LoadBalancerSettingsConsistentHashLB struct {
	HttpHeaderName  *string                                         `json:"httpHeaderName,omitempty"`
	HttpCookie      *LoadBalancerSettingsConsistentHashLBHTTPCookie `json:"httpCookie,omitempty"`
	UseSourceIp     *bool                                           `json:"useSourceIp,omitempty"`
	MinimumRingSize uint64                                          `json:"minimumRingSize,omitempty"`
}
type LoadBalancerSettingsConsistentHashLBHTTPCookie struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Ttl  string `json:"ttl,omitempty"`
}

type ConnectionPoolSettings struct {
	Tcp  *ConnectionPoolSettingsTCPSettings  `json:"tcp,omitempty"`
	Http *ConnectionPoolSettingsHTTPSettings `json:"http,omitempty"`
}

// Settings common to both HTTP and TCP upstream connections.
type ConnectionPoolSettingsTCPSettings struct {
	MaxConnections int32   `json:"maxConnections,omitempty"`
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
}

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type ConnectionPoolSettingsHTTPSettings struct {
	Http1MaxPendingRequests  int32 `json:"http1MaxPendingRequests,omitempty"`
	Http2MaxRequests         int32 `json:"http2MaxRequests,omitempty"`
	MaxRequestsPerConnection int32 `json:"maxRequestsPerConnection,omitempty"`
	MaxRetries               int32 `json:"maxRetries,omitempty"`
}

type OutlierDetection struct {
	ConsecutiveErrors  int32   `json:"consecutiveErrors,omitempty"`
	Interval           *string `json:"interval,omitempty"`
	BaseEjectionTime   *string `json:"baseEjectionTime,omitempty"`
	MaxEjectionPercent int32   `json:"maxEjectionPercent,omitempty"`
}

type TLSSettings struct {
	Mode              string   `json:"mode,omitempty"`
	ClientCertificate string   `json:"clientCertificate,omitempty"`
	PrivateKey        string   `json:"privateKey,omitempty"`
	CaCertificates    *string  `json:"caCertificates,omitempty"`
	SubjectAltNames   []string `json:"subjectAltNames,omitempty"`
	Sni               *string  `json:"sni,omitempty"`
}
