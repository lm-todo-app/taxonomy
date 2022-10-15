#!/bin/bash
migrate -path db/migrations -database $DB_PATH -verbose down
