package clone

import (
	"github.com/openlab-red/etcd-data-migration/pkg/utils"
	"github.com/spf13/viper"
	"github.com/openlab-red/etcd-data-migration/pkg/etcd"
	log "github.com/sirupsen/logrus"
)

func Start() {

	flags := []string{"source", "target"}
	utils.Validate(flags)

	clone := Clone{
		Source:    viper.GetString("source"),
		Target:    viper.GetString("target"),
		Overwrite: viper.GetBool("overwrite"),
	}

	clone.start()
}

func (c *Clone) start() {

	source := etcd.Cli{}
	source.Connect([]string{c.Source})

	target := etcd.Cli{}
	target.Connect([]string{c.Target})

	resp := source.Get("")

	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		value := string(kv.Value)

		if c.Overwrite || target.Get(key).Count == 0 {
			target.Put(key, value, "")
		} else {
			log.WithFields(log.Fields{
				"key":       key,
				"overwrite": c.Overwrite,
			}).Infoln("Exiting")
		}

	}

}
