# sitoo-test-assignment
Sitoo Test Assignment.

## Usage

Start MySQL container:
```bash
docker-compose up
```

Create empty database:
```bash
mysql -h 127.0.0.1 -P 3306 -u root < create_database.sql
```

Start client:
```bash
mysql -h 127.0.0.1 -P 3306 -u root
```

Add a product:
```
curl "localhost:8080/api/products" \
    -d '{"title": "Awesome socks", "sku": "SCK-4511"}'
```

Get products based on filter:
```
curl "localhost:8080/api/products?sku=SCK-4511&start=0&num=15"
{
    "totalCount": 1,
    "items": [
        {
            "productId": 1,
            "title": "Awesome socks",
            "sku": "SCK-4511",
            "description": null
        }
    ]
}
```

## Third-party dependencies
* `github.com/gin-gonic/gin` -- For exposing the HTTP REST API
* `gorm.io/gorm` -- ORM library for making interacting with the database easier
* `gorm.io/driver/mysql` -- Supporting `gorm` library for MySQL
* `github.com/go-sql-driver/mysql` -- Used for parsing of MySQL error codes
