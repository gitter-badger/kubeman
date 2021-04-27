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

// Code generated by apiregister-gen. DO NOT EDIT.

package aggregation

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	InternalPodAggregation = builders.NewInternalResource(
		"podaggregations",
		"PodAggregation",
		func() runtime.Object { return &PodAggregation{} },
		func() runtime.Object { return &PodAggregationList{} },
	)

	ApiVersion = builders.NewApiGroup("aggregation.federated.bkbcs.tencent.com").WithKinds(
		InternalPodAggregation,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
	}).AddToScheme

	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)


// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// PodAggregation Functions and Structs
//
// +k8s:deepcopy-gen=false
type PodAggregationStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type PodAggregationStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodAggregation struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   corev1.PodSpec
	Status corev1.PodStatus
}

type PodAggregationList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PodAggregation
}

func (pc *PodAggregation) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *PodAggregation) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc PodAggregation) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store PodAggregation.
// +k8s:deepcopy-gen=false
type PodAggregationRegistry interface {
	ListPodAggregations(ctx context.Context, options *internalversion.ListOptions) (*PodAggregationList,
		error)
	GetPodAggregation(ctx context.Context, id string, options *metav1.GetOptions) (*PodAggregation, error)
	CreatePodAggregation(ctx context.Context, id *PodAggregation) (*PodAggregation, error)
	UpdatePodAggregation(ctx context.Context, id *PodAggregation) (*PodAggregation, error)
	DeletePodAggregation(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewPodAggregationRegistry(sp builders.StandardStorageProvider) PodAggregationRegistry {
	return &storagePodAggregation{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storagePodAggregation struct {
	builders.StandardStorageProvider
}

func (s *storagePodAggregation) ListPodAggregations(ctx context.Context,
	options *internalversion.ListOptions) (*PodAggregationList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*PodAggregationList), err
}

func (s *storagePodAggregation) GetPodAggregation(ctx context.Context, id string,
	options *metav1.GetOptions) (*PodAggregation, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*PodAggregation), nil
}

func (s *storagePodAggregation) CreatePodAggregation(ctx context.Context,
	object *PodAggregation) (*PodAggregation, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*PodAggregation), nil
}

func (s *storagePodAggregation) UpdatePodAggregation(ctx context.Context,
	object *PodAggregation) (*PodAggregation, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*PodAggregation), nil
}

func (s *storagePodAggregation) DeletePodAggregation(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}