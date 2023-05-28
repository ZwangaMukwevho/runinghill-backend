# runinghill-backend
Backend api that connects for GCP database

# Documentation
To view endpoints and payloads for each endpoint:
- Navigate to docs directory.
- Copy the contents on swagger.json or swagger.yml
- Visit [Swagger](https://editor.swagger.io/) and past the contents
- This will show endpoints for the api

# Running
- install go modules. ```go mod tidy```
- Build golang project ```go build .```
- Run api on root directory ```go run main.go```

# Database
- service-file.json has service configurations for the api.
- To use different Database replace the `service-file.json` with configs of other database and change database url in the `firebase_db_connection.go` file