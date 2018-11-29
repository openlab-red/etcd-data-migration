package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
)

type Cli struct {
	Etcd *clientv3.Client
	TLSInfo  *transport.TLSInfo
}
