# Simple Todo App 

A simple Todo application built using Wails.io

![alt text](image-1.png)

## Features
- Add, edit, and delete todos
- Mark todos as completed
- Save tasks in SQLite database
- UI with a modern design

## Backend
The backend is written in Go and uses SQLite for data storage. It includes the following components:

- database/: Contains the SQLite database file.
- internal/: Contains internal packages for configuration, error handling, handlers, repository, and services.
- models/: Contains the data models.

## Frontend
The frontend is built with Svelte and SvelteUI library. It includes the following components:

src/: Contains the Svelte components and other frontend assets.


## Prerequisites
- Go
- Node.js
- Wails CLI

## Project Structure
```
├── frontend/
│            ├── src/
│            │   ├── App.svelte
│            │   ├── assets/
│            │   │   ├── fonts/
│            │   │   │   ├── OFL.txt
│            │   │   │   └── nunito-v16-latin-regular.woff2
│            │   ├── main.js
│            │   ├── style.css
│            │   ├── taskService.js
│            │   └── vite-env.d.ts       
│                      ....
├── backend/
│            ├── app.go
│            ├── database/
│            │   └── tasks.db
│            ├── internal/
│            │   ├── config/
│            │   │   └── logger.go
│            │   ├── errors/
│            │   │   └── error.go
│            │   ├── handlers/
│            │   │   └── task_handler.go
│            │   ├── repository/
│            │   │   └── task_repository.go
│            │   └── services/
│            │       └── task_service.go
│            └── models/
│                └── task.go

```
## Installation

### Clone the Repository
```sh
git clone https://github.com/ilich956/dmark_test_task.git
cd dmark_test_task
```
### Install the frontend dependencies
```sh
cd frontend
npm install
```

### Run the Application
In project directory
```sh
wails dev
```

### Build for Production
In project directory
```sh
wails build
```
