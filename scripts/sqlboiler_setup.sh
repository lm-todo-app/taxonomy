#!/bin/bash
# import env vars to sqlboiler toml

# for field in DB_NAME DB_HOST DB_PORT DB_USER DB_PASSWORD DB_SSLMODE
# do
#   sed -i "s#{field}#$field#g" sqlboiler.toml
# done

cp sqlboiler_template.toml sqlboiler.toml

sed -i "s#{DB_NAME}#$DB_NAME#g" sqlboiler.toml
sed -i "s#{DB_HOST}#$DB_HOST#g" sqlboiler.toml
sed -i "s#{DB_PORT}#$DB_PORT#g" sqlboiler.toml
sed -i "s#{DB_USER}#$DB_USER#g" sqlboiler.toml
sed -i "s#{DB_PASSWORD}#$DB_PASSWORD#g" sqlboiler.toml
sed -i "s#{DB_SSLMODE}#$DB_SSLMODE#g" sqlboiler.toml
