build:
	go build -o sqlx-etude

mysql:
	mysql -h 127.0.0.1 -P 23306 -usqlxetude -psqlxetude sqlxetude

init_database:
	mysql -h 127.0.0.1 -P 23306 -uroot -proot < init.sql
