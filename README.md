# About this project 

Its a simple RESTful API that can encode and decode BASE64, I made it to get familiar with Go syntax and Gin framework. 

# How to run it

In the shell, just write

```
go run server.go
```

Remember to create the ```config.env``` file, otherwise it will throw an error. I included a sample of each enviroment file in the project. 

# Endpoints 

## Encode string to Base64

**Method** 

```
POST
```

**URL**
```
/api/v1/encode/base64
```

**Body**
```json 
{
    "value": "string value to be encoded"
}
```

## Decode base64 to String

**Method** 

```
POST
```

**URL**
```
/api/v1/decode/base64
```

**Body**
```json 
{
    "value": "base64 value to be decoded"
}
```

# Libraries used 

* [Gin Gonic](https://github.com/gin-gonic/gin)
* [GoDotEnv](https://github.com/joho/godotenv)