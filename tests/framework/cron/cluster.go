/*
Copyright (c) 2024 Diagrid Inc.
Licensed under the MIT License.
*/

package cron

import (
	"testing"

	"github.com/diagridio/go-etcd-cron/tests/framework/etcd"
)

type Cluster struct {
	*Cron
	Crons [3]*Cron
}

func TripplePartition(t *testing.T) *Cluster {
	t.Helper()
	client := etcd.EmbeddedBareClient(t)
	cr1 := newCron(t, client, 3, 0)
	cr2 := newCron(t, client, 3, 1)
	cr3 := newCron(t, client, 3, 2)
	return &Cluster{
		Cron:  cr1,
		Crons: [3]*Cron{cr1, cr2, cr3},
	}
}

func TripplePartitionRun(t *testing.T) *Cluster {
	t.Helper()
	crs := TripplePartition(t)
	for _, cr := range crs.Crons {
		cr.run(t)
	}
	return crs
}

func (c *Cluster) Stop(t *testing.T) {
	t.Helper()

	for _, cr := range c.Crons {
		if cr != nil {
			cr.Stop(t)
		}
	}
}
