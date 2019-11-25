### docker cli helpers 

build `docker build -t gcr.io/piggy-police/pig-api:latest -f ${PIGGY_ROOT}/cmd/server/Dockerfile ${PIGGY_ROOT}`

push `docker push gcr.io/piggy-police/pig-api:latest`

run `docker run -e DATABASE_URL=$DATABASE_URL -e PORT=$PORT -p $PORT:$PORT gcr.io/piggy-police/pig-api`