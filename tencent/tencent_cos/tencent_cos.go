// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/2/5 2:27 下午
package tencent_cos

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type Config struct {
	BucketUrl string `json:"bucket_url" yaml:"bucket_url" mapstructure:"bucket_url"`
	SecretId  string `json:"secret_id" yaml:"secret_id" mapstructure:"secret_id"`
	SecretKey string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
}

type TencentCos struct {
	*cos.Client
}

func NewClient(cfg Config) *TencentCos {
	u, _ := url.Parse(cfg.BucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
		},
	})
	return &TencentCos{
		c,
	}

}

func (t *TencentCos) Get(prefix string) error {
	opt := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 1000,
	}
	res, _, err := t.Bucket.Get(context.Background(), opt)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// 写入生命周期 (临时)
func (t *TencentCos) InsertLifecycle() {

	lifecycle, _, err := t.Bucket.GetLifecycle(context.Background())
	fmt.Println(lifecycle, err)

	lc := &cos.BucketPutLifecycleOptions{
		Rules: lifecycle.Rules,
	}

	for _, env := range []string{"qa", "prod"} {
		for i := 1; i <= 12; i++ {
			lc.Rules = append(lc.Rules, cos.BucketLifecycleRule{
				ID:     fmt.Sprintf("examcloud/%s/deep%d", env, i),
				Filter: &cos.BucketLifecycleFilter{Prefix: fmt.Sprintf("examcloud/%s/deep%d", env, i)},
				Status: "Enabled",
				Transition: &cos.BucketLifecycleTransition{
					Days:         i * 30,
					StorageClass: "DEEP_ARCHIVE",
				},
			})
		}
	}

	putLifecycle, err := t.Bucket.PutLifecycle(context.Background(), lc)
	fmt.Println(putLifecycle, err)
}

// 删除生命周期 (临时)
func (t *TencentCos) DeleteLifecycle() {
	lifecycle, _, err := t.Bucket.GetLifecycle(context.Background())
	if err != nil {
		return
	}
	lc := &cos.BucketPutLifecycleOptions{}
	for _, v := range lifecycle.Rules {
		if match, err := regexp.Match(`examcloud/(qa|prod)/deep[0-9]{1,2}`, []byte(v.ID)); err != nil || !match {
			if err != nil {
				logrus.Error(err)
			}
			logrus.Info("not match:", v.ID)
			lc.Rules = append(lc.Rules, v)

			continue
		}
		logrus.Info("match:", v.ID)
	}

	if len(lc.Rules) == 0 {
		logrus.Fatal("no rules")
	}

	_, err = t.Bucket.PutLifecycle(context.Background(), lc)
	fmt.Println(err)

}
