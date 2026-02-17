
### Dependencies for go + postgres + GORM
```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

`godotenv` lets you store database credentials in a `.env` file instead of hardcoding 




To get the DB working I did:

```bash
docker compose up -d

# connect to it (can just check with \dt when prompted, but is empty ofc)
docker exec -it api-project-postgres psql -U postgres -d api-project

# downloaded Chinook SQL file
curl -o chinook.sql https://raw.githubusercontent.com/neondatabase/postgres-sample-dbs/main/chinook.sql

# load it into docker postgres container
docker exec -i api-project-postgres psql -U postgres -d api-project < chinook.sql

# verify it worked
docker exec -it api-project-postgres psql -U postgres -d api-project
\dt

```


#### Useful docker commands
```bash
# Start Postgres
docker compose up -d

# Stop Postgres (data is preserved)
docker compose stop

# Stop and remove the container (data is still preserved in the volume)
docker compose down

# Wipe everything including data (useful for a fresh start)
docker compose down -v

# View Postgres logs if something goes wrong
docker compose logs postgres
```


