# PROJECT LATIHAN MAGANG

---

## Local development

* Copy file `config.example.yaml` to `config.yaml`
* Setup database (local) to create database schema equal with config `(schema: "project-latihan")`
* Setup Makefile param for migration connection database ex: `PG_DB_URL=postgresql://postgres:postgres@localhost:5439/project-latihan?sslmode=disable`
* run mod=vendor dependency with `make deps`
* Up / Run migrate with `make run pg.migrate.up` (Use WSL for OS Windows)
* run with `make run-service-local`

## License
[© 2023 taufikandrian]