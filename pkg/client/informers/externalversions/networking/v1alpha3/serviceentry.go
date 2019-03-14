/*
Portions Copyright 2019 The Kubernetes Authors.
Portions Copyright 2019 Aspen Mesh Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	time "time"

	networkingv1alpha3 "github.com/tokopedia/istio-client-go/pkg/apis/networking/v1alpha3"
	versioned "github.com/tokopedia/istio-client-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/tokopedia/istio-client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/tokopedia/istio-client-go/pkg/client/listers/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceEntryInformer provides access to a shared informer and lister for
// ServiceEntries.
type ServiceEntryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.ServiceEntryLister
}

type serviceEntryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceEntryInformer constructs a new informer for ServiceEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceEntryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceEntryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceEntryInformer constructs a new informer for ServiceEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceEntryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().ServiceEntries(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().ServiceEntries(namespace).Watch(options)
			},
		},
		&networkingv1alpha3.ServiceEntry{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceEntryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceEntryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceEntryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha3.ServiceEntry{}, f.defaultInformer)
}

func (f *serviceEntryInformer) Lister() v1alpha3.ServiceEntryLister {
	return v1alpha3.NewServiceEntryLister(f.Informer().GetIndexer())
}
