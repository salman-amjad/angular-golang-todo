# GO + Angular Todo

This is a simple Todo using GO in the backend and Angular 4 for client side.
It's uses the basic memory storage to maintain the list.

### Installation

Install [Go](https://golang.org/doc/install) and [Node.js](https://nodejs.org/)
Install [Angular CLI](https://github.com/angular/angular-cli) 
```sh
$ npm install -g @angular/cli
```

Clone the repo
```sh
$ git clone git@github.com:salman-amjad/angular-golang-todo.git
```

Install the dependencies and devDependencies and start the server.

```sh
$ cd angular-golang-todo
$ go get
$ cd client
$ npm install -d
```

## Development server

This can be done in two ways:
- Run `go run *.go`< this will start a server at `http://localhost:8080`
- Serve only Angular: `cd client`, `ng serve --watch`, this will serve client at `http://localhost:4200`
- Run using a task manager `gin --port 8000 --appPort 8080 run`, this will proxy our app at `http://localhost:8000` and watch for fiel chnages.

> Note: This code is just an example and isn't production ready. 

### Next
- Implement `MVC`
- Use `Mongo DB` for data storage.