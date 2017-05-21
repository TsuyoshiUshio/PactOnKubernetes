#!/bin/bash
# You can execute this shell on provider directory
# Yo can share the /myapp/report directory for 
echo "The go provider test starts"
./bin/pact-go_linux_amd64 daemon &
./product.test -test.v -test.coverprofile report/cover.txt 2>&1 | bin/go-junit-report_linux > report/test-Result.xml
bin/gocov_linux convert report/cover.txt | bin/gocov-xml_linux > report/coverage.xml