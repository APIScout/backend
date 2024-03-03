# API Scout - Backend

## Set ENV Variables

The following are the variables that need to be set:

- `GIN_MODE`: In which mode Gin should be running (can be `release` or `debug`)
- `MODELS_HOST`: The hostname of the DL models container (use `models` if in release mode, `127.0.0.1` if in debug mode)

## Downloading USE Model

To download the Universal Sentence Encoder (USE) model, run the python script in `scripts/download-use.py` by running the following commands (you should run these commands while in the `backend` directory):

```bash
conda env create --file=environment.yml
conda run -n api-scout python ./scripts/download-use.py
```

You will now have a new directory in `models` called `universal-encoder`. This model will be used by the `docker-compose.yml` file to serve the model in a container.

## Spinning up the Containers

For replication purposes, in this repo you will find both a `Dockerfile` and a `docker-compose.yml` file. The `Dockerfile` will create a Docker image with a build of the Golang backend in it. To create the image, simply run:

```bash
docker build -f Dockerfile -t api-scout-backend:latest .
```

Once the image and its dependencies have been downloaded, you can now spin up the backend containers. To do that, simply run:

```bash
docker-compose up -d
```

This will create a Docker container for the Golang backend and for the USE model. The backend will now be able to make HTTP calls to the USE model to embed queries and documents. N.B.: The USE container will not expose any ports, it will be called locally by the backend by means of the `be-network` shared network.

## Dependencies

For the backend to work, both an ElasticSearch instance and a MongoDB instance should be up and running.
