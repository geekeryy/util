// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/26 2:21 下午
package elastic_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/elastic"
	"testing"
)

func init() {
	elastic.Init(config.GetConfig().Elastic)
}

func TestInit(t *testing.T) {

	doc := map[string]interface{}{
		"title":   "this is title",
		"content": "this is content",
	}

	if err := elastic.Index("demo", doc); err != nil {
		t.Error()
	}
}
