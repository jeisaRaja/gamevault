#!/bin/bash

# Default values for environment variables
DEFAULT_PORT=8080
DEFAULT_DB_HOST=db
DEFAULT_DB_PORT=5432
DEFAULT_DB_USER=postgres
DEFAULT_DB_PASSWORD=yourpassword
DEFAULT_DB_NAME=yourdbname

# Prompt the user to enter values or use defaults
read -p "Enter application port (default: $DEFAULT_PORT): " APP_PORT
APP_PORT=${APP_PORT:-$DEFAULT_PORT}

read -p "Enter database host (default: $DEFAULT_DB_HOST): " DB_HOST
DB_HOST=${DB_HOST:-$DEFAULT_DB_HOST}

read -p "Enter database port (default: $DEFAULT_DB_PORT): " DB_PORT
DB_PORT=${DB_PORT:-$DEFAULT_DB_PORT}

read -p "Enter database user (default: $DEFAULT_DB_USER): " DB_USER
DB_USER=${DB_USER:-$DEFAULT_DB_USER}

read -p "Enter database password (default: $DEFAULT_DB_PASSWORD): " DB_PASSWORD
DB_PASSWORD=${DB_PASSWORD:-$DEFAULT_DB_PASSWORD}

read -p "Enter database name (default: $DEFAULT_DB_NAME): " DB_NAME
DB_NAME=${DB_NAME:-$DEFAULT_DB_NAME}

# Export environment variables so Docker Compose can use them
export DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME

# Use envsubst to replace the environment variables in docker-compose.yml dynamically
cat <<EOF > .env
DB_HOST=$DB_HOST
DB_PORT=$DB_PORT
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_NAME=$DB_NAME
APP_PORT=$APP_PORT
EOF

# Run Docker Compose with specified ports and environment variables
sudo docker compose up --build -d

# Wait for services to start up and provide feedback
echo "Starting Gamevault application on port $APP_PORT and connecting to PostgreSQL database..."
echo "Database Host: $DB_HOST"
echo "Database Port: $DB_PORT"
echo "Database User: $DB_USER"
echo "Database Name: $DB_NAME"
