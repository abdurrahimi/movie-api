# MOVIE API

## How to run ?
### Run in wsl and docker compose
```
  1. run  > docker compose up -d
  2. import db.sql
  3. create rename .env.example to .env
  4. run > go run .
```

## API ENDPOINT

###### 1. Get Movie
```
curl -X GET http://<base-url>/movie
```
```
1. response status code 200 (success)
{
	"message": "Success",
	"movies": [
		{
			"id": 7,
			"title": "John Wick",
			"description": "dalah sebuah film horor Indonesia tahun 2022 yan  Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
			"rating": 7,
			"image": "",
			"created_at": "2022-12-30T04:03:55.368532Z",
			"updated_at": "2022-12-30T04:03:55.368532Z"
		}
	]
}

2. response status code 500 (Internal server error)
{
  "message":"sample error message"
}

```

###### 2. Get Movie By ID
```
curl -X GET http://<base-url>/movie/:id
```
```
1. response status code 200 (success)
{
	"message": "Success",
	"movie": {
			"id": 7,
			"title": "John Wick",
			"description": "dalah sebuah film horor Indonesia tahun 2022 yan  Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
			"rating": 7,
			"image": "",
			"created_at": "2022-12-30T04:03:55.368532Z",
			"updated_at": "2022-12-30T04:03:55.368532Z"
		}
}

2. response status code 400 (Bad request)
{
  "message":"bad request",
  "error":"some error message"
}

3. response status code 500 (Internal server error)
{
  "message":"sample error message",
  "error":"some error message"
}

```

###### 3. Insert movie

```
curl -X POST http://<base-url>/movie/
-d '{
     "title" : "John Wick",
     "description" : "dalah sebuah film horor Indonesia tahun 2022 yan  Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
     "rating" : 7.0,
     "image" : "",
     "created_at" : "2022-08-01 10:56:31",
     "updated_at": "2022-08-13 09:30:23"
  }'
```
```
1. response status code 200 (success)
no content (http status 204)

2. response status code 400 (Bad request)
{
  "message":"bad request",
  "error":"some error message"
}

3. response status code 500 (Internal server error)
{
  "message":"sample error message",
  "error":"some error message"
}

```

###### 4. Update movie

```
curl -X PATCH http://<base-url>/movie/:id
-d '{
     "title" : "John Wick",
     "description" : "dalah sebuah film horor Indonesia tahun 2022 yan  Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
     "rating" : 7.0,
     "image" : "",
     "created_at" : "2022-08-01 10:56:31",
     "updated_at": "2022-08-13 09:30:23"
  }'
```
```
1. response status code 200 (success)
no content (http status 204)

2. response status code 400 (Bad request)
{
  "message":"bad request",
  "error":"some error message"
}

3. response status code 500 (Internal server error)
{
  "message":"sample error message",
  "error":"some error message"
}

```

###### 5. Delete movie

```
curl -X DELETE http://<base-url>/movie/:id
```
```
1. response status code 200 (success)
no content (http status 204)

2. response status code 400 (Bad request)
{
  "message":"bad request",
  "error":"some error message"
}

3. response status code 500 (Internal server error)
{
  "message":"sample error message",
  "error":"some error message"
}

```

