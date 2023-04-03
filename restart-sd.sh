#!/usr/bin/env bash

 sudo fuser -k 7866/tcp
 source /home/ubuntu/miniconda3/etc/profile.d/conda.sh
 conda activate sdwebui
 cd /home/ubuntu/apps/stable-diffusion-webui
 ./webui.sh &