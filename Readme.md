# Deposit Service

This service is used to deposit funds into the wallet

## API Docs

Import file `stockbit test.postman_collection.json` at directory `docs` to your Postman.

## Run linter

1. for install linter localy `$ make lint-prepare`
2. for run linter `$ make lint`

## Run test

1. for install test localy `$ make test`

## How to use

1. run the kafka server using `$ docker-compose up -d`
2. run the http server using `$ make run-http`
3. run the proccesor server using `$ make run-processor`
