# Concurrent Payment Queue

## Descripción del Proyecto

Este proyecto es un ejemplo **educativo** de cómo modelar una **cola de pagos concurrente** en Go. Se utiliza concurrencia con **goroutines** y **channels**, junto con **contextos** para lograr un cierre limpio del sistema.

El código está organizado siguiendo buenas prácticas de separación de responsabilidades, permitiendo comprender y extender fácilmente el flujo de trabajo de un sistema concurrente. Está pensado como material de estudio para entender patrones simples de concurrencia en Go y cómo manejar workers dedicados a tareas específicas.

---

## Estructura del Proyecto

```
payment-queue/
├── main.go                 # Punto de entrada, arranca workers y simula pagos
├── model/
│   └── payment.go          # Definición del modelo de pago y tipos de canal
├── dispatcher/
│   └── dispatcher.go       # Enrutador que deriva cada pago al canal correcto
└── worker/
    └── worker.go           # Lógica del worker que procesa pagos
```

---

## Comportamiento del Sistema

- Los pagos entran por un canal común y son redirigidos por un **dispatcher** al worker correspondiente según su tipo (`tarjeta`, `transferencia`, `débito`).
- Cada **worker procesa un pago por segundo**.
- El sistema escucha una señal de cierre (`Ctrl+C` o `SIGTERM`) y detiene los workers de forma ordenada con `context.Context`.

---

## Justificación Técnica

- **Go** es ideal para este tipo de problemas gracias a su modelo de concurrencia simple y eficiente.
- Se evita el uso de herramientas externas como `sync.WaitGroup`, a favor de un enfoque claro y comprensible para estudiar el comportamiento asincrónico.
- Cada componente está aislado: `model`, `worker`, `dispatcher`, lo que hace más fácil razonar y testear partes individuales del sistema.

---

## Cómo Ejecutar

```bash
go run main.go
```

El programa procesará pagos de ejemplo automáticamente y se detendrá cuando se cierre el canal de entrada o se reciba una señal del sistema.

---

## Objetivo de este repositorio

Este repositorio tiene fines exclusivamente educativos, como referencia para practicar:

- Diseño concurrente en Go
- Uso de `context.Context` para control de cancelación
- Modelado de canales y workers dedicados
- Cierre limpio de procesos asincrónicos
