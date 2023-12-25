# Building and Running `go-gin-jwt` Application

## Introduction

This document provides a step-by-step guide to building and executing the `go-gin-jwt` application using the Go programming language. The application is based on the Gin framework and implements JSON Web Token (JWT) authentication.

## Prerequisites

Ensure that you have the following prerequisites installed on your system:

- Go programming language (version 1.16 or higher)
- Git (optional but recommended for cloning the repository)

## Steps

### 1. Clone the Repository

If you haven't already, clone the `go-gin-jwt` repository from the source.

```bash
git clone https://github.com/vivekjha1213/go-gin-jwt.git
```

### 2. Navigate to the Project Directory

```bash
cd go-gin-jwt
```

### 3. Build the Application

Run the following command to build the Go application:

```bash
go build
```

This command compiles the Go code in the current directory and generates an executable file.

### 4. Run the Application

Execute the compiled application by running the following command:

```bash
./go-gin-jwt
```

This will start the `go-gin-jwt` application, and you should see output indicating that the server is running.

### 5. Access the Application

Once the application is running, you can access it by opening a web browser and navigating to `http://localhost:PORT`, where `PORT` is the port number specified in the application or the default port if not explicitly set.
