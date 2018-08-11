#! /bin/bash

set ue

MODIFIED_FILE_SIZE=`git ls-files -m | wc -l | sed -e 's/[ \t]//g'`

print_error () {
  echo "$1 is $2"
  git status -s
  git --no-pager diff --color-words --word-diff-regex='\\w+|[^[:space:]]'
  exit 1
}

if [ $MODIFIED_FILE_SIZE -ne 0 ]; then
  print_error MODIFIED_FILE_SIZE $MODIFIED_FILE_SIZE
  exit 1
else
  exit 0
fi
