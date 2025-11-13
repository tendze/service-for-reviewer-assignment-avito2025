#!/bin/bash

set -e

MIGRATIONS_DIR="./migrations"

add_timestamp() {
    local file="$1"
    local direction="$2"
    local timestamp=$(date +%Y%m%d%H%M%S)
    local base_name=$(basename "$file" ".${direction}.sql")
    local new_name="${timestamp}_${base_name}.${direction}.sql"
    
    mv "$file" "$MIGRATIONS_DIR/$new_name"
}

# находим и переименовываем файлы без timestamp
find "$MIGRATIONS_DIR" -name "*.sql" | while read -r file; do
    filename=$(basename "$file")
    
    # скип миграций с норм. названием
    if [[ $filename =~ ^[0-9]{14}_.*\.(up|down)\.sql$ ]]; then
        continue
    fi
    
    if [[ $filename =~ \.up\.sql$ ]]; then
        base_name=$(basename "$filename" ".up.sql")
        down_file="$MIGRATIONS_DIR/${base_name}.down.sql"
        if [ -f "$down_file" ]; then
            timestamp=$(date +%Y%m%d%H%M%S)
            mv "$file" "$MIGRATIONS_DIR/${timestamp}_${base_name}.up.sql"
            mv "$down_file" "$MIGRATIONS_DIR/${timestamp}_${base_name}.down.sql"
        else
            add_timestamp "$file" "up"
        fi
    
    elif [[ $filename =~ \.down\.sql$ ]]; then
        base_name=$(basename "$filename" ".down.sql")
        up_file="$MIGRATIONS_DIR/${base_name}.up.sql"
        if [ ! -f "$up_file" ]; then
            add_timestamp "$file" "down"
        fi
    fi
done