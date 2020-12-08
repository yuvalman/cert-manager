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

package v1alpha3

import (
	"context"
	"time"

	v1alpha3 "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha3"
	scheme "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChallengesGetter has a method to return a ChallengeInterface.
// A group's client should implement this interface.
type ChallengesGetter interface {
	Challenges(namespace string) ChallengeInterface
}

// ChallengeInterface has methods to work with Challenge resources.
type ChallengeInterface interface {
	Create(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.CreateOptions) (*v1alpha3.Challenge, error)
	Update(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (*v1alpha3.Challenge, error)
	UpdateStatus(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (*v1alpha3.Challenge, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.Challenge, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.ChallengeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.Challenge, err error)
	ChallengeExpansion
}

// challenges implements ChallengeInterface
type challenges struct {
	client rest.Interface
	ns     string
}

// newChallenges returns a Challenges
func newChallenges(c *AcmeV1alpha3Client, namespace string) *challenges {
	return &challenges{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the challenge, and returns the corresponding challenge object, and an error if there is any.
func (c *challenges) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.Challenge, err error) {
	result = &v1alpha3.Challenge{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("challenges").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Challenges that match those selectors.
func (c *challenges) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.ChallengeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.ChallengeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("challenges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested challenges.
func (c *challenges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("challenges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a challenge and creates it.  Returns the server's representation of the challenge, and an error, if there is any.
func (c *challenges) Create(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.CreateOptions) (result *v1alpha3.Challenge, err error) {
	result = &v1alpha3.Challenge{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("challenges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(challenge).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a challenge and updates it. Returns the server's representation of the challenge, and an error, if there is any.
func (c *challenges) Update(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (result *v1alpha3.Challenge, err error) {
	result = &v1alpha3.Challenge{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("challenges").
		Name(challenge.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(challenge).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *challenges) UpdateStatus(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (result *v1alpha3.Challenge, err error) {
	result = &v1alpha3.Challenge{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("challenges").
		Name(challenge.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(challenge).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the challenge and deletes it. Returns an error if one occurs.
func (c *challenges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("challenges").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *challenges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("challenges").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched challenge.
func (c *challenges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.Challenge, err error) {
	result = &v1alpha3.Challenge{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("challenges").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
