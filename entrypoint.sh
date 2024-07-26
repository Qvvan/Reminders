#!/bin/sh

echo "Подключение к PostgreSQL..."
while ! nc -z $DB_HOST $DB_PORT; do
  sleep 0.2
done
echo "PostgreSQL подключен!"

exec /app/main
