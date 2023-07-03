package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "请求参数错误",
	ERROR_AiPromptToImage:   "根据prompt生成图片失败",
	ERROR_GetImageUrl:       "获取图片路径失败",
	ERROR_GetImageContent:   "获取图片数据失败",
	ERROR_UploadImageToIpfs: "上传图片数据到IPFS失败",
	ERROR_NOT_CREATE_FILE:   "创建文件失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
