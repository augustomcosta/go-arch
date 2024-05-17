#!/bin/sh

cd ..

while IFS='=' read -r key value; do
    echo "$key" | grep -q "^GOOSE_"
    if [ $? -eq 0 ]; then
        export "$key"="$value"
    fi
done < .env

MIGRATIONS_DIR="sql/migration"

goose -dir "$MIGRATIONS_DIR" "$@"