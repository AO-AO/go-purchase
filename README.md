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

    <server_address>:9401/api/v1/receipt/validate

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

## 获得product信息

### URL

    <server_address>:9401/api/v1/receipt/filter

### 参数说明

*'\*' 标识必选参数*

- validate_result * (json)

  validate接口返回详情。
 
  *PS.最佳做法是只传入 stataus/in_app 两个必要参数*

- product_list * (json)

  数据库中配置的proudcts列表，目前**能且只能**解析一下样例中的字段：

  ```json
  {
      "discount": 50,
      "best_deal": true,
      "effect": 100,
      "is_hot": true,
      "iap": {
          "price": "19.99",
          "product_id": "com.cocojulia1995.neon.vip0"
      },
      "kind": "vip",
      "offer_id": "getx.iap.vip.tier0",
      "way": "iap",
      "subscription": {
          "period_day": 30,
          "period_month": 1,
          "period_week": 4,
          "qualifying_period": 3
      }
  }
  ```

- transaction_id (string)

  可选参数，如果传入该参数，则尝试筛选出和传入 transaction\_id 一致的 in_app 购买项，匹配不到则返回[]。

- offer_id (string)

  可选参数，当筛选出有效的in_app购买项后；
  如果传入offer\_id，则通过offer_id筛选出匹配的dbProduct;
  如果没有传入offer\_id，则通过product_id筛选出匹配的dbProduct。

### 请求参数样例

- [filter接口请求样例](./examples/filter-request.json)

### 返回样例

- [filter接口返回样例](./examples/filter-response.json)