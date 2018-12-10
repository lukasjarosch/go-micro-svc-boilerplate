# go-micro service boilerplate

This is a boilerplate backend service using `go-micro`.
Goal for this project is to create a good starting-point for new microservices - batteries included.

Currently the project is under heavy development and may change at any time.
Feel free to contribute.

If you do not need to use RPCs, I will also provide a boilerplate project for **go-web**

## Getting started
All you need to do is to checkout this repo and start hacking. But for the sake of simplicity let's just walk through all steps required to customize the template.

1. Adjust `DOCKER_IMAGE` and `DOCKER_TAG` in the **Makefile** to your  needs
2. Create your own protobuf definitions and place them inside a folder in `proto/`
3. Adjust the `proto` command in the *Makefile* to match `go_out` to your protobuf setup
4. Change the service name inside `main.go`
5. Create the handler(s) you need inside `handler/` and register them in `main.go`

## Development setup
Although the target-platform is Kubernetes, we can omit it for local development and only require docker to be setup.

The `deploy/docker-compose.dev.yml` file contains all services required for local development (consul, micro, mysql, etc.) 
You can either start them normally or use the provided Makefile with `make start-local` and `make stop-local`.

Once started, you will have a running **consul** on `localhost:8500` as well as the **micro-api** running on `localhost:8080` in **rpc** mode.

### Building & running
Running and building the service is completely taken care of by the Makefile. One thing to keep in mind is that the service is configured
to run in *kubernetes-native* mode. Thus you will need to tell it that we are running without K8s by setting the environment variable `ENVIRONMENT=local`

To run your service then, use `make run` and to build a statically linked binary use `make build` 

It's just as easy to build a docker image, just use `make docker` and `make docker-run` to start the container (**Note:** The container is started in host-network mode).

## Deployment
The boilerplate ships with configured `deploy/k8s/deployment.yaml` and `deploy/k8s/service.yaml` which you can use to deploy the service into your Kubernetes cluster.


## Configuration
The boilerplate is preconfigured using **go-config**. It has two sources attached: *file* and *env* (in that order).
This means that you can always overwrite configuration defaults provided by the config.json by setting environment variables accordingly.

For example: It might be a good idea to change the log-format when deploying to your cluster.  The current default is **text** which is meant for development.
To change it to JSON logs, either change the config file or just overwrite it by setting `LOG_FORMAT=json`.
