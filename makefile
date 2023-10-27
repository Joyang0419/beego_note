
mysql:
	cd ./build && docker-compose up -d mysql

app:
	cd ./build && docker-compose up -d app

down:
	cd ./build && docker-compose down