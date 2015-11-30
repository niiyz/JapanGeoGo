#!/bin/bash

find ./geojson -type f | while read file
do
    arr=( `echo $file | tr -s '/' ' '`)
    pref=${arr[2]}
	# ディレクトリ
    dir="topojson/"$pref
    if [ ! -d $dir ]; then
      mkdir -p $dir
    fi
    basename=$(basename $file)
    if [ $basename = ".DS_Store" ]; then
      continue
    fi
    name=( `echo $basename | sed -e "s/\.json/.topojson/"`)
    echo $name
	# --simplify
    topojson --id-property N03_004 -p NAME=name -p name -o $dir"/"$name geojson/$pref/$basename
done