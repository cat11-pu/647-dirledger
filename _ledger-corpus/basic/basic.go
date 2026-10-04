package basic

//gofmts:json
const JsonCompact = `{"a":  1, "b":[1,2,3]}`

//gofmts:json
const JsonPretty = `{
  "a": 1,
  "b": [
    1,
    2,
    3
  ]
}`

//gofmts:sql
const SqlLower = `select  id , name from users where id = 1`

//gofmts:sql
const SqlCanon = `SELECT id, name FROM users WHERE id = 1`

//gofmts:go
const Expr = "1 +  2"

//gofmts:go
const ExprCanon = "1 + 2"
