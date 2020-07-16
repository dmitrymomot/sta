# sta
Test task

## Running

### With docker

If you have installed docker, use make-helpers.

#### Build and run the API:
```
make build docker up
```

#### Down:
```
make down
```

#### To look at all available commands:
```
make help
```

### Without docker
In terminal from the root of this repo:
```
APP_PORT=8080 go run ./
```

## Dependencies
- github.com/go-chi/chi
- github.com/TV4/graceful (server graceful shutdown)

## Annotation
...