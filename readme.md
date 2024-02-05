# README.md - API REST con Go

## Descripción

Este proyecto consiste en una API REST desarrollada en Go que utiliza SQLite como base de datos. La aplicación proporciona operaciones CRUD para eventos y usuarios, además de la autenticación mediante JWT y el cifrado de contraseñas mediante bcrypt.

## Configuración

### Base de Datos

La base de datos utilizada es SQLite. Asegúrate de tener SQLite instalado en tu entorno de desarrollo. La conexión a la base de datos se inicializa en el archivo `db.go`. Puedes configurar la conexión modificando la siguiente línea:

```go
DB, err = sql.Open("sqlite3", "api.db")
```

### Dependencias

Asegúrate de tener las dependencias necesarias instaladas. Puedes hacerlo ejecutando:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/mattn/go-sqlite3
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/dgrijalva/jwt-go
```

## Ejecución

Para ejecutar la aplicación, sigue estos pasos:

1. Inicializa le proyecto ejecutando el siguiente comando:

```bash
go run .
```

Esto iniciará el servidor en `localhost:8080`.

## Endpoints

### Obtener todos los eventos

```http
GET /events
```

### Obtener un evento por ID

```http
GET /events/:id
```

### Crear un nuevo evento

```http
POST /events
```

### Actualizar un evento por ID

```http
PUT /events/:id
```

### Eliminar un evento por ID

```http
DELETE /events/:id
```

### Registro de usuario

```http
POST /signup
```

## Estructura del Proyecto

```plaintext
.
├── models
│   ├── event.go
|   └── user.go
├── db
│   └── db.go
├── routes
│   ├── routes.go
|   ├── events.go
|   └── users.go
├── .gitignore
├── README.md
└── main.go
```

- `main.go`: Archivo principal que inicializa la base de datos y el servidor.
- `db`: Paquete que gestiona la conexión a la base de datos y las operaciones relacionadas.
- `routes`: Paquete que define las rutas y controladores de la aplicación.
- `models`: Paquete que define los modelos de datos.
- `README.md`: Archivo de documentación.

## Contribuciones

¡Las contribuciones son bienvenidas! Si encuentras algún problema o tienes sugerencias para mejorar el proyecto, no dudes en enviar un pull request.
