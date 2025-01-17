/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplications implements ApplicationInterface
type FakeApplications struct {
	Fake *FakeBkbcsV2
	ns   string
}

var applicationsResource = schema.GroupVersionResource{Group: "bkbcs", Version: "v2", Resource: "applications"}

var applicationsKind = schema.GroupVersionKind{Group: "bkbcs", Version: "v2", Kind: "Application"}

// Get takes name of the application, and returns the corresponding application object, and an error if there is any.
func (c *FakeApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationsResource, c.ns, name), &v2.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Application), err
}

// List takes label and field selectors, and returns the list of Applications that match those selectors.
func (c *FakeApplications) List(ctx context.Context, opts v1.ListOptions) (result *v2.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationsResource, applicationsKind, c.ns, opts), &v2.ApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.ApplicationList{ListMeta: obj.(*v2.ApplicationList).ListMeta}
	for _, item := range obj.(*v2.ApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applications.
func (c *FakeApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationsResource, c.ns, opts))

}

// Create takes the representation of a application and creates it.  Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Create(ctx context.Context, application *v2.Application, opts v1.CreateOptions) (result *v2.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationsResource, c.ns, application), &v2.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Application), err
}

// Update takes the representation of a application and updates it. Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Update(ctx context.Context, application *v2.Application, opts v1.UpdateOptions) (result *v2.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationsResource, c.ns, application), &v2.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Application), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplications) UpdateStatus(ctx context.Context, application *v2.Application, opts v1.UpdateOptions) (*v2.Application, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationsResource, "status", c.ns, application), &v2.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Application), err
}

// Delete takes name of the application and deletes it. Returns an error if one occurs.
func (c *FakeApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationsResource, c.ns, name), &v2.Application{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.ApplicationList{})
	return err
}

// Patch applies the patch and returns the patched application.
func (c *FakeApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationsResource, c.ns, name, pt, data, subresources...), &v2.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Application), err
}
