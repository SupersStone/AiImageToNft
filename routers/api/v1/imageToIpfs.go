package v1

import (
	"AiImageToNft/pkg/app"
	"AiImageToNft/pkg/e"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

// 验证图片链接是否正确
func ValidImageUrl(url string) ([]byte, error) {
	// 根据图片链接获取图片内容
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	// 获取图片内容
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// 判断图片是否存在
	if res.StatusCode != http.StatusOK {
		return nil, nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 根据图片链接上传到IPFS
func ImageUrlToIpfs(c *gin.Context) {
	appG := app.Gin{C: c}
	image_url := c.Query("image_url")
	if image_url == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	// 获取图片数据
	imageData, err := ValidImageUrl(image_url)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GetImageContent, nil)
		return
	}
	// 图片数据上传到IPFS
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

	fileName := strings.Split(image_url, "/")[len(strings.Split(image_url, "/"))-1]
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
