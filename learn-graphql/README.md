# Learn GraphQL

## Schema

See in the main file.

## Start

- To run the server
```
air
```

- To use API
```bash
# With cURL
curl --request POST \
  --url http://localhost:8080/query \
  --data '{"query":"{\n  name\n  school {\n    name\n    address\n    students {\n      id\n      name\n      totalCourses\n    }\n  }\n}"}'

# With HTTPie
http POST :8080/query --raw '{"query":"{\n  name\n  school {\n    name\n    address\n    students {\n      id\n      name\n      totalCourses\n    }\n  }\n}"}'
```