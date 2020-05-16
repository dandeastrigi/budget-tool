# Budget Tool

----

## App

Inside of the *app/src* directory run the follow command:

```bash
npm install
```

Back inside of the *app* directory, and run:

```bash
docker-compose up --build
```

## Database

Go inside *database* path, then run:

```bash
docker-compose up --build
```

## Service

Inside of the *service* directory run the follow command:

```go
go build .
go run server.go
```

Or you can run:

```bash
bash start.sh
```

### Using
Open your browser at *http://localhost:4200*