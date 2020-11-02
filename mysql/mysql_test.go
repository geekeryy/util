// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/30 3:32 下午
package mysql_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/mysql"
	"log"
	"testing"
)



func TestMysql(t *testing.T) {
	mysql.Init(config.LoadConfig("./../config.yaml").Mysql)


	db := mysql.Conn()
	defer mysql.Close()

	var data []int
	err := db.Table("demo").Pluck("id",&data).Error
	if err != nil {
		t.Error(err)
	}
	log.Println(data)
}
