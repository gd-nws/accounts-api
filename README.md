# Accounts API

API for managing Good News user accounts.

## Setup

### Setup Environment Variables

Setup `GOPATH` and `GOBIN` environment variables:
```shell script
export GOPATH=$(pwd)/src
export GOBIN=$GOPATH/bin
```

Set the environment variables for postgres and application secret:
```.env
DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=
SCHEMA=
JWT_KEY=
```

### Setup Postgres:
The database should have a users table like so:
```postgresql
create table users
(
	id serial not null
		constraint users_pk
			primary key,
	email varchar(255) not null,
	password varchar(255) not null,
	refresh_token varchar(255),
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP
);

create unique index users_email_uindex
	on users (email);
```

## Build
```shell script
cd ./src && go build -a -installsuffix cgo -o main .
```

## Run
```shell script
./src/main
```

## Deploy
Build and tag the docker image. Or run with:
```shell script
docker-compose up --build -d
```