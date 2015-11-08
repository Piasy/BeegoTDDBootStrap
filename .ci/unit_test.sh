#!/bin/sh
.ci/prepare_db.sh && \
cd models && go test -v . && cd .. && \
cd utils && go test -v .

cd .. && .ci/clean_db.sh