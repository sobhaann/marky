# TECH STACK:
  - database: postgres
  - web framework: golang `net/http`
  - migration tool: [goose](https://pressly.github.io/goose/)
  - sql builder: [sqlc](https://docs.sqlc.dev/en/stable/tutorials/getting-started-sqlite.html)
  - log toolkit: golang `log/slog`...
  - config format: `toml`


## TODO:
- [] add route directory to separate public and private route
- [x] define user schema (because we are using `sqlite` i should completely ensure what i want to add to my users tables)
- [] implement a functional and pretty log system via `log/slog`
- [x] dockerize project
- [] implement oauth as an option for our auth
- [] implement the crud apis for markdown part
- [] render and show markdown file
- [] add an optional `-config` flag that user put its config file path in there. if it is empty, we use the `./config/config.toml` path
