package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/e"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func GetImageData(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PromptToNft(c *gin.Context) {
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
	// ai 图片生成需要时间
	time.Sleep(10 * time.Second)

	imageUrl, err := GetNftUrl(imageTaskId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GetImageUrl, nil)
		return
	}

	imageData, err := GetImageData(imageUrl)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GetImageContent, nil)
		return
	}

	file := bytes.NewReader(imageData)
	dstFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CREATE_FILE, nil)
		return
	}
	defer os.Remove(dstFile.Name())
	defer dstFile.Close()
	_, err = io.Copy(dstFile, file)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CREATE_FILE, nil)
		return
	}

	fileName := prompt
	cid, err := UploadImageToIpfs(dstFile.Name(), fileName)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UploadImageToIpfs, nil)
		return
	}

	ImagePathUrl := "https://" + string(cid) + ".ipfs.nftstorage.link/" + fileName
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url": ImagePathUrl,
	})
}
