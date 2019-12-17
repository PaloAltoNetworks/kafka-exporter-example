#!/bin/bash
# shellcheck disable=SC2068
launchsrv "kafka-exporter" ".env" ${@}
