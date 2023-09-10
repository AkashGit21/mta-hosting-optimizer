# mta-hosting-optimizer

## Short description
Currently, about 35 physical servers host 482 mail transfer agents (MTAs) each having a dedicated public IP address. To be economical while hosting MTAs as a software engineer I want to create a service that uncovers the inefficient servers hosting only few active MTAs.

## Acceptance Criteria

- [X] Project with the name “mta-hosting-optimizer”.
- [X] A major programming language (Golang) is utilized to build the service.
- [X] A HTTP/REST endpoint to retrieve hostnames having less or equals X active IP
addresses exists.
- [X] X is configurable using an environment variable, the default value is set to 1.
- [X] IP Configuration data (IpConfig) is provided by a mock service using sample data below.
- [X] Unit & integration tests are present. 
- [ ] Code coverage exceeds 90%.
- [X] Integrate the test and build phases to Github action.

## Setup required
1. Golang v1.20+ (or 1.16+ should also work)
1. make (to execute the commands from Makefile)

## Steps to run the application locally
1. Fetch all golang dependencies by typing the following from root directory:
    ```sh
        go mod download
    ```
1. Configure the sample data in similar format in the `ipconfig/data.json` file. 
1. Run the application by typing the following from root directory:
    ```sh
        make run
    ```
1. Navigate from your browser to `http://localhost:8082/hosts/inefficient`.

## Features / Points added
1. Logging mechanism
1. Panic recover handler
1. Github workflow -> test and build phases
1. env file separation
1. Unit/Integration test cases

## Scope of Improvement
1. Addition of **Dockerfile** to enhance collaboration.

