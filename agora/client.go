package agora

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// record mode
const (
	MixRecord         = 2
	IndividualReocord = 3
)

// constant
var (
	baseURL        = ""
	kickURL        = ""
	appID          = ""
	appCertificate = ""
	customerID     = ""
	certificate    = ""
	cosAccessKey   = ""
	cosSecretKey   = ""
	cosBucket      = ""
)

// 初始化配置
func Init() {
	baseURL = viper.GetString("agora.base_url")
	kickURL = viper.GetString("agora.kick_url")
	appID = viper.GetString("agora.appid")
	appCertificate = viper.GetString("agora.app_certificate")
	customerID = viper.GetString("agora.customer_id")
	certificate = viper.GetString("agora.certificate")
	cosAccessKey = viper.GetString("agora.cos_accesskey")
	cosSecretKey = viper.GetString("agora.cos_secretkey")
	cosBucket = viper.GetString("agora.cos_bucket")
}

func genAuth() string {
	key := fmt.Sprintf("%s:%s", customerID, certificate)
	base64Str := base64.StdEncoding.EncodeToString([]byte(key))
	authKey := fmt.Sprintf("Basic %s", base64Str)
	return authKey
}

// POSTAcquire return acquire
func POSTAcquire(cname string, uid string) (*AcquireModel, error) {
	url := fmt.Sprintf("%s/%s/cloud_recording/acquire", baseURL, appID)
	authHeader := req.Header{"Authorization": genAuth()}
	params := req.Param{
		"cname":         cname,
		"uid":           uid,
		"clientRequest": req.Param{},
	}
	req.Debug = true
	result, err := req.Post(url, authHeader, req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	var acquireModel AcquireModel
	err = result.ToJSON(&acquireModel)
	if err != nil {
		return nil, err
	}

	return &acquireModel, nil
}

// POSTMixStartRecord - 开启录制
func POSTMixStartRecord(cname string, uid string, resourceID string) (*StartRecordModel, error) {
	//uids:=[]string{"2718695915","1536386642"}
	if resourceID == "" {
		acquire, err := POSTAcquire(cname, uid)
		if err != nil {
			logrus.Infof("[startRecord] - acquire error: %s \n", err.Error())
			return nil, err
		}

		logrus.Infof("[startRecord] - acquire: %+v \n", acquire)
		if acquire == nil || acquire.ResourceID == "" {
			return nil, errors.New("not found")
		}

		resourceID = acquire.ResourceID
	}

	token, _, err := GetToken(cname, uid)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/cloud_recording/resourceid/%s/mode/mix/start", baseURL, appID, resourceID)
	authHeader := req.Header{"Authorization": genAuth()}
	params := req.Param{
		"cname": cname,
		"uid":   uid,
		"clientRequest": req.Param{
			"token": token,
			"recordingConfig": req.Param{
				"maxIdleTime": 480,
				"transcodingConfig": req.Param{
					"width":            424,
					"height":           240,
					"fps":              15,
					"bitrate":          300, // 码率
					"mixedVideoLayout": 1,
				},
				//"subscribeVideoUids": uids,
				//"subscribeAudioUids": uids,
			},
			"storageConfig": req.Param{
				"vendor":    0,
				"region":    2,
				"accessKey": cosAccessKey,
				"secretKey": cosSecretKey,
				"bucket":    cosBucket,
			},
		},
	}

	req.Debug = true
	result, err := req.Post(url, authHeader, req.BodyJSON(params))
	if err != nil {
		logrus.Infof("[startRecord] - request error: %s \n", err.Error())
		return nil, err
	}

	var startRecordModel StartRecordModel
	err = result.ToJSON(&startRecordModel)
	if err != nil {
		logrus.Infof("[startRecord] - tojson error: %s \n", err.Error())
		return nil, err
	}

	return &startRecordModel, nil
}

// POSTMixStartRecord - 开启录制
func POSTMixStartRecord2(cname string, uid string, resourceID string, userID uint) (*StartRecordModel, error) {
	uids := []string{
		GetExamineeIDStr(userID), GetCandidateExamineeIDStr(userID),
	}
	uid=strconv.Itoa(int(userID))
	if resourceID == "" {
		acquire, err := POSTAcquire(cname, uid)
		if err != nil {
			logrus.Infof("[startRecord] - acquire error: %s \n", err.Error())
			return nil, err
		}

		logrus.Infof("[startRecord] - acquire: %+v \n", acquire)
		if acquire == nil || acquire.ResourceID == "" {
			return nil, errors.New("not found")
		}

		resourceID = acquire.ResourceID
	}

	token, _, err := GetToken(cname, uid)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/cloud_recording/resourceid/%s/mode/mix/start", baseURL, appID, resourceID)
	authHeader := req.Header{"Authorization": genAuth()}
	params := req.Param{
		"cname": cname,
		"uid":   uid,
		"clientRequest": req.Param{
			"token": token,
			"recordingConfig": req.Param{
				"maxIdleTime": 480,
				"transcodingConfig": req.Param{
					"width":            424,
					"height":           240,
					"fps":              15,
					"bitrate":          300, // 码率
					"mixedVideoLayout": 1,
				},
				"subscribeVideoUids": uids,
				"subscribeAudioUids": uids,
			},

			"storageConfig": req.Param{
				"vendor":    3,
				"region":    1,
				"accessKey": cosAccessKey,
				"secretKey": cosSecretKey,
				"bucket":    cosBucket,
			},
		},
	}

	req.Debug = true
	result, err := req.Post(url, authHeader, req.BodyJSON(params))
	if err != nil {
		logrus.Infof("[startRecord] - request error: %s \n", err.Error())
		return nil, err
	}

	var startRecordModel StartRecordModel
	err = result.ToJSON(&startRecordModel)
	if err != nil {
		logrus.Infof("[startRecord] - tojson error: %s \n", err.Error())
		return nil, err
	}

	return &startRecordModel, nil
}

// POSTIndividualStartRecord - 开启录制
func POSTIndividualStartRecord(cname string, uid string, resourceID string) (*StartRecordModel, error) {

	if resourceID == "" {
		acquire, err := POSTAcquire(cname, uid)
		if err != nil {
			logrus.Infof("[startRecord] - acquire error: %s \n", err.Error())
			return nil, err
		}

		logrus.Infof("[startRecord] - acquire: %+v \n", acquire)
		if acquire == nil || acquire.ResourceID == "" {
			return nil, errors.New("not found")
		}

		resourceID = acquire.ResourceID
	}

	token, _, err := GetToken(cname, uid)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/cloud_recording/resourceid/%s/mode/individual/start", baseURL, appID, resourceID)
	authHeader := req.Header{"Authorization": genAuth()}
	params := req.Param{
		"cname": cname,
		"uid":   uid,
		"clientRequest": req.Param{
			"token": token,
			"recordingConfig": req.Param{
				"maxIdleTime":       480,
				"subscribeUidGroup": 2,
			},
			"storageConfig": req.Param{
				"vendor":    3,
				"region":    1,
				"accessKey": cosAccessKey,
				"secretKey": cosSecretKey,
				"bucket":    cosBucket,

				// "vendor":    2,
				// "region":    3,
				// "accessKey": CosAccessKey,
				// "secretKey": CosSecretKey,
				// "bucket":    CosBucket,
			},
		},
	}

	req.Debug = true
	result, err := req.Post(url, authHeader, req.BodyJSON(params))
	if err != nil {
		logrus.Infof("[startRecord] - request error: %s \n", err.Error())
		return nil, err
	}

	var startRecordModel StartRecordModel
	err = result.ToJSON(&startRecordModel)
	if err != nil {
		logrus.Infof("[startRecord] - tojson error: %s \n", err.Error())
		return nil, err
	}

	return &startRecordModel, nil
}

// POSTStopRecord - 停止录制
func POSTStopRecord(mode int8, sid string, rsid string, cname string, uid string) (*ResponseModel, error) {
	modeStr := "mix"
	if mode == IndividualReocord {
		modeStr = "individual"
	}
	url := fmt.Sprintf("%s/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/stop", baseURL, appID, rsid, sid, modeStr)
	logrus.Infof("[stopRecord] - request url: %s \n", url)
	authHeader := req.Header{"Authorization": genAuth()}
	params := req.Param{
		"cname":         cname,
		"uid":           uid,
		"clientRequest": req.Param{},
	}

	req.Debug = true
	result, err := req.Post(url, authHeader, req.BodyJSON(params))
	if err != nil {
		logrus.Infof("[stopRecord] - request error: %s \n", err.Error())
		return nil, err
	}

	var stopRecordModel ResponseModel
	err = result.ToJSON(&stopRecordModel)
	if err != nil {
		logrus.Infof("[stopRecord] - tojson error: %s \n", err)
		return nil, err
	}

	return &stopRecordModel, nil
}

// GETQueryRecord - 查询录制状态
func GETQueryRecord(mode int8, sid string, rsid string) (int, error) {
	modeStr := "mix"
	if mode == IndividualReocord {
		modeStr = "individual"
	}

	url := fmt.Sprintf("%s/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/query", baseURL, appID, rsid, sid, modeStr)
	logrus.Infof("[queryRecord] - request url: %s \n", url)
	authHeader := req.Header{"Authorization": genAuth()}

	req.Debug = true
	result, err := req.Get(url, authHeader)
	if err != nil {
		logrus.Infof("[queryRecord] - request error: %s \n", err.Error())
		return 0, err
	}

	if result.Response().StatusCode == http.StatusNotFound {
		return NotRecord, nil
	} else if result.Response().StatusCode == http.StatusOK {
		return Recording, nil
	} else {
		return 0, errors.New(fmt.Sprintf("unknow err: %d", result.Response().StatusCode))
	}

}

// KickingRule - 踢人
func KickingRule(cname string, uid string) (*KickingResponseModel, error) {
	url := fmt.Sprintf("%s/kicking-rule", kickURL)
	authHeader := req.Header{"Authorization": genAuth(), "Content-Type": "application/json;charset=utf-8"}

	params := req.Param{
		"appid": appID,
		"cname": cname,
		"time":  0,
	}
	if uid != "" {
		params["uid"] = uid
	}
	result, err := req.Post(url, authHeader, req.BodyJSON(&params))
	if err != nil {
		return nil, err
	}

	var responseModel KickingResponseModel
	err = result.ToJSON(&responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

// GetToken - 获取token
func GetToken(cname string, uidStr string) (string, string, error) {
	rtcToken, err := BuildTokenWithUserAccount(appID, appCertificate, cname, uidStr, RoleAttendee, 0)
	if err != nil {
		return "", "", err
	}

	rtmToken, err := BuildToken(appID, appCertificate, uidStr, RoleRtmUser, 0)
	if err != nil {
		return "", "", err
	}

	return rtcToken, rtmToken, nil
}

// GetRecordURL - 获取录制url
func GetRecordURL(mode int8, sid string, cname string, uid uint) (string, string) {
	if mode == IndividualReocord {
		uidFrontStr := fmt.Sprintf("10%d", uid)
		uidTailStr := fmt.Sprintf("11%d", uid)
		return fmt.Sprintf("%s_%s__uid_s_%s__uid_e_video.m3u8", sid, cname, uidFrontStr),
			fmt.Sprintf("%s_%s__uid_s_%s__uid_e_video.m3u8", sid, cname, uidTailStr)
	}
	return fmt.Sprintf("%s_%s.m3u8", sid, cname), ""
}
