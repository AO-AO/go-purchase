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

- URL

\<server_address\>:9401/api/v1/receipt/validate

- 参数说明
- 返回样例