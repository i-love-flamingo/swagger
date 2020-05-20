# Flamingo Swagger

This module adds routes that serve the swagger ui (https://swagger.io/tools/swagger-ui/) and a generated `swagger.json`

## Installation

Add the module to your Flaming bootstrap:

```go
package main

import "flamingo.me/swagger"

func main() {
	flamingo.App([]dingo.Module{
		...
		new(swagger.Module),
	})
}
```

Make sure that your swagger definition is located under `docs/swagger.json`

After starting your application please open the path "/api-console/". 
E.g.: http://localhost:3322/api-console/

## Generating swagger file

For generating we recommend to use https://github.com/swaggo/swag

```go
swag init --parseDependency=1 --generalInfo=main.go
```

This will parse the defined doc comments and generate a file `docs/swagger.json`

An even better way is to add the following line to your project main file:

```go
//go:generate go run github.com/swaggo/swag/cmd/swag init --parseDependency=1 --generalInfo=main.go
```




