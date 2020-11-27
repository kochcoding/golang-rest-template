# Template for REST-API's in Golang

This is a simple template I built over the past year. I have tried many different things and this works best for me.
This code is actually a part of a list-app I use for me and my roommates. It has many other features like realtime updates using websockets,
but that is to much for a template.

![alt text](https://github.com/kochcoding/memes/blob/master/api.png)

## Getting started

The code is self-explanatory but has was to less documentation (I'm sorry! If I find some time in future I will add documentation!). The API has basic 
CRUD-functionality and uses a postgres DB in this example, but the interesting part is how I built the services and let handlers handle the routes.

I used a YAML-file to manage the secrets. There is a Config-struct in the types package that describes how the YAML-file must look like.
You have to pass the path to that file as a flag:

```go run app/server.go --config=<path_to_config.yml>```

