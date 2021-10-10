// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/1/5 4:47 下午
package agora_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/comeonjy/util/agora"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/redis"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	cname = "demo"
	uid   = "90"
)

const RSID = "util_rsid_%s_%s"
const SID = "util_sid_%s_%s"
const Record = "util_record_%s_%s"

func init() {
	redis.Init(config.LoadConfig().Redis)
	agora.Init()
}

func TestStartRecord(t *testing.T) {

	acquire:=agora.AcquireModel{}
	if err := redis.GetConn().Get(context.Background(), fmt.Sprintf(Record, cname, uid)).Err(); err != nil {
		logrus.Error(err)
	}


	record, err := agora.POSTMixStartRecord(cname, uid, acquire.ResourceID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(record, err)

	marshal, _ := json.Marshal(record)

	if err := redis.GetConn().Set(context.Background(), fmt.Sprintf(Record, cname, uid), string(marshal), time.Minute*24*60).Err(); err != nil {
		logrus.Fatal(err)
	}
}

func TestStopRecord(t *testing.T) {
	recordModel:=agora.StartRecordModel{}
	res, err := redis.GetConn().Get(context.Background(), fmt.Sprintf(Record, cname, uid)).Result()
	if err != nil {
		logrus.Fatal(err)
	}
	if err := json.Unmarshal([]byte(res), &recordModel);err!=nil{
		logrus.Fatal(err)
	}
	record, err := agora.POSTStopRecord(2, recordModel.SID, recordModel.ResourceID, cname, uid)
	fmt.Println(record, err)
}

func TestQueryRecord(t *testing.T) {
	recordModel:=agora.StartRecordModel{}
	res, err := redis.GetConn().Get(context.Background(), fmt.Sprintf(Record, cname, uid)).Result()
	if err != nil {
		logrus.Fatal(err)
	}
	if err := json.Unmarshal([]byte(res), &recordModel);err!=nil{
		logrus.Fatal(err)
	}

	record, err := agora.GETQueryRecord(2, recordModel.SID, recordModel.ResourceID)
	fmt.Println(record, err)
}
