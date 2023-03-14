#!/bin/bash

# Prompt user for PostgreSQL username and password
read -p "Enter PostgreSQL username: " username
read -sp "Enter PostgreSQL password: " password

# Set environment variables in .env file
echo "POSTGRES_USER=$username" > .env
echo "POSTGRES_PASSWORD=$password" >> .env
echo "POSTGRES_DB=coffee-shop-db" >> .env

# Print success message
echo -e "\nEnvironment variables set successfully in .env file."
