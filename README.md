
<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>
go-auth-middleware
</h1>
<h3 align="center">📍 Secure Your Go with go-auth-middleware!</h3>
<h3 align="center">⚙️ Developed with the software and tools below:</h3>

<p align="center">
<img src="https://img.shields.io/badge/GNU%20Bash-4EAA25.svg?style=for-the-badge&logo=GNU-Bash&logoColor=white" alt="GNU%20Bash" />
<img src="https://img.shields.io/badge/Redis-DC382D.svg?style=for-the-badge&logo=Redis&logoColor=white" alt="Redis" />
<img src="https://img.shields.io/badge/V8-4B8BF5.svg?style=for-the-badge&logo=V8&logoColor=white" alt="V8" />
<img src="https://img.shields.io/badge/MySQL-4479A1.svg?style=for-the-badge&logo=MySQL&logoColor=white" alt="MySQL" />
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=for-the-badge&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens" alt="JWT" />
<img src="https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white" alt="SwaggerUI" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style=for-the-badge&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
</div>

---

## 📚 Table of Contents
- [📚 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [💫 Features](#-features)
- [📂 Project Structure](#project-structure)
- [🧩 Modules](#modules)
- [🚀 Getting Started](#-getting-started)
<!-- - [🗺 Roadmap](#-roadmap) -->
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---


## 📍 Overview

This Go-based authentication server includes a package for handling HTTP requests and several endpoints for user authentication and management. It uses Redis for role-based access control and a MySQL database for user data storage. The project's value proposition lies in its robust JWT authentication middleware, which ensures secure authentication for users accessing web applications built with this server. It also includes a Dockerized environment for easy deployment.

---

## 💫 Features

Feature | Description |
|---|---|
| **🏗 Structure and Organization** | The codebase is well-structured and organized, with separate folders for Docker configurations, database scripts, and application files. The use of a Go module makes dependency management straightforward. |
| **📝 Code Documentation** | The codebase is well-documented, with comments explaining the purpose and usage of functions and variables throughout the application. |
| **🧩 Dependency Management** | The codebase uses the Go module system for versioned dependency management, with clear dependencies outlined in the go.mod file. |
| **♻️ Modularity and Reusability** | The codebase is modular and follows the Single Responsibility Principle, with separate packages for server logic, middleware, models, and utilities, making it easily extendable and reusable. |
| **⚡️ Performance and Optimization** | The codebase leverages Redis caching for role-based access control, making the application faster and more efficient, as well as providing persistence for database connections. |
| **🔒 Security Measures** | The codebase follows best practices for secure authentication, hashing passwords and storing them securely, generating JWT tokens for authorized users, and checking the validity of API requests. |
| **🔄 Version Control and Collaboration** | The codebase is stored in a GitHub repository, with clear commit messages and a branching strategy using pull requests for collaboration and code review. |
| **🔌 External Integrations** | The codebase integrates with Redis and MySQL for database functionality, as well as using SwaggerUI for API documentation. |
| **📈 Scalability and Extensibility** | The codebase lends itself well to scalability and extensibility, with the ability to easily add new endpoints, middleware, and functionality as needed without disrupting existing code. |

<!-- | **✔️ Testing and Quality Assurance** | The codebase includes both unit tests and integration tests for the server and middleware packages, and uses a linter (golangci-lint) and a formatter (go fmt) to ensure code quality. |
-->
---


<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-github-open.svg" width="80" />

## 📂 Project Structure


```bash
repo
├── LICENSE
├── README.md
├── app
│   ├── Dockerfile
│   ├── config
│   │   └── config.go
│   ├── go-init.sh
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── models
│   │   └── user.go
│   ├── server
│   │   ├── api
│   │   │   └── api.yaml
│   │   ├── handlers
│   │   │   ├── auth.go
│   │   │   ├── rbac.go
│   │   │   └── user.go
│   │   ├── middleware
│   │   │   └── jwt.go
│   │   └── server.go
│   └── utils
│       ├── jwt.go
│       └── logger.go
├── db
│   ├── mysql
│   │   ├── Dockerfile
│   │   └── init_template.sql
│   └── redis
│       └── redis_template.conf
└── scripts
    ├── docker_middleware.sh
    ├── docker_mysql.sh
    ├── docker_redis.sh
    ├── prepare_env.sh
    ├── prepare_mysql_init.sh
    └── prepare_redis_init.sh

12 directories, 26 files
```

---

<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-src-open.svg" width="80" />

## 🧩 Modules

<details closed><summary>App</summary>

| File       | Summary                                                                                                                                                                                                                                                                                                                                   | Module         |
|:-----------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------|
| go.mod     | The code snippet is a module file that specifies the required packages and their respective versions for a Go project. It includes packages for handling JWT authentication, Swagger UI, Redis, MySQL, and logging. The module also includes indirect dependencies for xxhash and rendezvous hashing.                                     | app/go.mod     |
| go-init.sh | This Bash script initializes a Go module named "app" using the "go mod init" command.                                                                                                                                                                                                                                                     | app/go-init.sh |
| Dockerfile | The code snippet defines a Docker container that builds a Go server application. It copies the necessary files and dependencies, downloads Go modules, builds the server using the command "go build", and exposes port 3456 to run the server. Finally, the command "CMD ["./auth_server"]" is used to run the built server application. | app/Dockerfile |
| main.go    | The provided code snippet initializes the server, connects to a Redis RBAC database using the provided configuration, and starts listening on port 3456. It uses a logger to record errors and reports a successful connection to the database.                                                                                           | app/main.go    |

</details>

<details closed><summary>Config</summary>

| File      | Summary                                                                                                                                                                                                                                                                                                                                                                                  | Module               |
|:----------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------|
| config.go | This code snippet defines a Config struct and provides methods for loading configuration values from environment variables, retrieving database and Redis connection details, and getting the application's secret key. The Config struct contains fields for all the required configuration values. The getEnv function helps retrieve the value of the environment variables provided. | app/config/config.go |

</details>

<details closed><summary>Handlers</summary>

| File    | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Module                      |
|:--------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------|
| user.go | The code snippet provides two handler functions, GetUserHandler and UpdateUserHandler, for handling GET and PUT requests respectively. Both handlers verify a token obtained from the request header and retrieve the user information from the database based on the username extracted from the token's claims. GetUserHandler returns the user data as a response, while UpdateUserHandler updates the user's password field and saves the changes to the database. | app/server/handlers/user.go |
| auth.go | The code snippet provides handlers for handling login and signup requests. The LoginHandler verifies user credentials, retrieves the user from the database, generates a JWT token, and returns it as a response. The SignupHandler creates a new user, saves it to the database, sets the user's role in Redis, generates a JWT token, and returns it as a response. Both handlers use helper functions to hash and check passwords and encode JSON responses.        | app/server/handlers/auth.go |
| rbac.go | The provided code snippet contains HTTP request handlers for user promotion, demotion, and privilege checking. These handlers interact with Redis for updating and retrieving user roles, and they also validate user authentication by checking the presence and validity of JWT tokens. Additionally, the handlers return JSON responses to indicate the success or failure of these actions.                                                                        | app/server/handlers/rbac.go |

</details>

<details closed><summary>Middleware</summary>

| File   | Summary                                                                                                                                                                                                                                                          | Module                       |
|:-------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------|
| jwt.go | The code snippet defines a middleware function that checks for a valid JWT token in the request headers or cookies. If the token is missing or invalid, it returns an unauthorized HTTP status code. The valid token is passed to the next handler in the chain. | app/server/middleware/jwt.go |

</details>

<details closed><summary>Models</summary>

| File    | Summary                                                                                                                                                                                                                                                                                                | Module             |
|:--------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------|
| user.go | This code defines a User struct that can be used to interact with a MySQL database using methods for CRUD operations: Save (INSERT), GetUserByUsername (SELECT), and Update (UPDATE). It uses the database URL and table name retrieved from a configuration file and logs all actions using a logger. | app/models/user.go |

</details>

<details closed><summary>Mysql</summary>

| File              | Summary                                                                                                                                                                                                                                                                                                                                               | Module                     |
|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------------|
| init_template.sql | The code snippet creates a new database, table, user account, and assigns privileges to manage the database and table. The user account is granted full privileges for both the database and specific table. The code is used to set up a relational database management system.                                                                      | db/mysql/init_template.sql |
| Dockerfile        | The code snippet is for building a Docker image with MySQL version 8.0.33. The password for the root user is set to the value of the $MYSQL_ROOT_PASSWORD environment variable. An initialization SQL file is copied to the /docker-entrypoint-initdb.d/ directory to be executed upon database start. Port 3306 is exposed for external connections. | db/mysql/Dockerfile        |

</details>

<details closed><summary>Redis</summary>

| File                | Summary                                                                                                                                                                                                                                                                                                                                                                        | Module                       |
|:--------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------|
| redis_template.conf | The provided code snippet contains configuration options for running a Redis server, including setting a password, specifying network settings, enabling persistence, setting logging preferences, and defining maximum client limits. It also defines save modes for automatic data snapshotting based on the number of keys that have changed within certain time intervals. | db/redis/redis_template.conf |

</details>

<details closed><summary>Scripts</summary>

| File                  | Summary                                                                                                                                                                                                                                                                                                                                                   | Module                        |
|:----------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------------------------------|
| prepare_mysql_init.sh | This Bash script sources another script, prepares the file paths for input and output SQL files, and replaces environment variable placeholders in the input SQL file with their corresponding values. The "envsubst" command is used to perform the replacement.                                                                                         | scripts/prepare_mysql_init.sh |
| docker_middleware.sh  | The code snippet is a bash script that prepares the environment by sourcing some variables, cleans previous docker instances, builds and pushes a docker image, and then runs it with specific configurations and environment variables. The script also dynamically configures the nginx proxy to be in the same network as the host.                    | scripts/docker_middleware.sh  |
| prepare_env.sh        | This code snippet sets environment variables used for a MySQL database and an authentication server, including hostname, port, secret key, and admin token. It also configures integration between the app and MySQL and Redis by setting database and table names, usernames, passwords, and Redis password.                                             | scripts/prepare_env.sh        |
| docker_redis.sh       | The code snippet is a Bash script that prepares Redis configurations, removes any existing Redis containers and images, and starts a new Redis container with persistence and network configuration parameters. The Redis container is named REDIS-RBAC-LOCAL and exposed on ports 6379 and 8001. The script is meant to be used in a Docker environment. | scripts/docker_redis.sh       |
| prepare_redis_init.sh | This Bash script sources a separate script (prepare_env.sh) and then generates a Redis configuration file using environment variables substituted into a template config file (redis_template.conf). The generated config file is saved to a specified location (redis.conf).                                                                             | scripts/prepare_redis_init.sh |
| docker_mysql.sh       | This script prepares and runs a MySQL database in a Docker container. It first builds and pushes an image, then reads environment variables before running the container with specified configurations. The script allows the database to be accessed from the host network and with a specific root password.                                            | scripts/docker_mysql.sh       |

</details>

<details closed><summary>Server</summary>

| File      | Summary                                                                                                                                                                                                                                                                                                                                                               | Module               |
|:----------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------|
| server.go | The code snippet is a package for a server that handles HTTP requests and has several endpoints for user authentication and management. It includes a SwaggerUI handler, a JWT middleware, and handlers for login, signup, user retrieval and modification, user promotion and demotion, and user checking. It also includes a Redis client for database interaction. | app/server/server.go |

</details>

<details closed><summary>Utils</summary>

| File      | Summary                                                                                                                                                                                                                                                                                                                             | Module              |
|:----------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------------------|
| logger.go | The code defines a package "utils" that imports logrus and os packages. It sets up a logger instance with JSON formatter, debug log level and output to standard output. The package provides a GetLogger function that returns the logger instance for logging purposes.                                                           | app/utils/logger.go |
| jwt.go    | The provided code snippet contains functions for extracting, generating, verifying, and retrieving claims from JWT tokens. It also initializes the secret key for token signing based on a configuration file. These functionalities are encapsulated in the `utils` package and can be utilized in other parts of the application. | app/utils/jwt.go    |

</details>

---

## 🚀 Getting Started

<!-- ### ✅ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:
> - [📌  PREREQUISITE-1]
> - [📌  PREREQUISITE-2]
> - ... -->

### 🖥 Installation

1. Clone the go-auth-middleware repository:
```sh
git clone ../go-auth-middleware
```

2. Change to the project directory:
```sh
cd go-auth-middleware
```

3. Install the dependencies:
```sh
go build -o myapp
```

### 🤖 Using go-auth-middleware

```sh
./myapp
```

### 🧪 Running Tests
```sh
go test
```


---

## 🗺 Roadmap

> - [ ] [📌  Task 1: Implement Unit Tests]
> - [ ] [📌  Task 2: Run Go Build Validation in GitHub]

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a pull request to the original repository.
Open a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `Apache-2.0` License. See the [LICENSE](LICENSE) file for additional info.

---

## 👏 Acknowledgments

> - Personal Project

---

## Credits

This awesome documentation was automatically generated using the [README-AI Project](https://github.com/eli64s/README-AI)

---