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

## Editor Setup

### VS Code

-  Fix wire.go in di Package Warning

   1. Create a `.vscode` directory at the root of the project.
   2. Create a `settings.json` file in the .vscode directory.
   3. Add the following JSON to `settings.json`

      ```json
      {
         "gopls": {
            "buildFlags": ["-tags=wireinject"]
         }
      }
      ```

## Developing ConQ

-  Use `wire` for Dependency Injection

   ```shell
   wire ./...
   ```
