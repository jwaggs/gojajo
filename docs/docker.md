### docker cli helpers 

build `docker build -t gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest -f ${PIGGY_ROOT}/cmd/${PIGGY_CMD}/Dockerfile ${PIGGY_ROOT}`

push `docker push gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest`

run `docker run -e DATABASE_URL=$DATABASE_URL -e PORT=$PORT -p $PORT:$PORT gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest`