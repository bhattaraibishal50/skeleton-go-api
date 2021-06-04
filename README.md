## API for behealth project

APi for be health project


### Development
- Copy `.env.example` to `.env` and update according to requirement.
- Create `serviceAccountKey.json` file for firebase admin sdk.
- To run `docker-compose up` (set configuration for adminer on `ADMINER_PORT` and debugger on `DEBUG_PORT`)
#### Migration Commands

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |
