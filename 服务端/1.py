import requests
res = requests.post("https://switchiot.dadiqq.cn/reflash").text
print(res)