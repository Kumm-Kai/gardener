// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupBucket holds details about backup bucket
type BackupBucket struct {
	metav1.TypeMeta
	// Standard object metadata.
	metav1.ObjectMeta
	// Specification of the Backup Bucket.
	Spec BackupBucketSpec
	// Most recently observed status of the Backup Bucket.
	Status BackupBucketStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupBucketList is a list of BackupBucket objects.
type BackupBucketList struct {
	metav1.TypeMeta
	// Standard list object metadata.
	metav1.ListMeta
	// Items is the list of BackupBucket.
	Items []BackupBucket
}

// BackupBucketSpec is the specification of a Backup Bucket.
type BackupBucketSpec struct {
	// Provider holds the details of cloud provider of the object store.
	Provider BackupBucketProvider
	// SecretRef is a reference to a secret that contains the credentials to access object store.
	SecretRef corev1.SecretReference
	// Seed holds the name of the seed allocated to BackupBucket for running controller.
	Seed *string
}

// BackupBucketStatus holds the most recently observed status of the Backup Bucket.
type BackupBucketStatus struct {
	// LastOperation holds information about the last operation on the BackupBucket.
	LastOperation *LastOperation
	// LastError holds information about the last occurred error during an operation.
	LastError *LastError
	// ObservedGeneration is the most recent generation observed for this BackupBucket. It corresponds to the
	// BackupBucket's generation, which is updated on mutation by the API Server.
	ObservedGeneration int64
}

// BackupBucketProvider holds the details of cloud provider of the object store.
type BackupBucketProvider struct {
	// Type is the type of provider.
	Type string
	// Region is the region of the bucket.
	Region string
}
