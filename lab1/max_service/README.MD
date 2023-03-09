## devops docker task
### Run from dockerhub:

```shell
docker run -d -p 80:80 --cpus="2" --memory="200m" maxutka/fist_task
```

### Or build locally:
```bash
docker build -t first-task .
docker run -d -p 80:80 --cpus="2" --memory="200m" first-task
```