alias k=kubectl
alias pigbuild='docker build -t gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest -f ${PIGGY_ROOT}/cmd/${PIGGY_CMD}/Dockerfile ${PIGGY_ROOT}'
alias pigpush='docker push gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest'

export PIGGY_CMD=api
pigbuild
pigpush

export PIGGY_CMD=scan
pigbuild
pigpush

k apply -f piggy-api-deployment.yml
k apply -f piggy-api-service.yml
