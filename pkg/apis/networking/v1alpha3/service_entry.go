package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualService `json:"items"`
}

func (in *ServiceEntrySpec) DeepCopyInto(out *ServiceEntrySpec) {
	*out = *in
}

type ServiceEntrySpec struct {
	Hosts      []string                 `json:"hosts,omitempty"`
	Addresses  []string                 `json:"addresses,omitempty"`
	Ports      []*Port                  `json:"ports,omitempty"`
	Location   string                   `json:"location,omitempty"`
	Resolution string                   `json:"resolution,omitempty"`
	Endpoints  []*ServiceEntry_Endpoint `json:"endpoints,omitempty"`
}

type ServiceEntry_Endpoint struct {
	Address string            `json:"address,omitempty"`
	Ports   map[string]uint32 `json:"ports,omitempty"`
	Labels  map[string]string `json:"labels,omitempty"`
}
