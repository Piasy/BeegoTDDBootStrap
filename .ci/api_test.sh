#!/bin/bash
function doTest()
{
    pwd
    output=`go test -v -cover . | grep -E "coverage:|FAIL:"`
    pass=$?
    if (( $pass != 0 )); then
        return 1
    fi
    coverage=`echo $output | grep "coverage:"`
    echo $coverage
    fail=`echo $output | grep "FAIL:"`
    pass=$?
    if (( $pass == 0 )); then
        return 1
    fi
    return 0
}

.ci/prepare_db.sh && cd tests && doTest
passed=$?

echo "passed: $passed"
cd .. && .ci/clean_db.sh
exit ${passed}
