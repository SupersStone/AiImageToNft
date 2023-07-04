package routers

import (
	v1 "AiImageToNft/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		//  prompt 生成图片
		apiv1.GET("/prompt", v1.PromptToNFT)
		// 根据task_id 获取图片链接
		apiv1.GET("/task_id", v1.GetImageAws3Url)
		// 上传图片到IPFS
		apiv1.POST("/nft", v1.UploadToIpfs)

		// 整合以上三个接口
		apiv1.GET("/prompttonft", v1.PromptToNft)
	}
	return r
}
