# 智能开关
## 服务端
服务端采用GOLANG作为后台，采用gin框架，需要注意的是需要用户自行使用反向代理9090端口且在反向代理中做一层CORS，否则前端可能会出现CORS错误

## 舵机控制
舵机控制采用esp8266模块，使用Arduino进行开发，这一块比较简单就不详细概述了

## 网页端
网页端采用uniapp进行开发，打包后将静态文件放在服务器根目录即可。