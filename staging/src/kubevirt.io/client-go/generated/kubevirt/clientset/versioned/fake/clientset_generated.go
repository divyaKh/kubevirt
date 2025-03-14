/*
Copyright 2022 The KubeVirt Authors.

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

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kubevirt.io/client-go/generated/kubevirt/clientset/versioned"
	clonev1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/clone/v1alpha1"
	fakeclonev1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/clone/v1alpha1/fake"
	exportv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/export/v1alpha1"
	fakeexportv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/export/v1alpha1/fake"
	flavorv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/flavor/v1alpha1"
	fakeflavorv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/flavor/v1alpha1/fake"
	migrationsv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/migrations/v1alpha1"
	fakemigrationsv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/migrations/v1alpha1/fake"
	poolv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/pool/v1alpha1"
	fakepoolv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/pool/v1alpha1/fake"
	snapshotv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/snapshot/v1alpha1"
	fakesnapshotv1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/snapshot/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// CloneV1alpha1 retrieves the CloneV1alpha1Client
func (c *Clientset) CloneV1alpha1() clonev1alpha1.CloneV1alpha1Interface {
	return &fakeclonev1alpha1.FakeCloneV1alpha1{Fake: &c.Fake}
}

// ExportV1alpha1 retrieves the ExportV1alpha1Client
func (c *Clientset) ExportV1alpha1() exportv1alpha1.ExportV1alpha1Interface {
	return &fakeexportv1alpha1.FakeExportV1alpha1{Fake: &c.Fake}
}

// FlavorV1alpha1 retrieves the FlavorV1alpha1Client
func (c *Clientset) FlavorV1alpha1() flavorv1alpha1.FlavorV1alpha1Interface {
	return &fakeflavorv1alpha1.FakeFlavorV1alpha1{Fake: &c.Fake}
}

// MigrationsV1alpha1 retrieves the MigrationsV1alpha1Client
func (c *Clientset) MigrationsV1alpha1() migrationsv1alpha1.MigrationsV1alpha1Interface {
	return &fakemigrationsv1alpha1.FakeMigrationsV1alpha1{Fake: &c.Fake}
}

// PoolV1alpha1 retrieves the PoolV1alpha1Client
func (c *Clientset) PoolV1alpha1() poolv1alpha1.PoolV1alpha1Interface {
	return &fakepoolv1alpha1.FakePoolV1alpha1{Fake: &c.Fake}
}

// SnapshotV1alpha1 retrieves the SnapshotV1alpha1Client
func (c *Clientset) SnapshotV1alpha1() snapshotv1alpha1.SnapshotV1alpha1Interface {
	return &fakesnapshotv1alpha1.FakeSnapshotV1alpha1{Fake: &c.Fake}
}
