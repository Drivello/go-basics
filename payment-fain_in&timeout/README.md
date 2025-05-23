# Fan-In Payment Sources (Select + Context Timeout)

## Descripción del Proyecto

Este proyecto es un ejemplo **educativo** que demuestra cómo implementar el patrón **fan-in** en Go utilizando múltiples fuentes de datos que compiten por enviar un resultado primero. Para esto se usa la instrucción `select` y control de tiempo con `context.WithTimeout`.

Cada fuente de pago simula un delay aleatorio, y el sistema captura **la primera respuesta disponible**. Si ninguna responde a tiempo, se cancela todo el flujo de forma ordenada gracias al contexto.

---

## ¿Qué es el patrón Fan-In?

El patrón **Fan-In** proviene del mundo de los sistemas concurrentes y se refiere a la **convergencia de múltiples flujos de entrada** hacia un único canal de salida. Es un patrón útil cuando tenés varias fuentes de datos (o procesos) que compiten por entregar una respuesta, y vos solo necesitás procesar **el primer resultado disponible**.

En Go, el patrón fan-in se implementa naturalmente con:

- **Múltiples goroutines** escribiendo en el mismo canal.
- **`select`** para leer del canal apenas haya un dato disponible.
- Opcionalmente, `context.Context` para cortar procesos que tarden demasiado o se vuelvan innecesarios.

Este patrón es común en:
- Sistemas de red que consultan varios servidores en paralelo y usan el que responde primero.
- Búsquedas paralelas donde se quiere el primer resultado válido.
- Agregación de datos con respuesta temprana.

---

## Estructura del Proyecto

```
fanin-payment/
└── main.go                # Código principal que lanza fuentes y escucha el primer pago válido
```

---

## Comportamiento del Sistema

- Se lanzan múltiples fuentes de pago (ej: Visa, Mastercard, etc) como goroutines.
- Cada fuente simula una demora aleatoria (hasta 1.5 segundos).
- Se usa `select` para capturar **la primera fuente que responda**.
- Si no se recibe respuesta en 1 segundo, el `context.WithTimeout` cancela todo.

---

## Justificación Técnica

- Se emplea `context.WithTimeout` para evitar espera infinita.
- Se prioriza **la respuesta más rápida** con `select`, un patrón ideal para fan-in.
- `rand.Seed(time.Now().UnixNano())` asegura que las simulaciones de delay sean diferentes en cada ejecución.
- El canal `output` es compartido por todas las fuentes: este patrón representa el concepto de fan-in con claridad.

---

## Cómo Ejecutar

```bash
go run main.go
```

La salida mostrará el primer pago recibido o un mensaje de timeout si ninguna fuente responde dentro del tiempo límite.

---

## Objetivo de este repositorio

Este ejemplo tiene fines educativos. Es útil para aprender:

- Concurrencia básica en Go
- Patrón fan-in con múltiples goroutines
- Uso práctico de `select` y `context.Context` para tiempo límite y cancelación de procesos
