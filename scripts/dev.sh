#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/.." || exit 1

make dev

cd "$currentDir" || exit 1

exit 0
