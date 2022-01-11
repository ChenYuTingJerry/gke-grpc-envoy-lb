#!/bin/sh

# Catch SIGTERM
_term() {
  # Send SIGTERM to child
  # And keep check child process is running
  kill -TERM "$child"
  stoped=0
  sec=0
  while [ $stoped -eq 0 ]
  do
    sec=$((sec + 1))
    if [ ! -e /proc/"$child" ]; then
      stoped=1
    else
      if [ $sec -eq 10 ]; then
        echo term timeout 10 sec
      fi
      sleep 1
    fi
  done
  # wait is used to capture the return code
  wait "$child"
  exit_status=$?
  echo graceful shut down with code $exit_status
  exit 0
}
trap _term TERM


cmd="${WORK_DIR}/main"
echo -----------------------------------------------------

date
$cmd &
child=$!
# first wait will be interrupted be a signal
wait $child
# second wait is used to capture the return code
wait $child
exit_status=$?
timestamp=$(date +"%s")
echo end of child process with code $exit_status
# send abnormal exit status to /logs for monitoring
exit $exit_status
