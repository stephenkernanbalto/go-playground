* [Managing Dependencies](https://golang.org/doc/modules/managing-dependencies#naming_module)

    When you run go mod init to create a module for tracking dependencies, you specify a module path that serves as the module’s name. The module path becomes the import path prefix for packages in the module. Be sure to specify a module path that won’t conflict with the module path of other modules.

    At a minimum, a module path need only indicate something about its origin, such as a company or author or owner name. But the path might also be more descriptive about what the module is or does.

    The module path is typically of the following form:

    `<prefix>/<descriptive-text>`

* [Redis in Go](https://github.com/go-redis/redis)
* [Postgres in Go](https://bun.uptrace.dev/postgres/#pgdriver)
* [Logging in Go](https://www.datadoghq.com/blog/go-logging/#use-logrus-for-formatted-logs)
* [Websockets in Go](https://github.com/gorilla/websocket/)
* [Time in Go](https://pkg.go.dev/time)