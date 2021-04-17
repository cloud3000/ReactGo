# Getting Started with ReactGo

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

#### The thought here is that Golang makes a better backend than the commonly used languages, and React is by far the winner of the frontend.

The Go server will NOT run using `yarn start`, so all script references have been removed from package.json except build, so use `yarn build` to build th React frontend, Go has it's own build command.

### **Prerequisites**

1. You have installed nodejs with yarn or npm
2. You have install Go (the Golang compiler and dev tools) to build main.go

> With commit title "Implemented //Go Embed::", and commits thereafter, you will need Go 1.16 or higher. If you are using an older version of Go you'll need to use a previous commit, So after you clone this repo and cd into the ReactGo directory, use the following 'git' command: (and 'oh yeah' you'll need to have 'git' installed)

```
git checkout cb3c6ed00631ab12b56b7ea4cd98c53c16eb2eff
```

To try ReactGo, after you clone this repo, you only need three commands `yarn build`, `go build`, and `./ReactGo`.

1. `yarn build` builds the production React frontend.
2. `go build` compiles the main.go source into a executable file 'ReactGo'
3. `./ReactGo` Starts the server, listening on port 3000

The `yarn build` command builds the app for production to the `build` folder.
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.

NOTE: Your app is NOT ready to be deployed yet!
If the `yarn build` is successful, next you will build/compile the Golang server using `go build`

You must cd into src to compile using `go build` because **main.go** source file is in the 'src' directory. the 'go build' command wil create the **ReactGo** binary executable. Depending on which OS you are using, the executable filename may be different. ea Windows will add .exe to the end.

# **IMPORTANT NOTE**

In main.go, the following statement is used to define the path where the files are to be served from:

```go
fs := http.FileServer(http.Dir("../build"))
```

If all goes well, then start the Go server at the command-line by typing `./ReactGo`

Then navigate your browser to `http://localhost:3000`

> This path is a relative path, and because ReactGo is in the 'src' folder, then `../build` is the correct relative path. If you decide to move 'ReactGo' to the parent directory of 'src', then `./build` will be the correct relative path. It's relative to whatever path 'pwd' is pointing to when you execute the ReactGo server. So if you get a error 404, it's probably related to the relative path.

> This source has been tested on Linux, not Windows or Mac. Therefore you may need to play with the path `../build` It will probably work just the way it is on Mac, but if it doesn't work on Windows, you might want to try something like `..\build` or just `build`, just stay away from trying the full path, as that's just a bad idea.
> git checkout
