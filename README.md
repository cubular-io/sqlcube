# sqlcube
A simple golang tool, that reduces the size of sqlc code generation by replacing Structs that are equal in Field.Names and Field.Types with Type Alias

## Install

```bash
go install github.com/cubular-io/sqlcube@latest
```

## Use

sqlcube look for a file sqlcube.yaml with source and target folder.
If not available, it will use the sqlc.yaml source to get source and target.
Only sqlc.yaml version 2 is supported

```bash 
sqlcube
```

Using the --debug flag will print out all changes done.


## Example Yamls

sqlc.yaml
```
version: "2"
sql:
  - engine: "mysql"
    queries: "./db/query"
    schema: "./db/migration/" #schema in this folder
    gen:
      go:
        emit_enum_valid_method: true
        package: "example/models"
        out: "example" 
       # sql_package: "mysql" #go
```

sqlcube.yaml

```yaml
version: "1"
go:
  source: "sqlc"
  target: "result"
```