package agora

type AcquireModel struct {
	ResourceID string `json:"resourceId"`
}


type QueryResp struct {
	ResourceID     string `json:"resourceId"`
	Sid            string `json:"sid"`
	ServerResponse struct {
		FileListMode string `json:"fileListMode"`
		FileList     []struct {
			Filename       string `json:"filename"`
			TrackType      string `json:"trackType"`
			UID            string `json:"uid"`
			MixedAllUser   bool   `json:"mixedAllUser"`
			IsPlayable     bool   `json:"isPlayable"`
			SliceStartTime int64  `json:"sliceStartTime"`
		} `json:"fileList"`
		Status         string `json:"status"`
		SliceStartTime int64  `json:"sliceStartTime"`
	} `json:"serverResponse"`
}

type StartRecordModel struct {
	AcquireModel
	SID string `json:"sid"`
}

type ServerResponse struct {
	FileList        string `json:"fileList"`
	UploadingStatus string `json:"uploadingStatus"`
	Status          string `json:"status"`
	SliceStartTime  string `json:"sliceStartTime"`
}

type ResponseModel struct {
	StartRecordModel
	ServerResp ServerResponse `json:"serverResponse"`
}

type StatusModel struct {
	ResourceID     string `json:"resourceId"`
	SID            string `json:"sid"`
	ServerResponse struct {
		FileList        string `json:"fileList"`
		UploadingStatus string `json:"uploadingStatus"`
		Status          int    `json:"status"`
		SliceStartTime  int    `json:"sliceStartTime"`
	}
}

type KickingResponseModel struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

const (
	Recording = 1
	NotRecord = 2
)
