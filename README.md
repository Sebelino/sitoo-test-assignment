# sitoo-test-assignment
Sitoo Test Assignment.

## Usage

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
