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

package gardenercontrollermanager

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"
	"k8s.io/utils/pointer"

	controllermanagerv1alpha1 "github.com/gardener/gardener/pkg/controllermanager/apis/config/v1alpha1"
	"github.com/gardener/gardener/pkg/logger"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	kubernetesutils "github.com/gardener/gardener/pkg/utils/kubernetes"
)

const (
	configMapControllerManagerPrefix = "gardener-controller-manager-config"
	dataConfigKey                    = "config.yaml"
)

var controllerManagerCodec runtime.Codec

func init() {
	controllerManagerScheme := runtime.NewScheme()
	utilruntime.Must(controllermanagerv1alpha1.AddToScheme(controllerManagerScheme))

	var (
		ser = json.NewSerializerWithOptions(json.DefaultMetaFactory, controllerManagerScheme, controllerManagerScheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: false,
			Strict: false,
		})
		versions = schema.GroupVersions([]schema.GroupVersion{
			controllermanagerv1alpha1.SchemeGroupVersion,
		})
	)

	controllerManagerCodec = serializer.NewCodecFactory(controllerManagerScheme).CodecForVersions(ser, ser, versions, versions)
}

func (g *gardenerControllerManager) configMapControllerManagerConfig() (*corev1.ConfigMap, error) {
	controllerManagerConfig := &controllermanagerv1alpha1.ControllerManagerConfiguration{
		GardenClientConnection: componentbaseconfigv1alpha1.ClientConnectionConfiguration{
			QPS:        100,
			Burst:      130,
			Kubeconfig: gardenerutils.PathGenericKubeconfig,
		},
		Controllers: controllermanagerv1alpha1.ControllerManagerControllerConfiguration{
			ControllerRegistration: &controllermanagerv1alpha1.ControllerRegistrationControllerConfiguration{
				ConcurrentSyncs: pointer.Int(20),
			},
			Project: &controllermanagerv1alpha1.ProjectControllerConfiguration{
				ConcurrentSyncs: pointer.Int(20),
				Quotas:          g.values.Quotas,
			},
			SecretBinding: &controllermanagerv1alpha1.SecretBindingControllerConfiguration{
				ConcurrentSyncs: pointer.Int(20),
			},
			Seed: &controllermanagerv1alpha1.SeedControllerConfiguration{
				ConcurrentSyncs:    pointer.Int(20),
				ShootMonitorPeriod: &metav1.Duration{Duration: 300 * time.Second},
			},
			SeedExtensionsCheck: &controllermanagerv1alpha1.SeedExtensionsCheckControllerConfiguration{
				ConditionThresholds: []controllermanagerv1alpha1.ConditionThreshold{{
					Duration: metav1.Duration{Duration: 1 * time.Minute},
					Type:     "ExtensionsReady",
				}},
			},
			SeedBackupBucketsCheck: &controllermanagerv1alpha1.SeedBackupBucketsCheckControllerConfiguration{
				ConditionThresholds: []controllermanagerv1alpha1.ConditionThreshold{{
					Duration: metav1.Duration{Duration: 1 * time.Minute},
					Type:     "BackupBucketsReady",
				}},
			},
			Event: &controllermanagerv1alpha1.EventControllerConfiguration{
				ConcurrentSyncs:   pointer.Int(10),
				TTLNonShootEvents: &metav1.Duration{Duration: 2 * time.Hour},
			},
			ShootMaintenance: controllermanagerv1alpha1.ShootMaintenanceControllerConfiguration{
				ConcurrentSyncs: pointer.Int(20),
			},
			ShootReference: &controllermanagerv1alpha1.ShootReferenceControllerConfiguration{
				ConcurrentSyncs: pointer.Int(20),
			},
		},
		LeaderElection: &componentbaseconfigv1alpha1.LeaderElectionConfiguration{
			LeaderElect:       pointer.Bool(true),
			ResourceName:      controllermanagerv1alpha1.ControllerManagerDefaultLockObjectName,
			ResourceNamespace: metav1.NamespaceSystem,
		},
		LogLevel:  g.values.LogLevel,
		LogFormat: logger.FormatJSON,
		Server: controllermanagerv1alpha1.ServerConfiguration{
			HealthProbes: &controllermanagerv1alpha1.Server{Port: probePort},
			Metrics:      &controllermanagerv1alpha1.Server{Port: metricsPort},
		},
		FeatureGates: g.values.FeatureGates,
	}

	data, err := runtime.Encode(controllerManagerCodec, controllerManagerConfig)
	if err != nil {
		return nil, err
	}

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapControllerManagerPrefix,
			Namespace: g.namespace,
			Labels:    GetLabels(),
		},
		Data: map[string]string{
			dataConfigKey: string(data),
		},
	}

	utilruntime.Must(kubernetesutils.MakeUnique(configMap))
	return configMap, nil
}
