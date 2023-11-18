# MYSQL_URL(mysql://user:password@tcp(host:port)/dbname
mysql_url = mysql://root:root@tcp(localhost:3306)/dev

mysql:
	cd ./build && docker-compose up -d mysql

app:
	cd ./build && docker-compose up -d app

down:
	cd ./build && docker-compose down

dbup:
	  migrate -database "${mysql_url}" -path build/mysql/migrations up

dbdown:
	migrate -database "${mysql_url}" -path build/mysql/migrations down

dbscript:
	@read -r -p "filename: " filename \
	&& migrate create -ext sql -dir build/mysql/migrations -seq "$${filename}"
