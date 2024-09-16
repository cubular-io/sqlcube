# sqlcube
A simple golang tool, that reduces the size of sqlc code generation by replacing Structs that are equal in Field.Names and Field.Types with Type Alias

## Install

```bash
go install github.com/cubular-io/sqlcube@latest
```

## Use

sqlcube reduce look for a file sqlcube.yaml with source and target folder.

```bash 
sqlcube reduce
```

sqlcube generate looks for views,procedures and schema folder, deletes the target folder, and recreates it and copys
schema folder, than sums up the views and add it to the target folder as a file called x_views.sql, does the same for
z_procedures.

```bash
sqlcube generate
```


## Example Yaml
sqlcube.yaml

```yaml
version: "1"
go:
  source: "example/sqlc"
  target: "example/result"
generation:
  views: "example_schema/views"
  procedures: "example_schema/procedures"
  schema: "example_schema/migrations"
  target: "example_schema/sqlc"
```