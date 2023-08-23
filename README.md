# short-it

## Start
- clone repo `git clone git@github.com:daussho/short-it.git`
- install dependencies `go mod tidy`
- copy env `cp .env.example .env`
- run migration
	- `make run-migration`
		- `0001` > create table `users` & `urls`
		- `0002` > insert user with credentials `admin:admin`
- build binaries `make build-http`
- run binaries either directly or use `pm2`
	- `./bin/short-it`
	- `pm2 start bin/short-it`