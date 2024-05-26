#!/usr/bin/env bash

# Short commit id
echo STABLE_GIT_COMMIT $(git describe --always)

# Full commit id
#echo "STABLE_GIT_COMMIT $(git rev-parse HEAD)"
