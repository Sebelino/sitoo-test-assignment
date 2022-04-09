# sitoo-test-assignment
Sitoo Test Assignment.

Start MySQL container:
```bash
docker-compose up
```

Seed database:
```bash
mysql -h 127.0.0.1 -P 3306 -u root < database.sql
```

Start client:
```bash
mysql -h 127.0.0.1 -P 3306 -u root
```
