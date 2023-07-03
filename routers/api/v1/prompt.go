package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/e"
	"AiImageToNft/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// 根据prompt生成图片
func PromptToImage(prompt string) (string, error) {
	payload := fmt.Sprintf("{\"prompt\": \"%s\",\"negative_prompt\":\"watermark, facial distortion, lip deformity, redundant background, extra fingers, Abnormal eyesight, ((multiple faces)), ((Tongue protruding)), ((extra arm)), extra hands, extra fingers, deformity, missing legs, missing toes, missin hand, missin fingers, (painting by bad-artist-anime:0.9), (painting by bad-artist:0.9), watermark, text, error, blurry, jpeg artifacts, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, artist name, (worst quality, low quality:1.4), bad anatomy\",\"sampler_name\":\"Euler a\",\"batch_size\":1,\"n_iter\":1,\"steps\":20,\"cfg_scale\":7,\"seed\":-1,\"height\":512,\"width\":512,\"model_name\":\"AnythingV5_v5PrtRE.safetensors\"}", prompt)
	newPayload := strings.NewReader(payload)
	// http client
	client := &http.Client{}
	req, err := http.NewRequest("POST", setting.AppSetting.PromptUrl, newPayload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// image id
	var data Result
	if err := json.Unmarshal(body, &data); err == nil {
		fmt.Println(data.Data.TaskId)
		return data.Data.TaskId, nil
	} else {
		return "", err
	}
}

func PromptToNFT(c *gin.Context) {
	appG := app.Gin{C: c}
	prompt := c.Query("prompt")
	if prompt == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	imageTaskId, err := PromptToImage(prompt)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AiPromptToImage, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"task_id": imageTaskId,
	})

}
