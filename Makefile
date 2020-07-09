migrate: migrate/databases migrate/tabels

migrate/databases:
	mysql -h 0.0.0.0 -P 13306 --protocol tcp -u root -ppass < databases.sql
migrate/tabels:
	mysql -h 0.0.0.0 -P 13306 --protocol tcp -u root -ppass test_db < tables.sql
mysql:
	mysql -h 0.0.0.0 -P 13306 --protocol tcp -u root -ppass -D test_db
