/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	schedulerconfig "k8s.io/kube-scheduler/config/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoschedulingArgs defines the scheduling parameters for Coscheduling plugin.
type CoschedulingArgs struct {
	metav1.TypeMeta `json:",inline"`

	// PermitWaitingTime is the wait timeout in seconds.
	PermitWaitingTimeSeconds *int64 `json:"permitWaitingTimeSeconds,omitempty"`
	// PodGroupGCInterval is the period to run gc of PodGroup in seconds.
	PodGroupGCIntervalSeconds *int64 `json:"podGroupGCIntervalSeconds,omitempty"`
	// If the deleted PodGroup stays longer than the PodGroupExpirationTime,
	// the PodGroup will be deleted from PodGroupInfos.
	PodGroupExpirationTimeSeconds *int64 `json:"podGroupExpirationTimeSeconds,omitempty"`
}

// modes type.
type ModeType string

const (
	Least ModeType = "Least"
	Most  ModeType = "Most"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourcesAllocatableArgs holds arguments used to configure NodeResourcesAllocatable plugin.
type NodeResourcesAllocatableArgs struct {
	metav1.TypeMeta `json:",inline"`

	// Resources to be considered when scoring.
	// Allowed weights start from 1.
	// An example resource set might include "cpu" (millicores) and "memory" (bytes)
	// with weights of 1<<20 and 1 respectfully. That would mean 1 MiB has equivalent
	// weight as 1 millicore.
	Resources []schedulerconfig.ResourceSpec `json:"resources,omitempty"`

	// Whether to prioritize nodes with least or most allocatable resources.
	Mode *ModeType `json:"mode,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CapacitySchedulingArgs defines the scheduling parameters for CapacityScheduling plugin.
type CapacitySchedulingArgs struct {
	metav1.TypeMeta `json:",inline"`

	// KubeConfigPath is the path of kubeconfig.
	KubeConfigPath *string `json:"kubeconfigpath,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BestFitBinPackArgs holds arguments used to configure BestFitBinPack plugin.
type BestFitBinPackArgs struct {
	metav1.TypeMeta 	`json:",inline"`

	// Node target CPU Utilisation for bin packing
	TargetCPUUtilization float64 `json:"mode,omitempty"`
	// Default CPU requests to use for best effort QoS
	DefaultCPURequests   int64 `json:"mode,omitempty"`
}

func (b BestFitBinPackArgs) DeepCopyObject() runtime.Object {
	panic("implement me")
}

var _ runtime.Object = &BestFitBinPackArgs{}