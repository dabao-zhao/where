## where

gorm where 条件辅助生成

<div align=center>

[![Release](https://img.shields.io/github/v/release/dabao-zhao/where.svg)](https://github.com/dabao-zhao/where)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/dabao-zhao/where/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/dabao-zhao/where)](https://goreportcard.com/report/github.com/dabao-zhao/where)

</div>

### 安装

```
go get -u github.com/dabao-zhao/where
```

### 使用 

```
cond := []where.Expr{
    where.Eq{"status" : 1}
}

query, args := where.ToQueryAndArgs(cond)
db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
db.Table("users").Where(query, args...).Rows()
```

