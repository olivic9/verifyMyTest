#!/bin/bash -e

#setup_db() {
#  echo "Configuring the database"
#  ./verify-my-test migrate
#}
#
#setup_db
echo "Starting server"
./verify-my-test server
