#!/bin/sh
.ci/prepare_db.sh && \
cd tests && go test -v .

cd .. && .ci/clean_db.sh