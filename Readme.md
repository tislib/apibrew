API Brew [![build](https://github.com/tislib/apibrew/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/tislib/apibrew/actions/workflows/build.yml)  - [website](http://apibrew.tislib.net/)
======================
[![Go Report Card](https://goreportcard.com/badge/github.com/tislib/apibrew)](https://goreportcard.com/report/github.com/tislib/apibrew)
[![Go Reference](https://pkg.go.dev/badge/github.com/tislib/apibrew.svg)](https://pkg.go.dev/github.com/tislib/apibrew)
[![Docker Pulls](https://img.shields.io/docker/pulls/tislib/apibrew)](https://hub.docker.com/r/tislib/apibrew)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/tislib/apibrew)](https://hub.docker.com/r/tislib/apibrew)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/tislib/apibrew)


## Introduction
**Everything is a *Resource***. And **Everything has a *CRUD API***

With Data handler, you can create CRUD APIs for your data in a few minutes

country.yml
```yaml
type: resource
name: country
properties:
  - name: name # name of the property
    type: STRING # type of the property
    length: 255
    required: true
    unique: true
  - name: description # name of the property
    type: STRING # type of the property
    length: 255
```

```bash
dhctl apply -f country.yml
```
So you are ready, you have fully established Rest API for country resource

```javascript
axios.post('http://localhost:9009/country', {
  name: 'Country1',
  description: 'Sample Country 1'
})
```

You can build entire application with resources and references between them (like relations in relational databases)

Everything can be written by resources, not?
**Let's extend our country resource**

```javascript
countryExtension.onCreate(async (country) => {
        country.description = country.description + ' - Extended'

        return country
    })
```
So we have extended our country resource with modification description on create

So, with **API Brew**, you can create your application with resources and you can customize behavior of your resources with extensions

## About
API Brew is a **Low Code software** that allows to create various Grpc and Rest APIs from various database platforms

## Features

* ***Declarative*** - *Data handler* is declarative. You can define your schema in a declarative way, it will create your APIs
* ***Low Code*** - With API Brew, you can create APIs for your data without coding. But you can also extend your APIs with
  extensions, so you can customize behavior of your Resources/Apis
* ***Rest API*** - As you create resources, Rest Apis for them is made automatically
* ***Grpc*** - As you create resources, Grpc Apis for them is made automatically
* ***Database agnostic*** - Data handler is using Postgresql database by default, but it also supports various databases. Including Mongo, Mysql, Redis, etc.
* ***CRUD*** - Crud is on the heart of API Brew.
* ***Swagger*** - Swagger docs are generated automatically
* ***Authentication*** - Data handler supports various authentication methods. Including JWT authentication etc.
* ***Authorization*** - Data handler supports authorization. You can define permissions for your resources
* ***Multi Database*** - You can define multiple databases and do operations on top of them
* ***Scalable*** - Data handler is scalable. You can run it on multiple instances, and it will work as expected, as Data handler does not have any data internally, you can scale it.
* ***Extensible*** - Data handler is extensible. You can extend your resources with extensions. You can also extend your APIs with extensions
* ***CLI support*** - Data handler has a cli tool to manage your resources, dataSources, etc. It is called `dhctl`
* ***Docker*** - Data handler is dockerized. You can run it on docker
* ***Docker Compose*** - Data handler is docker-compose ready. You can run it on docker-compose, see [docker-compose](deploy/docker-compose)
* ***Kubernetes*** - Data handler is kubernetes ready. You can run it on kubernetes, see [kubernetes](deploy/kubernetes)

## Use Cases

* Creating backend for your mobile application or website
* Creating backend for your existing database
* Managing your data in a CRUD fashion
* Creating Standardized, well documented APIs for your data


## Quick Example

### Easy Installation 

```bash
curl -L https://raw.githubusercontent.com/tislib/apibrew/master/deploy/easy-install/run.sh | bash
```

For more detailed installation, see [Installation](https://apibrew.tislib.net/docs/installation)

### Let's create a resource

country.yml

```
type: resource
name: country
properties:
  - name: name
    type: STRING
    length: 255
    required: true
    unique: true
  - name: description
    type: STRING
    length: 255
```

```
dhctl apply -f country.yml
```

Swagger: http://localhost:9009/docs/index.html

```
# Create Country
curl -X POST --location "http://localhost:9009/country" \
    -H "Authorization: <token>" \
    -d "{
   \"name\": \"Country1\",
   \"description\": \"Sample Country 1\"
}"

# List Countries
curl "http://localhost:9009/country" -H "Authorization: <token>"
```

# [Docs](https://apibrew.tislib.net)

- [Tutorials](https://apibrew.tislib.net/docs/tutorials/index) - Tutorials
- [Installation](https://apibrew.tislib.net/docs/installation) - Installation
- [General](https://apibrew.tislib.net/docs/general) - General information about API Brew
- [API](https://apibrew.tislib.net/docs/api) - GRPC and Rest API documentation
- [CLI](https://apibrew.tislib.net/docs/dhctl) - CLI interface(dhctl) documentation
- [Proto](https://apibrew.tislib.net/docs/proto) - Proto documentation
- [SDK](https://apibrew.tislib.net/docs/sdk) - SDK documentation
