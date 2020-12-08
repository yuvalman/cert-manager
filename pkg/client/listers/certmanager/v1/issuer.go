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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IssuerLister helps list Issuers.
// All objects returned here must be treated as read-only.
type IssuerLister interface {
	// List lists all Issuers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Issuer, err error)
	// Issuers returns an object that can list and get Issuers.
	Issuers(namespace string) IssuerNamespaceLister
	IssuerListerExpansion
}

// issuerLister implements the IssuerLister interface.
type issuerLister struct {
	indexer cache.Indexer
}

// NewIssuerLister returns a new IssuerLister.
func NewIssuerLister(indexer cache.Indexer) IssuerLister {
	return &issuerLister{indexer: indexer}
}

// List lists all Issuers in the indexer.
func (s *issuerLister) List(selector labels.Selector) (ret []*v1.Issuer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Issuer))
	})
	return ret, err
}

// Issuers returns an object that can list and get Issuers.
func (s *issuerLister) Issuers(namespace string) IssuerNamespaceLister {
	return issuerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IssuerNamespaceLister helps list and get Issuers.
// All objects returned here must be treated as read-only.
type IssuerNamespaceLister interface {
	// List lists all Issuers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Issuer, err error)
	// Get retrieves the Issuer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Issuer, error)
	IssuerNamespaceListerExpansion
}

// issuerNamespaceLister implements the IssuerNamespaceLister
// interface.
type issuerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Issuers in the indexer for a given namespace.
func (s issuerNamespaceLister) List(selector labels.Selector) (ret []*v1.Issuer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Issuer))
	})
	return ret, err
}

// Get retrieves the Issuer from the indexer for a given namespace and name.
func (s issuerNamespaceLister) Get(name string) (*v1.Issuer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("issuer"), name)
	}
	return obj.(*v1.Issuer), nil
}
