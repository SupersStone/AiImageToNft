package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/e"
	"AiImageToNft/pkg/setting"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func GetNftUrl(taskId string) (string, error) {
	url := setting.AppSetting.PromptDownload + taskId

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

	// get image url
	var dataImage ImagePromptResult
	if err := json.Unmarshal(body, &dataImage); err == nil {
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

	imageUrl, err := GetNftUrl(imageTaskId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GetImageUrl, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url": imageUrl,
	})

}
