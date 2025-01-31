#!/bin/sh

## Constants
TRUE=1
FALSE=0
TODO="TODO"
UNDEF="UNDEF"

minitest_start() {
    success_count=0
    failure_count=0
}

minitest_stop() {
    echo "SUCCESS count: $success_count"
    echo "FAILURE count: $failure_count"
}

assert_eq() {
    expect="$1"; actual="$2"; msg="$3"
    expect_pretty=${expect//$'\n'/␞}; expect_pretty=${expect_pretty//$'\t'/␟};
    actual_pretty=${actual//$'\n'/␞}; actual_pretty=${actual_pretty//$'\t'/␟}
    if [ "$expect" == "$actual" ]; then
        success_count=$((success_count+1))
        echo "assert_eq $msg expect:$expect_pretty actual_pretty: SUCCESS"
    else
        failure_count=$((failure_count+1))
        echo "assert_eq $msg expect:$expect_pretty actual_pretty: FAILURE"
    fi
}

assert_eq_todo() {
    assert_eq $TODO "$2" "$3"
}
