# personal-blog

Simple Go CRUD app using **net/http** and **HTML templates**.

## Features

* Create, read, update, delete articles
* Basic templates + static CSS
* Clean and minimal structure

## Run

```bash
go run main.go
```

App runs at **[http://localhost:8080](http://localhost:8080)**

## Routes

* `/` – list articles
* `/new` – create
* `/view/{id}` – view
* `/edit/{id}` – update
* `/delete/{id}` – delete (POST)

---

https://roadmap.sh/projects/personal-blog
