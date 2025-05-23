# Debt Scheduler – Generador de Pagos Mensuales

## Descripción del Proyecto

Este proyecto es un ejemplo **educativo y reusable** que permite dividir un monto total de deuda en **pagos mensuales iguales**, considerando:

- Redondeo correcto a centavos
- Compensación para que la suma sea exacta
- Fechas mensuales futuras bien generadas, independientemente del número de días de cada mes

Es una solución simple y profesional para agendar pagos, ideal para sistemas de préstamos, financiamiento o recordatorios automáticos.

---

## ¿Qué problema resuelve?

Cuando se divide una deuda (por ejemplo, $1000) en partes iguales con redondeo a centavos, puede que los importes no sumen exactamente el total. Este sistema:

- Calcula las cuotas con dos decimales
- Distribuye centavos sobrantes a las primeras cuotas
- Genera correctamente la fecha de vencimiento mes a mes

---

## Estructura del Proyecto

```
debt-scheduler/
├── main.go                # Programa de ejemplo que genera un cronograma de pagos
└── schedule/
    ├── model.go           # Estructuras de datos: Payment y ScheduleRequest
    └── generator.go       # Lógica principal de distribución de deuda y fechas
```

---

## Comportamiento del Sistema

Dado:
- Un monto total (ej: $1000)
- Una cantidad de cuotas (ej: 3)
- Una fecha de inicio (ej: 22/05/2025)

El sistema devuelve:
```
Cuota 1 | Fecha: 22/05/2025 | Importe: $333.34
Cuota 2 | Fecha: 22/06/2025 | Importe: $333.33
Cuota 3 | Fecha: 22/07/2025 | Importe: $333.33
```

✅ Suma total = $1000.00  
✅ Fechas correctas, sin errores por febrero o meses de 31 días

---

## Justificación Técnica

- Se usa `math.Floor()` para asegurar redondeo estable a 2 decimales
- El remanente se distribuye en las primeras cuotas
- `time.AddDate(0, i, 0)` garantiza que se suman **meses calendario reales**, sin errores

---

## Cómo Ejecutar

```bash
go run main.go
```

---

## Objetivo de este repositorio

Este código está pensado como módulo educativo y base de desarrollo para:

- Generación de planes de pago automáticos
- Finanzas personales o corporativas
- Aplicaciones con planes de cuotas
