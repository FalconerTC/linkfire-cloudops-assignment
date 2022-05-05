# linkfire-cloudops-assignment

# Run locally
To run the webserver locally, execute
```bash
go run .
```

# Running with docker
To build and run a docker image, execute
```bash
docker build -t linkfire-webapp .
docker run linkfire-webapp
```

# Releases
Releases are triggered but Github Actions when a tag is pushed. A docker image will be built and uploaded to Dockerhub


# Deployments
Deployments are executed manually using Github Actions. Specify the git tag to deploy
