# people

people is a full stack social media web app inspired by Twitter.

## About

Back end is a REST API built in API-First Approach using OpenAPI.
It uses JWT for authentication and WebSocket for real-time chat.

Front end is a React SPA with custom fully type-safe React Query hooks 
generated from OpenAPI specification.

The whole application is containerized and deployed on Linode using a simple
GitHub Actions CD.

[Demo Website](http://139.177.176.74)  
OpenAPI specification is available at `/swagger.html`

## Technologies used

* Go
* PostgreSQL
* TypeScript
* React
* React Query
* Mantine

## Dependencies

* make
* docker
* npm

## Installation

You can clone the repo with `git clone https://github.com/toxeeec/people`.

## Usage

1. Modify the `.env` file (you can omit this step if you want to use default values).
2. Run the `make` command to build and run the project.
3. You can stop the docker containers with Ctrl+C.
4. Run `make down` to remove stopped containers.

Docker compose will create the database, build and then run both backend and frontend.

## TODO

* Notification System
