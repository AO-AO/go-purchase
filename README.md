# Purchase微服务说明

## 部署说说明

```shell
cd pincloud.purchase
version=1.0
docker build -t purchase:$version
docker run --env [SERVICE_ENV=] -p 8018:9401<dest:source> -itd
docker run --env SERVICE_ENV=production -p 8016:9401 -itd review:$version
```

## 校验 receipt

### URL

    \<server_address\>:9401/api/v1/receipt/validate

### 参数说明

*'\*' 标识必选参数*

- receipt * (string/json)

  receipt数据。苹果是编码的string，谷歌是json。

- market * (string)

  "ios" 或 "android"。表示校验商店。

- iap_config (json)

  没有传入正确的参数会校验失败

  - apple_password * (string)

    苹果必须配置

  - google_client_id * (string)

    谷歌必须配置

  - google_client_secret * (string)

    谷歌必须配置

  - google_client_secret * (string)

    谷歌必须配置

  - google_refresh_token * (string)

    谷歌必须配置

- sandbox_mode (bool)

  如果是苹果的沙盒模式，必须设为true

- product (string)

  日志用

- platform (string)

  日志用

- version (string)

  日志用

### 请求参数样例

- [apple 订阅购买](./examples/receipt-apple-subscription.json)

- [apple 内建购买](./examples/receipt-apple-build-in.json)

- [google 订阅购买](./examples/receipt-google-subscription.json)

- [google 内建购买](./examples/receipt-google-build-in.json)

### 返回样例

- [apple 订阅购买](./examples/response-apple-subscription.json)

- [apple 内建购买](./examples/response-apple-build-in.json)

- [google 订阅购买](./examples/response-google-subscription.json)

- [google 内建购买](./examples/response-google-build-in.json)