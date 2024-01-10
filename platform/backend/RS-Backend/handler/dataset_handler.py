import requests
import json
import base64
import pdb
import cv2


def get_all_datasets_api(api_url):
    headers = {'Content-Type': 'application/json'}
    
    response = requests.get(api_url, headers=headers)
    return response.json()


def get_dataset_by_id_api(api_url,id):
    payload = json.dumps({"id": id})
    
    headers = {'Content-Type': 'application/json'}
    response = requests.get(api_url, data=payload, headers=headers)
    return response.json()


# 新建数据集 
# data={
#   "Name": "Test",
#   "Comment": "This is a test dataset",
#   "Path": "/test",
#   "StoreFormat": "CSV"
# }
def post_dataset(api_url,data):
    payload = json.dumps(data)

    headers = {'Content-Type': 'application/json'}
    response = requests.post(api_url, data=payload, headers=headers)
    
    return response.json()


# def upload_file(api_url,file_path,name,datasetId):
#     # Read the image as a file to be sent as form data
#         # Set up the data dictionary for additional form data parameters
#     with open(file_path, 'rb') as file:

#         data = {
#             'datasetId': datasetId,
#             'files':file,
#             'name': name
#         }
#     # files = {}
#     # with open(file_path, 'rb') as file:
#     #     files['files'] = (file_path, file, 'multipart/form-data')

#         # Send POST request with files and additional form data parameters
#     response = requests.post(api_url, data=data)

#         # Print the response content
#     print(response.json())  # Access JSON keys or other handling as needed

# api_url="http://118.178.134.87:8080/api/v1/uploadFile"
# file_paths='/root/workspace/src/RS-platform/platform/backend/RS-Backend/handler/test.csv'


# upload_file(api_url,file_paths,"Text",1)