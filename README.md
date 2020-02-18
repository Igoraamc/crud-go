# GO CRUD

GO CRUD is a simple crud written in [GO](https://golang.org/)

## Configuration

Change your $GOPATH to your workspace. This can be made with this command: 

```bash
    go env -w GOPATH=$HOME/Desktop/go
```

## Usage

After the configuration, you can start the project with the command:

```bash
    go run main.go
```

## MongoDB (Docker required for this steps)

You can install a mongodb image with docker. After installing docker you need to install a mongo image with this:

```bash
    docker pull tutum/mongodb
```

After installing the mongo image from docker, you need to run the image with this command:

```bash
    docker run -d -p 27017:27017 -p 28017:28017 -e AUTH=no tutum/mongodb
```

With the mongo image running, check the ID of the image running and then start the image with:

```bash
    docker ps -a // check images running

    docker start 611g6dc90o47 // start the image with the CONTAINER ID
```

And finally check if mongo is running with:

```bash
    mongo
```