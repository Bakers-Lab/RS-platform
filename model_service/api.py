from flask import Flask, request, jsonify
import os
import numpy as np
import cv2
from mmseg.apis import init_model, inference_model
import base64
from io import BytesIO

app = Flask(__name__)

# Model initialization
config_file = './MyConfigs/loveda_KNet_20231204.py'  # Training config
checkpoint_file = './work_dirs/loveda-KNet/best_mIoU_iter_11500.pth'  # Trained model
device = 'cuda:0'
model = init_model(config_file, checkpoint_file, device=device)

# Palette as per your existing code
palette = [['background', [255, 255, 255]], ['building', [255, 0, 0]], ['road', [255, 255, 0]], ['water', [0, 0, 255]],
           ['barren', [159, 129, 183]], ['forest', [0, 255, 0]], ['agricultural', [255, 195, 128]]]

palette_dict = {idx: each[1] for idx, each in enumerate(palette)}

opacity = 0.3  # Opacity for overlay


def process_single_img(img):
    result = inference_model(model, img)
    pred_mask = result.pred_sem_seg.data[0].cpu().numpy()

    pred_mask_bgr = np.zeros((pred_mask.shape[0], pred_mask.shape[1], 3))
    for idx in palette_dict.keys():
        pred_mask_bgr[np.where(pred_mask == idx)] = palette_dict[idx]
    pred_mask_bgr = pred_mask_bgr.astype('uint8')

    pred_viz = cv2.addWeighted(img, opacity, pred_mask_bgr, 1 - opacity, 0)
    return pred_mask_bgr, pred_viz


def base64_to_image(base64_string):
    img_data = base64.b64decode(base64_string)
    img = cv2.imdecode(np.frombuffer(img_data, np.uint8), cv2.IMREAD_COLOR)
    return img


def image_to_base64(image):
    _, buffer = cv2.imencode('.jpg', image)
    return base64.b64encode(buffer).decode()


@app.route('/predict', methods=['POST'])
def predict():
    content = request.json
    base64_img = content['img']

    img = base64_to_image(base64_img)
    pred_mask_bgr, pred_viz = process_single_img(img)

    base64_pred_mask_bgr = image_to_base64(pred_mask_bgr)
    base64_pred_viz = image_to_base64(pred_viz)

    return jsonify({'mask': base64_pred_mask_bgr, 'viz': base64_pred_viz})


if __name__ == '__main__':
    app.run(debug=True)
