package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/e"
	"AiImageToNft/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// 根据imageTaskId 获取图片链接
func GetNftUrl(taskId string) (string, error) {
	url := setting.AppSetting.PromptDownload + taskId
	fmt.Println("url is :", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))

	// get image url
	var dataImage ImagePromptResult
	if err := json.Unmarshal(body, &dataImage); err == nil {
		fmt.Println("dataImage is：", dataImage)
		return dataImage.Data.Imgs[0], nil
	} else {
		return "", err
	}
}

func GetImageAws3Url(c *gin.Context) {
	appG := app.Gin{C: c}
	imageTaskId := c.Query("task_id")
	if imageTaskId == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 2. 获取图片路径
	imageUrl, err := GetNftUrl(imageTaskId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GetImageUrl, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url": imageUrl,
	})

}
