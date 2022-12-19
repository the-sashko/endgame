#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir" || exit 1

if [ ! -f ../bin/app ]; then
  echo "app binary not exists"
  exit 1;
fi

cd "$scriptDir/../bin" || exit 1

if [ ! -d ../logs ]; then
  mkdir ../logs
  chmod -R 775 ../logs
  exit 1;
fi

./app

cd "$currentDir" || exit 1

exit 0
