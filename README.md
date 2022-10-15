# Setup:

 Set env vars (see below for dev settings)

`./setup.sh`  
`./scripts/migrate_up.sh`  
`./scripts/sqlboiler.sh`  
`docker-compose up`  
`air`  


# Dev Env Variables:
````
export DB_NAME="taxonomy"
export DB_HOST="localhost"
export DB_PORT="5433"
export DB_USER="docker"
export DB_PASSWORD="docker"
export DB_SSLMODE="disable"

export DB_PATH="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSLMODE"
````
If you are using a different location for postgres run ./scripts/sqlboiler_setup.sh to change sqlboiler config file after changing these variables
