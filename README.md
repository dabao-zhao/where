## where

gorm where 条件辅助生成

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

