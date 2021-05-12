module github.com/koetsuserizawa/wire_example/blogservice

go 1.14

replace github.com/koetsuserizawa/wire_example/userservice => ../userservice

replace github.com/koetsuserizawa/wire_example => ./

require github.com/go-sql-driver/mysql v1.6.0
