# Readme

## 更新日志
### v1.3 （2022年5月13日）
1. 新增get方法接口控制，方便嵌入式设备直接调用
2. 优化代码结构
### v1.2 （2021年12月22日）
1. 新增阿里云短信
2. 使用redis进行缓存
3. 众多bug修复
### v1.1（2021年12月19日）
1. 对直接部署情况下的CORS通过注释给出，如果需要直接绕过CORS请使用。
2. 可使用mysql连接到数据库，提供用户注册功能。
3. 衍生自制简单mysql插件，sqld，详细参考[sqld github地址](https://github.com/PonyWilliam/sqld)

### v1.0
1. 服务端使用gin开发，简单且并发量较高
2. 使用rest API 标准化接口



## 关于本程序

嵌入式平台通过mqtt协议连接到阿里云物联网设备，服务端使用golang作为中间件负责客户端与阿里云服务器的双向通信，同时提供部分特别功能，如不通过服务端订阅直接监控设备在线状态（阿里云对服务端订阅要求过于严苛···）  

## 注意
部署到服务端后，为了防止CORS，请在前端使用JSONP策略或使用nginx反向代理解决。  