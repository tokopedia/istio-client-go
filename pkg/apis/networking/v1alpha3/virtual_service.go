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

// VirtualService is a Istio VirtualService resource
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VirtualServiceSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualServiceList is a list of VirtualService resources
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualService `json:"items"`
}

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *VirtualServiceSpec) DeepCopyInto(out *VirtualServiceSpec) {
	*out = *in
}

type VirtualServiceSpec struct {
	Hosts []string `json:"hosts,omitempty"`

	Gateways []string `json:"gateways,omitempty"`

	Http []*HTTPRoute `json:"http,omitempty"`

	Tls []*TLSRoute `json:"tls,omitempty"`

	Tcp []*TCPRoute `json:"tcp,omitempty"`
}

type Destination struct {
	Host string `json:"host,omitempty"`

	Subset string `json:"subset,omitempty"`

	Port *PortSelector `json:"port,omitempty"`
}

type HTTPRoute struct {
	Match []*HTTPMatchRequest `json:"match,omitempty"`

	Route []*DestinationWeight `json:"route,omitempty"`

	Redirect *HTTPRedirect `json:"redirect,omitempty"`

	Rewrite *HTTPRewrite `json:"rewrite,omitempty"`

	WebsocketUpgrade bool `json:"websocketUpgrade,omitempty"`

	Timeout string `json:"timeout,omitempty"`

	Retries *HTTPRetry `json:"retries,omitempty"`

	Mirror *Destination `json:"mirror,omitempty"`

	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`

	AppendHeaders map[string]string `json:"append_headers,omitempty"`

	RemoveResponseHeaders []string `json:"removeResponseHeaders,omitempty"`
}

type TLSRoute struct {
	Match []*TLSMatchAttributes `json:"match,omitempty"`

	Route []*DestinationWeight `json:"route,omitempty"`
}

type TCPRoute struct {
	Match []*L4MatchAttributes `json:"match,omitempty"`

	Route []*DestinationWeight `json:"route,omitempty"`
}

type HTTPMatchRequest struct {
	Uri *StringMatch `json:"uri,omitempty"`

	Scheme *StringMatch `json:"scheme,omitempty"`

	Method *StringMatch `json:"method,omitempty"`

	Authority *StringMatch `json:"authority,omitempty"`

	Headers map[string]*StringMatch `json:"headers,omitempty"`

	Port uint32 `json:"port,omitempty"`

	SourceLabels map[string]string `json:"sourceLabels"`

	Gateways []string `json:"gateways,omitempty"`
}

type DestinationWeight struct {
	Destination *Destination `json:"destination,omitempty"`

	Weight int32 `json:"weight,omitempty"`
}

type L4MatchAttributes struct {
	DestinationSubnets []string `json:"destinationSubnets,omitempty"`

	Port uint32 `json:"port,omitempty"`

	SourceSubnet string `json:"sourceSubnet,omitempty"`

	SourceLabels map[string]string `json:"sourceLabels,omitempty"`

	Gateways []string `json:"gateways,omitempty"`
}

type TLSMatchAttributes struct {
	SniHosts []string `json:"sniHosts,omitempty"`

	DestinationSubnets []string `json:"destinationSubnets,omitempty"`

	Port uint32 `json:"port,omitempty"`

	SourceSubnet string `json:"sourceSubnet,omitempty"`

	SourceLabels map[string]string `json:"sourceLabels,omitempty"`

	Gateways []string `json:"gateways,omitempty"`
}

type HTTPRedirect struct {
	Uri string `json:"uri,omitempty"`

	Authority string `json:"authority,omitempty"`
}

type HTTPRewrite struct {
	Uri string `json:"uri,omitempty"`

	Authority string `json:"authority,omitempty"`
}

type StringMatch struct {
	Exact  *string `json:"exact,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	Regex  *string `json:"regex,omitempty"`
}

type HTTPRetry struct {
	Attempts int32 `json:"attempts,omitempty"`

	PerTryTimeout string `json:"perTryTimeout,omitempty"`
}

type CorsPolicy struct {
	AllowOrigin []string `json:"allowOrigin,omitempty"`

	AllowMethods []string `json:"allowMethods,omitempty"`

	AllowHeaders []string `json:"allowHeaders,omitempty"`

	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	MaxAge string `json:"maxAge,omitempty"`

	AllowCredentials bool `json:"allowCredentials,omitempty"`
}

type PortSelector struct {
	Number *int32 `json:"number,omitempty"`
}
