package agora

import (
	"fmt"
	"math"
)

const pageSize = 4
const pageSize8 = 8

// GetPageSize method
func GetPageSize(isDualCamera int8) int {
	if isDualCamera == 1 {
		return pageSize
	} else {
		return pageSize8
	}
}

// GetTotalPage method.
func GetTotalPage(total int, isDualCamera int8) int8 {
	if isDualCamera == 1 {
		return int8(math.Ceil(float64(total) / float64(pageSize)))
	} else {
		return int8(math.Ceil(float64(total) / float64(pageSize8)))
	}

}

// GetWrittenMeetingNo - 获取笔试考场id
func GetWrittenMeetingNo(env string, examroomID uint, uids []uint, uid uint, isDualCamera int8) string {
	// w
	pos := int(0)
	for _idx, _uid := range uids {
		if _uid == uid {
			pos = _idx + 1
			break
		}
	}

	pageN := GetTotalPage(pos, isDualCamera)
	return ToWrittenMeetingNo(env, examroomID, pageN)
}

// GetInviteMeetingNo - 获取面试候考区考场id
func GetInviteMeetingNo(env string, examroomID uint, uids []uint, uid uint, isDualCamera int8) string {
	// w
	pos := int(0)
	for _idx, _uid := range uids {
		if _uid == uid {
			pos = _idx + 1
			break
		}
	}

	pageN := GetTotalPage(pos, isDualCamera)
	return ToInviteMeetingNo(env, examroomID, pageN)
}

// ToWrittenMeetingNo - 拼接笔试考场id
func ToWrittenMeetingNo(env string, examroomID uint, pageN int8) string {
	return fmt.Sprintf("%s_w_%d_%d", env, examroomID, pageN)
}

// 拼接面试候考区考场id
func ToInviteMeetingNo(env string, examroomID uint, pageN int8) string {
	return fmt.Sprintf("%s_c_%d_%d", env, examroomID, pageN)
}

// GetMeetingNo - 获取考场id
func GetMeetingNo(env string, id uint) string {
	// m
	return fmt.Sprintf("%s_m_%d", env, id)
}

// GetCandidateMeetingNo - 获取侯考场id
func GetCandidateMeetingNo(env string, id uint) string {
	return fmt.Sprintf("%s_c_%d", env, id)
}

// GetExamineeIDStr - 获取考生idstr
func GetExamineeIDStr(uid uint) string {
	return fmt.Sprintf("10%d", uid)
}

// GetCandidateExamineeIDStr - 获取考生idstr
func GetCandidateExamineeIDStr(uid uint) string {
	return fmt.Sprintf("11%d", uid)
}

// GetServerIDStr - 获取服务端id，用于录制
func GetServerIDStr() string {
	return fmt.Sprintf("90")
}
