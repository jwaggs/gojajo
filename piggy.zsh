# some script shorthands
alias k=kubectl
alias pigbuild='docker build -t gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest -f ${PIGGY_ROOT}/cmd/${PIGGY_CMD}/Dockerfile ${PIGGY_ROOT}'
alias pigpush='docker push gcr.io/piggy-police/piggy-${PIGGY_CMD}:latest'

# delete everything in the default namespace
#k delete all --all
#kubectl delete namespace piggy
#kubectl create namespace piggy

k delete deployment piggy-api
k delete svc piggy-api
k delete cronjob piggy-scan

# TODO: apply secrets

# rebuild our api and push to conatiner registry
export PIGGY_CMD=api # our aliases use this env var to determine which dockerfile to build
pigbuild # see the alias above
pigpush  # see the alias above

# apply api deployment spec (contains the container images and number of replicas necessary for pods)
k apply -f piggy-api-deployment.yml
# apply api service (exposes the api deployment to other pods in the cluster)
k apply -f piggy-api-service.yml

# rebuild our scan cronjob and push to conatiner registry
export PIGGY_CMD=scan
pigbuild
pigpush

# apply the scan cronjob
k apply -f piggy-scan-chase.yml
