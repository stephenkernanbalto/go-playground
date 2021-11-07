### Project Structure

#### Modules vs Packages

Go programs are organized into packages, which are then organized further into modules. A package is a collection of
files in the same directory that get compiled together. Packages can then be imported from anywhere within a module.

Modules are collections of related packages that get released together.

In terms of our `/web` directory, the `/ui`, `/tests`, and `/public` directories would all be *packages*, while
the `/web` root directory is a `module`.

In Go, all modules will have a `go.mod` file that specify the module's import path. For instance, in the `/hello`
module, the first line of the `go.mod` reads as follows:

```gotemplate
module go_playground/hello
```

This is typically a place where Go's import tools can reach the module, such as a GitHub repository.

## TODO: importing by path from private repo

#### Basic Project File Structure

```
- module_name
    \_ go.mod
    \_ go.sum
    \_ main.go
    \_ package_name
        \_ package_name.go
    \_ second_package
        \_ second_package.go
```

* `go.mod` - module config file that declares module name, imports, and other things
* `go.sum` - dependency list
* `main.go` - package named `main` that lives at module root and is the executable when `go run .` is called.

#### Complex Project Structure

As the number of packages, entry point files, and private files increase, it can be helpful to expand the structure to
subdirectories. [Many articles](https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/) recommend a
pattern like this:

```
- module_name
    \_ go.mod
    \_ go.sum
    \_ cmd
        \_ file_name
            \_ main.go
        \_ second_file_name
            \_ main.go
    \_ pkg
        \_ package_name
            \_ package_name.go
        \_ second_package
            \_ second_package.go
    \_ internal
        \_ private_name
            \_ ...
        \_ second_private_name
```

### Package Management

#### Setting Up Project for Package Management

In order to enable package management in your Go project, you must create a `go.mod` file. Use the following command:

```shell
    go mod init {your_module_name}
```

You'll notice that you have to name your module.
Here's [a quick guide](https://golang.org/doc/modules/managing-dependencies#naming_module) on that.

#### Using Internal Packages

Any packages in a module can be imported by other packages in the module. To import a local package, you can use the
module's file path as declared in the `go.mod` file. For instance, if the `go.mod` is listed
as `module example.com/module`, then you can import a package like this:

```gotemplate
import "example.com/module/local_package"
```

#### Downloading External Packages for Your Project

Before using methods from an external module in your code, you will need to run `go mod tidy` so that Go knows to
download external packages for use. This only needs to be done when introducing new packages.

#### Checking Dependencies

***go why***

If you're curious why a specific dependency ended up in your `go.mod` file, you can run the following command do explain
where the dependency came from:

```shell
go why {package_name}
```

***go vendor***

If you want to track all of your dependencies, you can obtain a local copy of all dependencies (similar to node_modules)
by running:

```shell
go mod vendor
```

This creates a `vendor` directory with a `modules.txt` file that shows each of your dependencies along with their peer
dependencies.

#### Using Internal Modules

If your module is not published in a place that Go tools can find it, you'll need to point the module's path name to
your local directory. This is done with the `go mod edit` command with the following pattern:

```shell
go mod edit -replace {imported-module-path-name}={actual-file-path}
go mod edit -replace example.com/greetings=../greetings
```

This adds the following line to your `go.mod` file:

```gotemplate
replace example.com/greetings => ../greetings
```

Then, just like above, you can run `go mod tidy` to have Go bring in that module.

### Running an Application

When making frequent local changes, it is usually sufficient to run an application without compiling or installing the
packages.

To run a file in this basic way, use the `go run` command in the same way that you would the `python` CLI command:

```shell
go run {pathname}
```

This command runs the `main` package in your directory.

### Compiling and Installing Application

While the `go run` command is extremely useful for making quick changes, this command *does not* produce a binary
executable. When you want to create that binary executable, you must use one of the following commands:

* `go build` - compiles the packages, along with dependencies, but it doesn't install the results
* `go install` - compiles and installs the packages

#### _go build_

1. From the command line in the package directory, run:

   ```shell
      go build
   ```

2. This will create a `{module_name}.exe`.
    1. To run on Linux/MacOS:
    ```shell
      ./{module_name}
    ```
    2. To run on Windows:
    ```shell
      {module_name}.exe
    ```

3. Adding Go install directory to your system's shell path
    1. Discover the path with this:
   ```shell
      go list -f '{{.Target}}'
   ```
   > **_NOTE_**: This will output something like `/go/bin/{module_name}`. Your path name *does not* include the `module_name`. In this case, the pathname you'll need is `/go/bin`.
    2. Add directory to your shell path
        1. On Linux or Mac, run the following command:
       ```shell
          export PATH=$PATH:/path/to/your/install/directory
       ```
        2. On Windows, run the following command:
       ```shell
         set PATH=%PATH%;C:\path\to\your\install\directory
       ```

4. Once you've updated shell path, you can install the package like this:
    ```shell
      go install
    ```

### Websockets in Go

While Go has its own websocket library, the external package from Gorilla has an active community and provides a more
pleasant developer experience. The documentation can be
found [here](https://pkg.go.dev/github.com/gorilla/websocket#pkg-overview).

To install the package, you must do the following:

1. Install the package by downloading it in the target directory with this command:

```shell
go get github.com/gorilla/websocket
```

2. Import the package in your file

```gotemplate
// main.go
import "github.com/gorilla/websocket"
```

