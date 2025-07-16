<div align="center">
    <img src="/web/static/logos/mums.svg" alt="mums logo" title="mums logo" width="256">
    <h1>mums</h1>
    <h4>
        Serving beverages since 1337!
        <br />
        <a href="https://mums.dsek.se/">mums.dsek.se</a>
    </h4>
</div>

## About


A web application designed to easily track and manage *mums* throughout the introductory period.  

This project is built with simplicity at its core, deliberately avoiding the
build steps that are common in modern web development. It combines
[HTMX](https://htmx.org/) on the frontend with
[Echo](https://echo.labstack.com/) on the backend to create a lightweight,
server-driven architecture. Core features like authentication, session
management, and routing are implemented from scratch, giving full control while
keeping complexity to a minimum. All data is stored in a lightweight
[SQLite](https://www.sqlite.org/) database, making the setup easy to run and
maintain.

## Development Guide

### Prerequisites

Ensure the following tools are installed:

- [Go](https://go.dev/) – The Go Programming Language  
- [Docker](https://www.docker.com/) *(optional)* – For containerized environments

### Setup

1. Clone or fork this repository.

### Running

1. Run `go run cmd/mums/main.go` from the project root,  
   or alternatively run `go tool air` for live reloading.
