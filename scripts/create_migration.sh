#!/bin/bash
migrate create -ext sql -dir db/migrations -seq $1
