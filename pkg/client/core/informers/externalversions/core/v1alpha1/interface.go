// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/core/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackupBuckets returns a BackupBucketInformer.
	BackupBuckets() BackupBucketInformer
	// BackupEntries returns a BackupEntryInformer.
	BackupEntries() BackupEntryInformer
	// ControllerInstallations returns a ControllerInstallationInformer.
	ControllerInstallations() ControllerInstallationInformer
	// ControllerRegistrations returns a ControllerRegistrationInformer.
	ControllerRegistrations() ControllerRegistrationInformer
	// Plants returns a PlantInformer.
	Plants() PlantInformer
	// Seeds returns a SeedInformer.
	Seeds() SeedInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackupBuckets returns a BackupBucketInformer.
func (v *version) BackupBuckets() BackupBucketInformer {
	return &backupBucketInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// BackupEntries returns a BackupEntryInformer.
func (v *version) BackupEntries() BackupEntryInformer {
	return &backupEntryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ControllerInstallations returns a ControllerInstallationInformer.
func (v *version) ControllerInstallations() ControllerInstallationInformer {
	return &controllerInstallationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ControllerRegistrations returns a ControllerRegistrationInformer.
func (v *version) ControllerRegistrations() ControllerRegistrationInformer {
	return &controllerRegistrationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Plants returns a PlantInformer.
func (v *version) Plants() PlantInformer {
	return &plantInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Seeds returns a SeedInformer.
func (v *version) Seeds() SeedInformer {
	return &seedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
