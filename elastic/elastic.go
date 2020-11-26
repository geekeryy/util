// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/26 2:14 下午
package elastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var client *elasticsearch.Client

type Config struct {
	Address string `json:"address" yaml:"address"`
}

func Init(cfg Config) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.Address},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatal("elasticsearch info res is error")
	}
	client = es
	log.Println(res)
}

func Conn() *elasticsearch.Client {
	return client
}

func Index(index string,doc map[string]interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		return err
	}
	res, err := client.Index(index, &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return errors.New("index res err")
	}
	return nil
}