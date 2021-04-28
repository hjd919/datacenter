### 用户服务，含登陆


### 创建api 文件
```shell 生成接口
goctl api -o user.api
```
### 生成user api服务
```
goctl api go -api user.api -dir .
```

### 进入微服务目录
```
goctl rpc proto -src user.proto -dir .
```