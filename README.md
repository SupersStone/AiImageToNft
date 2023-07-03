# AiImageToNft
ai generate image output NFT



#### Api 
##### 1. prompt 生成图片  
- 请求方法：`GET`
- 请求地址：`/api/v1/prompt`
- 请求参数：`prompt` 生成图片的文字描述
  - 请求示例：
    - `127.0.0.1:8000/api/v1/prompt?prompt=dog`
- 返回参数：`task_id` 图片id

```json
 {
  "code": 200,
  "msg": "ok",
  "data": {
    "task_id": "429c4a5a-b094-4648-be5d-53b5f3d6b1f5"
  }
}
```


##### 2. 根据task_id 获取图片链接
- 请求方法：`GET`
- 请求地址：`/api/v1/task_id`
- 请求参数：`task_id` 生成图片的图片id 
  - 请求示例：
    - `127.0.0.1:8000/api/v1/task_id?task_id=429c4a5a-b094-4648-be5d-53b5f3d6b1f5`
- 返回参数：`image_url` 图片的aws3存储链接

```json
{
  "code": 200,
  "msg": "ok",
  "data": {
    "image_url": "https://stars-test.s3.amazonaws.com/free-prod/429c4a5a-b094-4648-be5d-53b5f3d6b1f5-0.png"
  }
}
```

##### 3. 上传图片到IPFS
- 请求方法：`POST`
- 请求地址：`/api/v1/nft`
- 请求参数：`image` 上传图片文件
- 返回参数：`image_url` 图片的IPFS链接地址

```json
{
  "code": 200,
  "msg": "ok",
  "data": {
    "image_url": "https://bafybeibvimit3vy5avilz2czc67xb36pfmdtkr3fx7pzaucrujhr62qifi.ipfs.nftstorage.link/111.png"
  }
}
```