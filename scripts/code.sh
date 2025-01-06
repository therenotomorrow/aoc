#!/usr/bin/env bash

set -e

isort .
black .

flake8 .
mypy .
