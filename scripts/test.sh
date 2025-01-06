#!/usr/bin/env bash

set -e

case "$1" in
  smoke)
    pytest -n auto .
    ;;
  cover)
    pytest --cov=. -n auto .
    ;;
  *)
    echo "Usage: ./test.sh [smoke|cover]" && exit 1
    ;;
esac
