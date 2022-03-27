#! /bin/sh

images=$(docker images | awk 'NR == 1 {next} {print $1}' | awk '!a[$0]++')
directory="docker-save"
rm -rf ${directory}
mkdir ${directory}
cd ${directory}
for image in ${images}; do
    echo "-----正在导出镜像${image}"
    image_name=$(echo ${image} | sed 's/\//_/')
    docker save -o ${image_name}.tar ${image}
done
cd ../
