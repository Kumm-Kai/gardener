// Copyright 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package reference_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	operatorv1alpha1 "github.com/gardener/gardener/pkg/apis/operator/v1alpha1"
	. "github.com/gardener/gardener/pkg/operator/controller/garden/reference"
)

var _ = Describe("Add", func() {
	Describe("#Predicate", func() {
		var garden *operatorv1alpha1.Garden

		BeforeEach(func() {
			garden = &operatorv1alpha1.Garden{
				Spec: operatorv1alpha1.GardenSpec{
					VirtualCluster: operatorv1alpha1.VirtualCluster{
						Kubernetes: operatorv1alpha1.Kubernetes{
							KubeAPIServer: &operatorv1alpha1.KubeAPIServerConfig{
								KubeAPIServerConfig: &gardencorev1beta1.KubeAPIServerConfig{},
							},
						},
						Gardener: operatorv1alpha1.Gardener{
							APIServer: &operatorv1alpha1.GardenerAPIServerConfig{},
						},
					},
				},
			}
		})

		It("should return false because new object is no garden", func() {
			Expect(Predicate(nil, nil)).To(BeFalse())
		})

		It("should return false because old object is no garden", func() {
			Expect(Predicate(nil, garden)).To(BeFalse())
		})

		It("should return false because there is no ref change", func() {
			Expect(Predicate(garden, garden)).To(BeFalse())
		})

		It("should return true because the kube-apiserver audit policy field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Kubernetes.KubeAPIServer.AuditConfig = &gardencorev1beta1.AuditConfig{AuditPolicy: &gardencorev1beta1.AuditPolicy{ConfigMapRef: &corev1.ObjectReference{Name: "audit-policy"}}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the gardener-apiserver audit policy field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Gardener.APIServer.AuditConfig = &gardencorev1beta1.AuditConfig{AuditPolicy: &gardencorev1beta1.AuditPolicy{ConfigMapRef: &corev1.ObjectReference{Name: "audit-policy"}}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the kube-apiserver audit webhook secret field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Kubernetes.KubeAPIServer.AuditWebhook = &operatorv1alpha1.AuditWebhook{KubeconfigSecretName: "webhook-secret"}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the gardener-apiserver audit webhook secret field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Gardener.APIServer.AuditWebhook = &operatorv1alpha1.AuditWebhook{KubeconfigSecretName: "webhook-secret"}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the ETCD backup secret field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.ETCD = &operatorv1alpha1.ETCD{Main: &operatorv1alpha1.ETCDMain{Backup: &operatorv1alpha1.Backup{SecretRef: corev1.LocalObjectReference{Name: "secret-name"}}}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the SNI secret field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Kubernetes.KubeAPIServer.SNI = &operatorv1alpha1.SNI{SecretName: "secret-sni"}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the authentication webhook secret field changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Kubernetes.KubeAPIServer.Authentication = &operatorv1alpha1.Authentication{Webhook: &operatorv1alpha1.AuthenticationWebhook{KubeconfigSecretName: "auth-secret"}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the kube-apiserver admission plugin secret fields changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Kubernetes.KubeAPIServer.AdmissionPlugins = []gardencorev1beta1.AdmissionPlugin{{KubeconfigSecretName: pointer.String("foo")}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})

		It("should return true because the gardener-apiserver admission plugin secret fields changed", func() {
			oldShoot := garden.DeepCopy()
			garden.Spec.VirtualCluster.Gardener.APIServer.AdmissionPlugins = []gardencorev1beta1.AdmissionPlugin{{KubeconfigSecretName: pointer.String("foo")}}
			Expect(Predicate(oldShoot, garden)).To(BeTrue())
		})
	})
})
