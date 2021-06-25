# ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like setting up the database or cache server instance and storing migrations.

- `./platform/cache` folder with in-memory cache setup functions
- `./platform/database` folder with database configuration
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
