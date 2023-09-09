# mta-hosting-optimizer

## Short description
Currently, about 35 physical servers host 482 mail transfer agents (MTAs) each having a dedicated public IP address. To be economical while hosting MTAs as a software engineer I want to create a service that uncovers the inefficient servers hosting only few active MTAs.

## Acceptance Criteria

- [X] Project with the name “mta-hosting-optimizer”.
- [X] A major programming language (Golang) is utilized to build the service.
- [ ] A HTTP/REST endpoint to retrieve hostnames having less or equals X active IP
addresses exists.
- [X] X is configurable using an environment variable, the default value is set to 1.
- [ ] IP Configuration data (IpConfig) is provided by a mock service using sample data below.
- [ ] Code coverage exceeds 90%. 7. Unit & integration tests are present.
- [ ] Integrate the test and build phases to Github action.
