#!/bin/bash

# Command: nohup pipe-docker <pipe-path> &
# Description: 파이프를 통해 명령어를 입력하면, 호스트에서 그대로 실행하는 스크립트
#    - 다른 터미널에서 명령을 보낼 때 echo "<명령어>" > <pipe-path>
#    - 실행 종료: jobs -> kill %<job 번호>

PIPE=$1

if [[ -z "$PIPE" ]]; then
  mkdir -p pipe
  mkfifo pipe/docker
  PIPE='pipe/docker'
fi

while true; 
do 
    eval "$(cat $PIPE)"; 
done