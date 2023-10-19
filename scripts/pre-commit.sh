#!/usr/bin/env bash

set -x

./scripts/code.sh
./scripts/test.sh -fast

git add .

ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
