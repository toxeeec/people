# people

people is a full stack social media web app.

## Built with

* Go
* PostgreSQL
* OpenAPI
* TypeScript
* React

## Dependencies

* make
* docker
* npm

## Installation

You can clone the repo with `git clone https://github.com/toxeeec/people`.

## Usage

1. Modify the `.env` file (you can omit this step if you want to use default values).
2. Type `make` to build and run the project.
3. You can stop the docker containers with Ctrl+C.
4. To remove stopped containers type `make down`.

Docker compose will create the database, build and then run both backend and frontend.

## TODO

* Editing posts
* Profile pictures
* Editing profile 
* Real-time chat
* Log out from all devices
* Deleting account
* Notifications
