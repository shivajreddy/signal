
# Introduction
- This is a startup project with the following techstack, that you can use to
  get started with least amount of friction.  
- Things like database, migrations are all taken care of.  
- If you want to change any tool in the stack, you can just replace that, 
  since all the tools are isolated cleanly.

# Pre-Reqs
- Docker, Docker-Desktop(for development)/Docker-Compose(for production)

# Tech Stack
Backend:
 - GO
    - godotenv: for reading .env files
    - GIN : GO framework
    - GORM - GO psql ORM

Database:
 - PSQL

Dev-Ops:
 - Git
 - Docker
    - Reflex (for hot-reload)


# Project Structure
helloworld
    |- .git
    |- .gitignore
    - docker-compose.yml
    - reflex.conf
    - server/
    - client/
    - database/


# Usage

- Step1: choose a project name
Current Project name: helloworld
Make sure to edit the following files to your new project name
    - server/go.mod
    - client/package.json
    - database/init.sql all the places you see 'helloworld'
    - database/backups/backup.sh
    - docker-compose.yml
    - docker-compose.prod.yml
    - NOTE: do searchgrep in the project for 'helloworld'

- Step2: Uncomment the environment files in .gitignore
        - These are commented so that you can get the files, and set the final
        environment values for your new project



