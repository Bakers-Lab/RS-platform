import requests
import json
import base64
import cv2


def get_all_evaljobs_api(api_url):
    headers = {'Content-Type': 'application/json'}
    
    response = requests.get(api_url, headers=headers)
    return response.json()

def get_evaljob_by_id_api(api_url,id):
    payload = json.dumps({"id": id})
    
    # Set headers to specify JSON content type
    headers = {'Content-Type': 'application/json'}

    # Send POST request and get the response
    response = requests.get(api_url, data=payload, headers=headers)
    return response.json()
    