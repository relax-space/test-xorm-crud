# test-xorm-crud

## start

``` shell
$ docker-compose up -d
$ go test
```

## result
当xorm根据id执行update操作时，如果id在数据库中不存在，xorm不会返回错误，只会返回受影响的行为0
当xorm更新的的字段有唯一索引的时候，如果发生唯一索引冲突的时候，xorm会返回错误