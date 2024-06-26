/*
Copyright (c) 2024 Diagrid Inc.
Licensed under the MIT License.
*/

package key

import (
	"path/filepath"
	"strconv"
)

type Options struct {
	// Namespace is the key namespace of all objects. Should be the same for all
	// cron partition replicas.
	Namespace string

	// PartitionID is the partition ID of the key.
	PartitionID uint32
}

// Key returns the correct namespaced key for the given object name.
type Key struct {
	// namespace is the key namespace of all objects.
	namespace string

	// partitionID is the partition ID of the key.
	partitionID string
}

// New returns a new Key with the given namespace.
func New(opts Options) *Key {
	return &Key{
		namespace:   opts.Namespace,
		partitionID: strconv.Itoa(int(opts.PartitionID)),
	}
}

// JobKey returns the job key for the given job name.
func (k *Key) JobKey(name string) string {
	return filepath.Join(k.namespace, "jobs", name)
}

// CounterKey returns the counter key for the given job name.
func (k *Key) CounterKey(name string) string {
	return filepath.Join(k.namespace, "counters", name)
}

// LeadershipNamespace returns the namespace key for the leadership keys.
func (k *Key) LeadershipNamespace() string {
	return filepath.Join(k.namespace, "leadership")
}

// LeadershipKey returns the leadership key for this partition ID.
func (k *Key) LeadershipKey() string {
	return filepath.Join(k.namespace, "leadership", k.partitionID)
}

// JobNamespace returns the job namespace key.
func (k *Key) JobNamespace() string {
	return filepath.Join(k.namespace, "jobs")
}

// JobName returns the job name from the given key.
func (k *Key) JobName(key []byte) string {
	return filepath.Base(string(key))
}
