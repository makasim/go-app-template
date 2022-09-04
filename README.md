# go-app-template

The repo contains an app template that I use most of the time.

* It shows how to decouple app logic from main file. 
* How to write unit tests with mocks.
* How to write func tests.
* How to use Taskfile.


## Walk through 

### main.go file

The `main.go` file purpose is to:
* Set up logger (level, format, output and so on).
* Set up and validate app`s config (envs, json file and so on)
* Create root context.
* Create app instance, pass config and logger to it and run it. 
* Handle signals. Cancel the root context.
* It must exit with status one if `app.Run` return an error, otherwise zero. 

I put `main.go` file into a separate directory as a convention. 
All projects follow this rule and I have some scripts that rely on it.

### internal directory

This is the directory for app specific packages (if there are any).
No special rules. Just follow Go's advise on package naming.

### internal/app directory

The application we are going to run. 
It should contain an `App` and its `Config` structs.

### app.Run method

The method purpose is to:

* Set up any services the app needs.
* Resolve dependencies between services.
* Start all daemons, servers, workers and so on.
* Wait for the root context to be done.
* Gracefully stop services and other stuff in right order.

The app must be run synchronously from `main.go` file.
It must block until the root context is canceled.

### Unit tests

Run all tests:
```shell
task test:unit
```

Run some tests:
```shell
task test:unit -- ./internal/greeter/...
```

Filter tests:
```shell
task test:unit -- ./... -v -run TestGreeter_Greet/Hi
```

### Func tests

Run all tests:
```shell
task test:func
```

Run some tests:
```shell
task test:func -- ./tests/api/...
```

Enter test container:
```shell
task test:func:enter
# you'll get into container shell
# you can run there several commands 
task test:func 
task test:func -- ./tests/... -v 
```


