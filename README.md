# veil
docker build -t veil .
docker run --publish 6060:8080 --name veil --rm veil

# To ssh into container, first get the container id
docker ps

# then run to get a bash shell in the container.
docker exec -it <container name> /bin/bash

# generically run this to run whatever command
docker exec -it <container name> <command>