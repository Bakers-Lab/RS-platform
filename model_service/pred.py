import os
import numpy as np
import cv2
from mmseg.apis import init_model, inference_model, show_result_pyplot
import mmcv
from tqdm import tqdm
import matplotlib.pyplot as plt
# %matplotlib inline

config_file = './MyConfigs/loveda_KNet_20231204.py'   #训练时的config
checkpoint_file = './work_dirs/loveda-KNet/best_mIoU_iter_11500.pth'  #训练得到的模型
device = 'cuda:0'
model = init_model(config_file, checkpoint_file, device=device)
palette=[['background',[255, 255, 255]], 
         ['building',[255, 0, 0]],
         ['road', [255, 255, 0]],
         ['water', [0, 0, 255]],
         ['barren',[159, 129, 183]],
         ['forest', [0, 255, 0]],
         ['agricultural', [255, 195, 128]]]

palette_dict = {}
for idx, each in enumerate(palette):
    palette_dict[idx] = each[1]

if not os.path.exists('outputs/testset-pred'):
    os.mkdir('outputs/testset-pred')

PATH_IMAGE = 'data/loveDA/img_dir/val'

opacity = 0.3 # 透明度，越大越接近原图
def process_single_img(img_path, save=False):
    img_bgr = cv2.imread(img_path)
    # print(img_bgr)
    # 语义分割预测
    result = inference_model(model, img_bgr)
    pred_mask = result.pred_sem_seg.data[0].cpu().numpy()
    # 将预测的整数ID，映射为对应类别的颜色
    pred_mask_bgr = np.zeros((pred_mask.shape[0], pred_mask.shape[1], 3))
    for idx in palette_dict.keys():
        pred_mask_bgr[np.where(pred_mask==idx)] = palette_dict[idx]
    pred_mask_bgr = pred_mask_bgr.astype('uint8')
    # 将语义分割预测图和原图叠加显示
    pred_viz = cv2.addWeighted(img_bgr, opacity, pred_mask_bgr, 1-opacity, 0)
    # 保存图像至 outputs/testset-pred 目录
    if save:
        save_path = os.path.join('outputs', 'testset-pred', 'pred-'+img_path.split('/')[-1])
        cv2.imwrite(save_path, pred_viz)

for each in tqdm(os.listdir(PATH_IMAGE)):
    process_single_img(os.path.join(PATH_IMAGE,each), save=True)