/*
Copyright 2020 The cert-manager contributors.

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

package scheme

import (
	acmev1 "github.com/jetstack/cert-manager/pkg/apis/acme/v1"
	acmev1alpha2 "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha2"
	acmev1alpha3 "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha3"
	acmev1beta1 "github.com/jetstack/cert-manager/pkg/apis/acme/v1beta1"
	certmanagerv1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	certmanagerv1alpha3 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha3"
	certmanagerv1beta1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	acmev1alpha2.AddToScheme,
	acmev1alpha3.AddToScheme,
	acmev1beta1.AddToScheme,
	acmev1.AddToScheme,
	certmanagerv1alpha2.AddToScheme,
	certmanagerv1alpha3.AddToScheme,
	certmanagerv1beta1.AddToScheme,
	certmanagerv1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
