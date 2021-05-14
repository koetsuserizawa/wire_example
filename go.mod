module github.com/koetsuserizawa/wire_example

go 1.14

replace github.com/koetsuserizawa/wire_example/blogservice => ./blogservice

replace github.com/koetsuserizawa/wire_example/userservice => ./userservice

require github.com/go-sql-driver/mysql v1.6.0
