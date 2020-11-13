// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/13 9:45 上午
package redis_test

import (
	"context"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/redis"
	"log"
	"testing"
)

func init() {
	config.LoadConfig()
}

func TestGetConn(t *testing.T) {
	redis.Init(config.GetConfig().Redis)

	ctx := context.Background()
	client := redis.GetConn()

	log.Println(client.Set(ctx, "demo", 1, 0).Err())
	log.Println(client.Get(ctx, "demo").Result())

	log.Println(client.MSet(ctx, "num", 2, "num2", 4).Result())
	log.Println(client.MGet(ctx, "num"))
	log.Println(client.MGet(ctx, "num2"))

	log.Println(client.HMSet(ctx, "hm", "1", 2, "3", 4))
	log.Println(client.HMGet(ctx, "hm", "3"))
	log.Println(client.HGetAll(ctx, "hm"))

	log.Println(client.HSet(ctx, "h", "1", 1))
	log.Println(client.HGetAll(ctx, "h"))


}
