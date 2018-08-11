# go-vue-starter

Copyright 2018 Sharful Islam

## Golang Starter project with Vue.js single page client

### Work in progress...

### Features:
- Middleware: [Negroni](https://github.com/urfave/negroni)

- Router: [Gorilla](https://github.com/gorilla/mux)

- Orm: [Gorm](https://github.com/jinzhu/gorm) (sqlite or postgres)

- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)

- [Vue.js](https://vuejs.org/) spa client with webpack

- User management

### TODO:
- config from file

- email confirmation

- logrus

- letsencrypt tls

### To get started:
Go starts up on port 3000 (and does not hot-reload), Vue client ( in a separate process - both must run alongside ) starts up on port 8080 (see below).  
``` bash
# clone repository
go get github.com/shantopagla/ShantoGoVue
cd $GOPATH/src/github.com/shantopagla/ShantoGoVue

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run server.go

# open a new terminal and change to the client dir
cd $GOPATH/src/github.com/shantopagla/ShantoGoVue/client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

### License

MIT License  - see LICENSE for more details