# [WIP] coffee-api-go

## Set up username and password for DB 
> This step is needed only for the first time you're setting up the database or in case you delete the `.env` file

> If you're doing it because you deleted the `.env` file, you will also need to put down the docker-compose before re-running the script.

It executes a bash script that prompts the user to enter their PostgreSQL username and password, 
and sets them as environment variables in a `.env` file.

- The script that creates a new `.env` file if it doesn't exist or overwrites the existing one if it does. 

To make the script executable, you need to give it execute permission using the chmod command:

    chmod +x set-env.sh

Then, you can then run the script using:

    bash ./set-env.sh

## stand up the application

    docker-compose up --build -d
    
it will build and stand up the DB and App docker images
