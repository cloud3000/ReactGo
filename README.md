# Getting Started with ReactGo

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

#### The thought here is that Golang makes a better backend than the commonly used languages, and React is by far the winner of the frontend battle.

The Go server will NOT run using `yarn start`, so all script references have been removed from package.json except build, so use `yarn build` to build th React frontend, Go has it's own build command.

### **Prerequisites**

1. You must have nodejs, with yarn or npm, installed.
2. You must have Go (the Golang compiler and dev tools) installed to build main.go

> With the commit that begins with "Implemented //go.embed", and commits thereafter, you will need Go 1.16 or higher. If you are using an older version of Go you'll need to use a previous commit, So after you clone this repo and cd into the ReactGo directory, use the following 'git' command: (and 'oh yeah' you'll need to have 'git' installed).

```
git checkout cb3c6ed00631ab12b56b7ea4cd98c53c16eb2eff
```

The commit above was the last commit that did NOT require Go 1.16, and therefore did NOT import "embed". I highly recommend that you don't use the older commit, and just upgrade to Go 1.16 or higher and try the "embed" package.

## What is //go.embed ?

**I should explain:** Go source files that import "embed" can use the //go:embed directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time.

This means that you can embed an entire file system (type FS) into the executable load file produced by the Go compiler.

**For Example:** This project serves static html/css/javascript generated by the React build script, so all the production React files are compiled into the Go executable.

**This is just a small example** Theoretically I should be able to develope a large application with hundreds of React pages, thousands of components, including any content of valid mime types on the Frontend, Go code on the Backend which includes a HTTP/HTTPS server and router, all compile into a single executable file.

Try it yourself, after you clone this repo, you only need three commands `yarn build`, `go build`, and `./ReactGo`.

1. `yarn build` builds the production React frontend.
2. `go build` compiles the main.go source into a executable file 'ReactGo'
3. `./ReactGo` Starts the server, listening on port 3000

> To prove it works, transfer the executable `ReactGo` to another machine, one that doesn't have nodejs or the go compiler and devtools on it. It wil just work, without any missing dependencies, because they are all in the executable.

The `yarn build` command builds the app for production to the `build` folder.
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.

NOTE: Your app is NOT ready to be deployed yet!
If the `yarn build` is successful, next you will build/compile the Golang server using `go build`

Depending on which OS you are using, the executable filename may be different. ea Windows will add .exe to the end.

# **IMPORTANT NOTE**

In main.go, the following statements are used to define the path where the files are to be compiled from:

```go
//go:embed build
var f embed.FS

	contentStatic, err := fs.Sub(fsys, "build")
	fs := http.FileServer(http.FS(contentStatic))
```

The Go linter in VS Code and the Go compiler will complain if the path is wrong, so if it compiles clean, it should execute without 404 errors. Start the Go server at the command-line by typing `./ReactGo`

Then navigate your browser to `http://localhost:3000`

> This path is a relative path, `ReactGo` is in the same directory as the `build` directory, so "build" will be the correct relative path at compile time. It's relative to whatever path 'pwd' is pointing to when you compile main.go After compile, it doesn't matter where you run the ReactGo server. It's a little more complicated using the gorilla mux.

> This source has been tested on Linux, not Windows or Mac. Therefore you may need to play with the path `build` It will probably work just the way it is on Mac, and windows. The Go docs say that the "embed" package makes Windows, Mac, and Linux file systems look the same when embedded.

I am working on larger example of this using "embed" with "gorilla/mux". It will have a REST api, with user login and registration, a blog, and will be compatible with Sqlite, MySQL, and Postgres. All compiled in to one executable file. Additionally the React Frontend, and the Backend REST api will be severed from one port. This will be a simpler and more a secure method for production web applications.
