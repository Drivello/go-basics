# Proyecto Educativo: CRUD con Fiber y Ent ORM

Este proyecto demuestra cómo construir una API REST básica en Go utilizando el framework web [Fiber](https://gofiber.io/) y el ORM [Ent](https://entgo.io/), centrado en una entidad de usuarios.

---

## 📦 ¿Qué es Ent ORM?

[Ent](https://entgo.io/) es un ORM (Object Relational Mapping) para Go que permite definir tu esquema de base de datos utilizando código Go. Es moderno, fuertemente tipado y enfocado en el rendimiento, la mantenibilidad y la experiencia del desarrollador.

### 🧠 ¿Cuándo conviene usar Ent?

Ent es ideal cuando:
- Querés mantener el esquema de la base de datos versionado y tipado.
- Necesitás relaciones complejas entre entidades.
- Buscás generar consultas y migraciones sin escribir SQL a mano.
- Valorás la integración con herramientas modernas (GraphQL, gRPC, OpenAPI).

---

## ✨ Beneficios de usar Ent

- Generación automática de código y migraciones.
- Código tipado 100% Go.
- Validaciones y hooks integrados.
- Compatible con SQLite, PostgreSQL, MySQL, MariaDB, SQL Server.
- Arquitectura extensible y modular.

---

## 🚀 Instrucciones para usar Ent en Go

### 1. Instalar Ent

```bash
go install entgo.io/ent/cmd/ent@latest
```

### 2. Crear el esquema

Por ejemplo, en `ent/schema/user.go`:

```go
package schema

import "entgo.io/ent/schema/field"

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.String("email").Unique(),
        field.String("username").Unique(),
        field.Int("age"),
    }
}
```

### 3. Generar el código

```bash
go generate ./ent
```

Asegurate de que tu archivo `ent/generate.go` contenga:

```go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

Esto generará automáticamente el cliente, migraciones y el código necesario para operar con la base de datos.

---

## ▶️ Cómo correr este proyecto

```bash
go mod tidy
go run main.go
```

La aplicación quedará corriendo en `http://localhost:3000`.

---

## 📫 Colección de Postman

Incluimos una colección Postman para probar la API de usuarios:

- `GET /users` → Listar todos los usuarios.
- `GET /users/:id` → Obtener un usuario por ID.
- `POST /users` → Crear un nuevo usuario.
- `PUT /users/:id` → Actualizar usuario.
- `DELETE /users/:id` → Eliminar usuario.

> 📎 Archivo incluido: [Ent Users API.postman_collection.json](EntUsersAPI.postman_collection.json)

---

## 🛠️ Requisitos

- Go 1.20+
- SQLite (por defecto) o configurar otra base

---

## 📁 Estructura del proyecto

```
ent-orm-basic/
├── ent/                # Código generado por Ent y esquema
│   └── schema/         # Definición de entidades
├── main.go             # CRUD principal con Fiber
├── go.mod / go.sum     # Dependencias
└── EntUsersAPI.postman_collection.json
```

---

## 📚 Recursos útiles

- [Documentación oficial de Ent](https://entgo.io/docs/)
- [Ejemplos en GitHub](https://github.com/ent/ent)

