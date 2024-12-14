/*
Copyright (c) 2024 Diagrid Inc.
Licensed under the MIT License.
*/

package etcd

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/go-logr/logr"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"

	"github.com/diagridio/go-etcd-cron/internal/client"
)

func Embedded(t *testing.T) client.Interface {
	t.Helper()
	return client.New(client.Options{
		Log:    logr.Discard(),
		Client: EmbeddedBareClient(t),
	})
}

func EmbeddedBareClient(t *testing.T) *clientv3.Client {
	t.Helper()

	cfg := embed.NewConfig()
	cfg.LogLevel = "error"
	cfg.Dir = t.TempDir()
	peerPort, err := freeport.GetFreePort()
	clientPort, err := freeport.GetFreePort()

	lpurl, err := url.Parse("http://127.0.0.1:" + strconv.Itoa(peerPort))
	require.NoError(t, err)
	lcurl, err := url.Parse("http://127.0.0.1:" + strconv.Itoa(clientPort))
	require.NoError(t, err)

	require.NoError(t, err)
	cfg.ListenPeerUrls = []url.URL{*lpurl}
	cfg.ListenClientUrls = []url.URL{*lcurl}

	etcd, err := embed.StartEtcd(cfg)
	require.NoError(t, err)
	t.Cleanup(etcd.Close)

	cl, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcd.Clients[0].Addr().String()},
	})
	require.NoError(t, err)

	select {
	case <-etcd.Server.ReadyNotify():
	case <-time.After(2 * time.Second):
		assert.Fail(t, "server took too long to start")
	}

	return cl
}
