import requests
import json
import base64
import pdb
import cv2

# 用户登录
# data={
#     "email": "user1@163.com",
#     "password": "123456"
# }
# url='http://118.178.134.87:8080/api/v1/login'
def login(data,api_url):
    payload = json.dumps(data)
    headers = {'Content-Type': 'application/json'}
    response = requests.post(api_url, data=payload, headers=headers)
    return response.json()


# 用户注册
# data={
#   "userName": "user1",
#   "password": "123456",
#   "email": "user1@163.com"
# }
# url='http://118.178.134.87:8080/api/v1/register'
def register(data,api_url):
    payload = json.dumps(data)
    headers = {'Content-Type': 'application/json'}
    response = requests.post(api_url, data=payload, headers=headers)
    return response.json()
