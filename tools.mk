GO_BUILD_ENVIRONMENT=GOOS=linux #GOARCH=amd64 this was removed as there is an issue when 64bit binary is built

DOCKER_BUILD_CMD=docker build --no-cache