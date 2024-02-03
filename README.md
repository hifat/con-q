# ConQ

lorem ip sum ........

## Run

```
make run
```

## Development Setup

### Install Go Wire

-  If you are unable to use the `wire` command, you need to add `$GOPATH/bin` to your `$PATH`
   ```shell
   # Fish shell example
   fish_add_path $(go env GOPATH)/bin
   ```

### Install Atlas

[How to install](https://atlasgo.io/guides/orms/gorm#installation)

## Editor Setup

### VS Code

Fix wire.go in di Package Warning

1. Create a `.vscode` directory at the root of the project.
2. Create a `settings.json` file in the .vscode directory.
3. Add the following JSON to `settings.json`

   ```json
   {
      "gopls": {
         "buildFlags": ["-tags=wireinject"]
      },
      "editor.tabSize": 4
   }
   ```
Config `launch.json` for debugging (optional)
```json
"version": "0.2.0",
"configurations": [
   {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd/rest",
      "cwd": "${workspaceFolder}"
   },
]
```

## Developing ConQ

### Add ENV variables

1. Add variable to `config/env/.env`
2. Add variable to `internal/app/config/config.go`

### Dependency Injection

-  Add all injection function in feature that you develop. to `/di/wire.go` (Only that are not already in the set)
-  Use `wire` for Dependency Injection

   ```shell
   wire ./...
   ```

### DB Fixing

-  When you fixed gorm model

   ```shell
   atlas migrate diff --env dev
   ```

-  When you want to migrate
   ```shell
   atlas migrate apply --env dev
   ```

## Deeployment

Create `uuid-ossp` extension if not exists in your postgres

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Generate API docs

```
make swag-init
```
