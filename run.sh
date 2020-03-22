#!/usr/bin/env bash

pid_file=run.pid

function start() {
  echo "go build -o $1 main.go"
  go build -o $1 main.go
  nohup "./$1" &
  if [[ $? -eq 0 ]]; then
      echo $! > ${pid_file}
  else exit 1
  fi
}


function stop() {
  if [ -f ${pid_file} ];then
    kill $(cat ${pid_file})
    echo "killed $(cat ${pid_file})"
    if [[ $? -eq 0 ]]; then
        rm -f ${pid_file}
    else exit 1
    fi
  else
  echo "cannot find ${pid_file}"
  fi

}

case $1 in
    'start')
        if [ $# != 2 ] ; then
          echo "USAGE: $0 program"
          echo " e.g.: $0 main"
          exit 1;
        fi
        start "$2"
        ;;
    'stop')
        stop
        ;;
    *)
        # shellcheck disable=SC2016
        echo 'Get invalid option, please input(as to $1):'
        echo -e '\t"start" -> start service'
        echo -e '\t"stop"  -> stop service'
  exit 1
esac