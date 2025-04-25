# TODO CLI

## Overview

This is a simple TODO CLI application, created with GoLang, to learn more about GoLang and to practice clean architecture.

## Features

The application supports the following operations:

### Add TODO

Create a new TODO with a name and an optional due date.

```bash
todo create "<name>"
```

### List all TODOs

Display the information of all created TODOs.

```bash
todo list
```

### Edit TODO

Update the name of a TODO item by ID.

```bash
todo update <id> "<new name>"
```

### Remove TODO

Delete a TODO item by ID.

```bash
todo remove <id>
```

### Complete TODO

Complete a TODO item by ID.

```bash
todo complete <id>
```

## Tools

### Migrate

The GoLang migrate CLI is used for handling database migrations.

The `migrate` command can be installed by running the command:

```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

OR

With the make command from the [Makefile](https://github.com/rykeroc/todo-cli/blob/32af4e5f64ff0d9f18212635017d49a030856451/Makefile) that is present in this repo:

```bash
make install_migrate
```

Refer [here](https://github.com/golang-migrate/migrate) for information about using `migrate`.

### SQLite

Refer [here](https://www.sqlitetutorial.net/what-is-sqlite/) for information about `SQLite` operations.

### Cobra

`cobra` is used to support the CLI behaviour.

Refer [here](https://github.com/spf13/cobra) for information about `cobra`.

### Logrus

`logrus` is used for handling logging across the application.

Refer [here](https://github.com/sirupsen/logrus) for more information about `logrus`.