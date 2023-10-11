## Bug conditions

### Bug

```
test2: SQL logic error: row value misused (1)
```

```go
		aDB.Nums = make(StringSlice, 1)
		aDB.Nums = append(aDB.Nums, 10)
```

```go
		aDB.Nums = make(StringSlice, 2)
```

### No bug

```go
		aDB.Nums = make(StringSlice, 1)
		aDB.Nums = append(aDB.Nums, 10)
```

```go
		aDB.Nums = make(StringSlice, 1)
```

## Analysis

https://github.com/go-gorm/gorm/issues/4879

https://github.com/go-gorm/gorm/issues/3585

Both are still open "sqlite3 'row value misused' error when preloading a model with composite primary keys and foreign key relations"

we are using	gorm.io/gorm v1.25.4.

This seems pretty arcane. The strange thing is that the reference model 

I forgot 

```go
func (ss *StringSlice) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), ss)
}

func (ss StringSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(ss)
	return string(b), err
}
```

There was no compilation issue. The bug was runtime