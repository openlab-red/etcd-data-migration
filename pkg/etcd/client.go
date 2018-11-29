package etcd

import (
	log "github.com/sirupsen/logrus"
	"time"
	"go.etcd.io/etcd/clientv3"
	"context"
	"crypto/tls"
	"strconv"
)

func (c *Cli) Connect(endpoints []string) {

	var err error
	var tlsConfig *tls.Config

	if c.Etcd == nil {
		if c.TLSInfo != nil {
			tlsConfig, err = c.TLSInfo.ClientConfig()
			if err != nil {
				log.Fatalln(err)
			}
		}

		c.Etcd, err = clientv3.New(clientv3.Config{
			Endpoints:          endpoints,
			DialTimeout:        5 * time.Second,
			TLS:                tlsConfig,
			MaxCallSendMsgSize: 20 * 1024 * 1024,
			MaxCallRecvMsgSize: 20 * 1024 * 1024,
		})
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Cli) Get(key string) (*clientv3.GetResponse) {
	cli := c.Etcd
	var resp *clientv3.GetResponse
	var err error

	if resp, err = cli.Get(context.Background(), key, clientv3.WithPrefix()); err != nil {
		log.Fatalln(err)
	}
	return resp

}

func (c *Cli) GetMapValue(key string) (map[string]interface{}) {
	data := make(map[string]interface{})
	if resp := c.Get(key); resp.Count > 0 {
		for _, kv := range resp.Kvs {
			key := string(kv.Key)
			value := string(kv.Value)
			log.Debugln(key + " | " + value)
			data[key] = value
		}
	}

	return data

}

func (c *Cli) Put(key string, value string, ttl string) {

	var err error
	cli := c.Etcd

	if ttl != "" {
		var sec int64
		sec, err = strconv.ParseInt(ttl, 10, 64)
		if err != nil {
			log.Println(err.Error())
		}
		var leaseResp *clientv3.LeaseGrantResponse
		leaseResp, err = cli.Grant(context.TODO(), sec)
		_, err = cli.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
	} else {
		_, err = cli.Put(context.Background(), key, value)
	}
	if err != nil {
		log.Fatalln(err)
	} else {
		if _, err := cli.Get(context.Background(), key, clientv3.WithPrefix()); err != nil {
			log.Fatalln(err)
		}
	}

}

func (c *Cli) ttl(lease int64) int64 {
	cli := c.Etcd
	resp, err := cli.Lease.TimeToLive(context.Background(), clientv3.LeaseID(lease))
	if err != nil {
		return 0
	}
	if resp.TTL == -1 {
		return 0
	}
	return resp.TTL
}
