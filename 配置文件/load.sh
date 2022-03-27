#! /bin/sh

directory="docker-save"
cd ${directory}
images=$(ll | awk 'NR == 1 {next} {print $9}')
for image in ${images}; do
    echo "-----正在导入镜像${image}"
    docker load -i ${image}
done
cd ..
