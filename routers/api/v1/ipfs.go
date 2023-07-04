package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/constant"
	"AiImageToNft/pkg/e"
	"AiImageToNft/pkg/setting"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageUpload struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	TaskId string `json:"task_id"`
}

type ImagePromptResult struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data ImageData `json:"data"`
}

type ImageData struct {
	Status       int      `json:"status"`
	Progress     int      `json:"progress"`
	Eta_relative int      `json:"eta_relative"`
	Imgs         []string `json:"imgs"`
}

type IpfsResult struct {
	Ok    bool        `json:"ok"`
	Value ValueResult `json:"value"`
}

type ValueResult struct {
	Cid string `json:"cid"`
}

func UploadImageToIpfs(path, fileName string) (string, error) {
	client := &http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	file, openErr := os.Open(path)
	if openErr != nil {
		return "", openErr
	}
	defer file.Close()

	fileWrite, _ := bodyWrite.CreateFormFile("file", fileName)
	_, copyErr := io.Copy(fileWrite, file)
	if copyErr != nil {
		return "", copyErr
	}
	bodyWrite.Close()
	contentType := bodyWrite.FormDataContentType()
	reqResult, err := http.NewRequest("POST", setting.AppSetting.IpfsNftStorageUrl, bodyBuf)
	if err != nil {
		return "", err
	}
	reqResult.Header.Set("Authorization", constant.Authorization)
	reqResult.Header.Set("Content-Type", contentType)
	reqResult.Header.Set("Accept", "application/json")

	resp, err := client.Do(reqResult)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// get image CID
	var data IpfsResult
	if err := json.Unmarshal(body, &data); err == nil {
		return data.Value.Cid, nil
	} else {
		return "", err
	}
}

func UploadToIpfs(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
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
	fileName := image.Filename
	cid, err := UploadImageToIpfs(dstFile.Name(), fileName)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UploadImageToIpfs, nil)
		return
	}

	ImagePathUrl := "https://" + string(cid) + ".ipfs.nftstorage.link/" + string(fileName)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url": ImagePathUrl,
	})
}
