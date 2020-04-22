# telepresence-remote-mounts
##  INSTALLATION
### build
```bash
$ git clone https://mafuyuk/telepresence-remote-mounts .
$ make
$ cp bin/telepresence-remote-mounts $GOPATH/bin
```

## Usage
```bash
$ telepresence \
    --context your-context \
    --namespace your-namespace \
    --swap-deployment your-deployment:your-container \
    --docker-run --rm  -p 80:80 \
    --name frontend \
    --volume $GOPATH/bin/telepresence-remote-mounts:/bin/telepresence_remote_mounts \
    nginx:latest

$ export CONTAINER_ID=$(docker ps --format "{{.ID}}" -f "name=frontend")
$ docker exec -it $CONTAINER_ID /bin/telepresence_remote_mounts
```
