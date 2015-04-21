#!/bin/sh
git status
git add -A
read commit
git commit -m $commit
echo "commit:"
git status
sleep 3