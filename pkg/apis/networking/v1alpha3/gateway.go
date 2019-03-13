package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GatewaySpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Gateway `json:"items"`
}

type GatewaySpec struct {
	Servers  []*Server         `json:"servers,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
}

func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
}

type Server struct {
	Port  *Port              `json:"port,omitempty"`
	Hosts []string           `json:"hosts,omitempty"`
	Tls   *Server_TLSOptions `json:"tls,omitempty"`
}

type Server_TLSOptions struct {
	HttpsRedirect     bool     `json:"httpsRedirect,omitempty"`
	Mode              string   `json:"mode,omitempty"`
	ServerCertificate string   `json:"serverCertificate,omitempty"`
	PrivateKey        string   `json:"privateKey,omitempty"`
	CaCertificates    string   `json:"caCertificates,omitempty"`
	SubjectAltNames   []string `json:"subjectAltNames,omitempty"`
}

type Port struct {
	Number   uint32 `json:"number,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name,omitempty"`
}
