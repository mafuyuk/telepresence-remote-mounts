# telepresence-remote-mounts
## DESCRIPTION
In the local container created by Telepresence, mount path is not in the intended place as it is like this
With this tool, you can achieve the expected behavior by putting a symbolic link on the intended path.

> Volume support requires a small amount of work on your part. The root directory where all the volumes can be found will be set to the TELEPRESENCE_ROOT environment variable in the shell run by telepresence. You will then need to use that env variable as the root for volume paths you are opening.  
> Telepresence will attempt to gather the mount points that exist in the remote pod and list them in the TELEPRESENCE_MOUNTS environment variable, separated by : characters. This allows automated discovery of remote volumes.
https://www.telepresence.io/howto/volumes

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
...Running...
...Add Symlink...
link:  /tmp/tel-k50p7t73/fs/etc/nginx/nginx.conf
mount:  /etc/nginx/nginx.conf
...Add Symlink...
link:  /tmp/tel-k50p7t73/fs/etc/nginx/conf.d
mount:  /etc/nginx/conf.d
...Add Symlink...
link:  /tmp/tel-k50p7t73/fs/var/run/secrets/kubernetes.io/serviceaccount
mount:  /var/run/secrets/kubernetes.io/serviceaccount
...End...

```

### with local change
When local and volume mounted, the creation of symbolic links by the tool at that location is skipped.
```bash
$ telepresence \
    --context your-context \
    --namespace your-namespace \
    --swap-deployment your-deployment:your-container \
    --docker-run --rm  -p 80:80 \
    --name frontend \
    --volume $GOPATH/bin/telepresence-remote-mounts:/bin/telepresence_remote_mounts \
    --volume $(pwd)/configs/front-nginx/nginx.conf:/etc/nginx/nginx.conf \
    nginx:latest

$ export CONTAINER_ID=$(docker ps --format "{{.ID}}" -f "name=frontend")
$ docker exec -it $CONTAINER_ID /bin/telepresence_remote_mounts
...Running...
unlinkat /etc/nginx/nginx.conf: device or resource busy
...Add Symlink...
link:  /tmp/tel-k50p7t73/fs/etc/nginx/conf.d
mount:  /etc/nginx/conf.d
...Add Symlink...
link:  /tmp/tel-k50p7t73/fs/var/run/secrets/kubernetes.io/serviceaccount
mount:  /var/run/secrets/kubernetes.io/serviceaccount
...End...
```
