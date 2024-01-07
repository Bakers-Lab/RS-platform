import requests
import json
import base64
import cv2


def image_to_base64(image_path):
    # Read the image using OpenCV
    img = cv2.imread(image_path)
    # Convert the image to a binary stream
    _, img_encoded = cv2.imencode('.jpg', img)
    # Convert the binary stream to a base64 string
    return base64.b64encode(img_encoded).decode()


def test_api_with_image(image_path, url):
    # Convert image to base64 string
    base64_img_string = image_to_base64(image_path)

    # JSON payload for the POST request
    payload = json.dumps({"img": base64_img_string})

    # Set headers to specify JSON content type
    headers = {'Content-Type': 'application/json'}

    # Send POST request and get the response
    response = requests.post(url, data=payload, headers=headers)

    # Print the response JSON
    # print(response.json())
    print(response.json().keys())


# URL of your Flask API
api_url = "http://127.0.0.1:5000/predict"

# Path to your image file
image_file_path = "./demo/0.png"  # Replace with your image file path

# Test the API with the image
test_api_with_image(image_file_path, api_url)
