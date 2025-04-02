#!/bin/bash
# Use this script to test if a given TCP host/port are available

set -x  # Enables debugging (logs all executed commands)

WAITFORIT_cmdname=${0##*/}

echoerr() { 
    if [[ $WAITFORIT_QUIET -ne 1 ]]; then 
        echo "[DEBUG] $@" 1>&2 
    fi 
}

usage() {
    cat << USAGE >&2
Usage:
    $WAITFORIT_cmdname host:port [-s] [-t timeout] [-- command args]
USAGE
    exit 1
}

wait_for() {
    if [[ $WAITFORIT_TIMEOUT -gt 0 ]]; then
        echoerr "Waiting $WAITFORIT_TIMEOUT seconds for $WAITFORIT_HOST:$WAITFORIT_PORT"
    else
        echoerr "Waiting indefinitely for $WAITFORIT_HOST:$WAITFORIT_PORT"
    fi

    WAITFORIT_start_ts=$(date +%s)

    while true; do
        if [[ $WAITFORIT_ISBUSY -eq 1 ]]; then
            nc -zv $WAITFORIT_HOST $WAITFORIT_PORT
            WAITFORIT_result=$?
        else
            (echo -n > /dev/tcp/$WAITFORIT_HOST/$WAITFORIT_PORT) >/dev/null 2>&1
            WAITFORIT_result=$?
        fi

        if [[ $WAITFORIT_result -eq 0 ]]; then
            WAITFORIT_end_ts=$(date +%s)
            echoerr "Connection successful! $WAITFORIT_HOST:$WAITFORIT_PORT available after $((WAITFORIT_end_ts - WAITFORIT_start_ts)) seconds"
            break
        fi

        echoerr "Still waiting for $WAITFORIT_HOST:$WAITFORIT_PORT..."
        sleep 1
    done

    return $WAITFORIT_result
}

# Process arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        *:*)
            IFS=":" read -r WAITFORIT_HOST WAITFORIT_PORT <<< "$1"
            shift
            ;;
        -q | --quiet)
            WAITFORIT_QUIET=1
            shift
            ;;
        -t | --timeout)
            WAITFORIT_TIMEOUT="$2"
            shift 2
            ;;
        --)
            shift
            WAITFORIT_CLI="$@"
            break
            ;;
        *)
            echoerr "Unknown argument: $1"
            usage
            ;;
    esac
done

if [[ -z "$WAITFORIT_HOST" || -z "$WAITFORIT_PORT" ]]; then
    echoerr "Error: Must provide a host and port"
    usage
fi

WAITFORIT_TIMEOUT=${WAITFORIT_TIMEOUT:-15}

echoerr "Connecting to $WAITFORIT_HOST:$WAITFORIT_PORT with timeout $WAITFORIT_TIMEOUT seconds..."

wait_for
WAITFORIT_RESULT=$?

if [[ -n "$WAITFORIT_CLI" ]]; then
    if [[ $WAITFORIT_RESULT -ne 0 ]]; then
        echoerr "Strict mode enabled. Command execution aborted."
        exit $WAITFORIT_RESULT
    fi
    exec "$WAITFORIT_CLI"
else
    exit $WAITFORIT_RESULT
fi