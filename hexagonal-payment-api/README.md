# Hexagonal Payment API

## Descripción del Proyecto

Este proyecto es una API de pagos desarrollada en Go, estructurada siguiendo el patrón de arquitectura **hexagonal (Ports and Adapters)**. 
La finalidad es mantener un fuerte desacoplamiento entre el núcleo del negocio (dominio y casos de uso) y las tecnologías externas como frameworks web, bases de datos o bibliotecas específicas.

Se eligió esta arquitectura para mejorar la **testabilidad, mantenibilidad y escalabilidad** del sistema. La separación clara de responsabilidades facilita también futuras migraciones tecnológicas o cambios de implementación sin afectar el núcleo del dominio.

Si bien esta implementación es pequeña, es suficiente para demostrar las capacidades de la arquitectura en cuestión.

---

## Estructura del Proyecto

```
hexagonal-payment-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── adapter/
│   │   ├── handler/           # Adaptador de entrada: implementación HTTP de la capa de entrada
│   │   └── repository/        # Adaptador de salida: repositorio en memoria
│   ├── domain/                # Modelos y lógica de dominio
│   ├── port/                  # Puertos: interfaces que definen las entradas y salidas del dominio
│   └── usecase/               # Casos de uso: lógica de aplicación
├── pkg/
│   └── middleware/            # Middleware reutilizable (por ejemplo, logger)
├── go.mod / go.sum            # Dependencias y módulo de Go
```

---

## Detalle de Carpetas

- **`cmd/`**: Punto de entrada de la aplicación. Aquí se instancia el servidor y se conectan los adaptadores al núcleo.
- **`internal/domain/`**: Define los modelos y entidades del negocio. Esta capa no depende de ninguna tecnología externa.
- **`internal/port/`**: Interfaces (puertos) que el dominio expone o requiere para operar. Define contratos, no implementaciones.
- **`internal/usecase/`**: Contiene la lógica de aplicación, implementando los casos de uso del dominio.
- **`internal/adapter/handler/`**: Implementación de la entrada de la aplicación, en este caso HTTP.
- **`internal/adapter/repository/`**: Implementación concreta del almacenamiento, aquí en memoria.
- **`pkg/middleware/`**: Código reutilizable y transversal, como logging o autenticación.

---

## Justificación Técnica

- **Go**, por su rendimiento y concurrencia nativa, es y será siempre mi lenguaje predilecto para desarrollar y enseñar a mis pares.
- **Arquitectura hexagonal** permite desarrollar y testear el dominio sin preocuparse por detalles de infraestructura.
- Separar `adapter`, `port` y `usecase` permite mockear cualquier dependencia fácilmente y facilita el testing unitario y de integración.

---

## Cómo Ejecutar

```bash
go run ./cmd/main.go
```

---

## Objetivo de este repositorio

- Este repositorio busca representar, sin abrumar, las caracteristicas principales de la arquitectura hexagonal. 
