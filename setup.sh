#!/bin/bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/cosmtrek/air@latest

# import env vars to sqlboiler toml
./scripts/sqlboiler_setup.sh
