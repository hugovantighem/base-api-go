# Run the application

`make start-db` then `make run-scratch`


NOTE: the application requires a running postgreSQL that can be run using: `make start-db` (docker required).

NOTE: the application requires a `.env` file for variable configuration. You can set you own configuration based on `.env-example` file. `make copy-env-file` creates a `.env` file compatible with `make start-db` command.

NOTE: requires docker

curl commands for testing on local server are provided into `client` directory.

# Development
The `Makefile` contains targets to run individual target, or any target (`*-scratch`) along with all previously required targets as if you just cloned the repository.

`make install-deps` for generation tools

`make generate` for api and mock generation

`make build` to build binary

`make run` to run the application. 


# Tests
run unit tests
```
make test-scratch
```
