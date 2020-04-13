#!/bin/sh
set -e

docker run --rm -v /Users/toreglia/dev/breadcrumbs/migration/sql:/flyway/sql flyway/flyway -user="toreglia" -password="anthony" -url='jdbc:postgresql://localhost/breadcrumbs' -table=flyway_schema_history migrate info