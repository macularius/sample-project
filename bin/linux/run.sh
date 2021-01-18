#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/sample-project" -importPath sample-project -srcPath "$SCRIPTPATH/src" -runMode dev
