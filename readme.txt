#build the project

docker build -t go-app-normal:latest .

docker run -d -p 8081:8081 --name web go-app-normal:latest

docker logs -f web

minikube ip

#remove the docker container - docker rm -f web

# check the go-app image size - docker image ls 
