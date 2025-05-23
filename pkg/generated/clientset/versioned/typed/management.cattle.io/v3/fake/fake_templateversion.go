/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplateVersions implements TemplateVersionInterface
type FakeTemplateVersions struct {
	Fake *FakeManagementV3
}

var templateversionsResource = v3.SchemeGroupVersion.WithResource("templateversions")

var templateversionsKind = v3.SchemeGroupVersion.WithKind("TemplateVersion")

// Get takes name of the templateVersion, and returns the corresponding templateVersion object, and an error if there is any.
func (c *FakeTemplateVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.TemplateVersion, err error) {
	emptyResult := &v3.TemplateVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(templateversionsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.TemplateVersion), err
}

// List takes label and field selectors, and returns the list of TemplateVersions that match those selectors.
func (c *FakeTemplateVersions) List(ctx context.Context, opts v1.ListOptions) (result *v3.TemplateVersionList, err error) {
	emptyResult := &v3.TemplateVersionList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(templateversionsResource, templateversionsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.TemplateVersionList{ListMeta: obj.(*v3.TemplateVersionList).ListMeta}
	for _, item := range obj.(*v3.TemplateVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templateVersions.
func (c *FakeTemplateVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(templateversionsResource, opts))
}

// Create takes the representation of a templateVersion and creates it.  Returns the server's representation of the templateVersion, and an error, if there is any.
func (c *FakeTemplateVersions) Create(ctx context.Context, templateVersion *v3.TemplateVersion, opts v1.CreateOptions) (result *v3.TemplateVersion, err error) {
	emptyResult := &v3.TemplateVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(templateversionsResource, templateVersion, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.TemplateVersion), err
}

// Update takes the representation of a templateVersion and updates it. Returns the server's representation of the templateVersion, and an error, if there is any.
func (c *FakeTemplateVersions) Update(ctx context.Context, templateVersion *v3.TemplateVersion, opts v1.UpdateOptions) (result *v3.TemplateVersion, err error) {
	emptyResult := &v3.TemplateVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(templateversionsResource, templateVersion, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.TemplateVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTemplateVersions) UpdateStatus(ctx context.Context, templateVersion *v3.TemplateVersion, opts v1.UpdateOptions) (result *v3.TemplateVersion, err error) {
	emptyResult := &v3.TemplateVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(templateversionsResource, "status", templateVersion, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.TemplateVersion), err
}

// Delete takes name of the templateVersion and deletes it. Returns an error if one occurs.
func (c *FakeTemplateVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(templateversionsResource, name, opts), &v3.TemplateVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplateVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(templateversionsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.TemplateVersionList{})
	return err
}

// Patch applies the patch and returns the patched templateVersion.
func (c *FakeTemplateVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.TemplateVersion, err error) {
	emptyResult := &v3.TemplateVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(templateversionsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.TemplateVersion), err
}
