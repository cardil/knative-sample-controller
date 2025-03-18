/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	samplesv1alpha1 "knative.dev/sample-controller/pkg/apis/samples/v1alpha1"
	scheme "knative.dev/sample-controller/pkg/client/clientset/versioned/scheme"
)

// AddressableServicesGetter has a method to return a AddressableServiceInterface.
// A group's client should implement this interface.
type AddressableServicesGetter interface {
	AddressableServices(namespace string) AddressableServiceInterface
}

// AddressableServiceInterface has methods to work with AddressableService resources.
type AddressableServiceInterface interface {
	Create(ctx context.Context, addressableService *samplesv1alpha1.AddressableService, opts v1.CreateOptions) (*samplesv1alpha1.AddressableService, error)
	Update(ctx context.Context, addressableService *samplesv1alpha1.AddressableService, opts v1.UpdateOptions) (*samplesv1alpha1.AddressableService, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, addressableService *samplesv1alpha1.AddressableService, opts v1.UpdateOptions) (*samplesv1alpha1.AddressableService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*samplesv1alpha1.AddressableService, error)
	List(ctx context.Context, opts v1.ListOptions) (*samplesv1alpha1.AddressableServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *samplesv1alpha1.AddressableService, err error)
	AddressableServiceExpansion
}

// addressableServices implements AddressableServiceInterface
type addressableServices struct {
	*gentype.ClientWithList[*samplesv1alpha1.AddressableService, *samplesv1alpha1.AddressableServiceList]
}

// newAddressableServices returns a AddressableServices
func newAddressableServices(c *SamplesV1alpha1Client, namespace string) *addressableServices {
	return &addressableServices{
		gentype.NewClientWithList[*samplesv1alpha1.AddressableService, *samplesv1alpha1.AddressableServiceList](
			"addressableservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *samplesv1alpha1.AddressableService { return &samplesv1alpha1.AddressableService{} },
			func() *samplesv1alpha1.AddressableServiceList { return &samplesv1alpha1.AddressableServiceList{} },
		),
	}
}
