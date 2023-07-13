# AiImageToNft
ai generate image output NFT

Link:
- Prompt Generate Images : https://omniinfer.io/
- Image Upload IPFS : https://nft.storage/



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


##### 4. prompt 生成图片并上传到IPFS（接口整合）
- 请求方法：`GET`
- 请求地址：`/api/v1/prompttonft`
- 请求参数：`prompt` 生成图片的文字描述
  - 请求示例：
    - `127.0.0.1:8000/api/v1/prompttonft?prompt=dog`
- 返回参数：`image_url` 图片的IPFS链接地址

```json
{
  "code": 200,
  "msg": "ok",
  "data": {
    "image_url": "https://bafybeicbivtgneo3vov3gjmod7eix3ioyqgpquphicqcwtnug53wr7juey.ipfs.nftstorage.link/dog"
  }
}
```

##### 5. 根据图片链接上传到IPFS（新增）
- 请求方法：`GET`
- 请求地址：`/api/v1/nft`
- 请求参数：`image_url` 图片链接地址
  - 请求示例：
    - `127.0.0.1:8000/api/v1/nft?image_url=https://stars-test.s3.amazonaws.com/free-prod/c8683e80-780e-43dd-aa58-b824c576b705-0.png`
- 返回参数：`image_url` 图片的IPFS链接地址

```json
{
  "code": 200,
  "msg": "ok",
  "data": {
    "image_url": "https://bafybeifrwu3aiica2dzyekhazvehfsn5mz5hsajtxk5asujsnm5dfnpz2i.ipfs.nftstorage.link/c8683e80-780e-43dd-aa58-b824c576b705-0.png"
  }
}
```