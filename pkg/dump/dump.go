package dump

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"github.com/openlab-red/etcd-data-migration/pkg/etcd"
	"encoding/json"
	"io/ioutil"
	"github.com/openlab-red/etcd-data-migration/pkg/utils"
)

func Start() {

	flags := []string{"endpoint", "output"}
	utils.Validate(flags)

	endpoints := []string{viper.GetString("endpoint")}
	d := Dump{
		Endpoints: endpoints,
		Key:       "",
		Output:    viper.GetString("output"),
	}

	d.dump()
}

func (d *Dump) dump() {
	var bytes []byte
	var err error

	cli := etcd.Cli{}

	cli.Connect(d.Endpoints)
	data := cli.GetMapValue(d.Key)

	if bytes, err = json.Marshal(data); err != nil {
		log.Fatalln(err)
	} else {
		ioutil.WriteFile(d.Output, bytes, 0644)
	}

}
