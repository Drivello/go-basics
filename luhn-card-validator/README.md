# Card Validator with Luhn Algorithm

## Descripción del Proyecto

Este proyecto es un ejemplo **educativo** que demuestra cómo validar un número de tarjeta de crédito utilizando el **algoritmo de Luhn**. También se identifica la marca de la tarjeta (Visa, MasterCard, Amex, Discover, etc.) mediante el análisis de prefijos y longitud.

La validación estructural de tarjetas es un paso común antes de enviar los datos a un **procesador de pagos real**, y es especialmente útil para realizar una **prevalidación rápida del cliente** o en entornos offline.

---

## ¿Qué es el algoritmo de Luhn?

El algoritmo de **Luhn**, también conocido como **"mod 10"**, es un método simple de validación de números de identificación. Fue desarrollado en 1954 por Hans Peter Luhn, un científico de IBM, y se utiliza ampliamente para validar números de tarjetas de crédito, números de identificación gubernamentales, IMEIs, etc.

## ¿Por qué se sigue usando hoy en día?

- Hoy en día está estandarizado por la ISO/IEC 7812 para tarjetas de identificación (como tarjetas de crédito).
- Es suficiente para prevalidar input del usuario antes de enviar a un procesador de pagos.
- Reduce la carga de errores y mejora la UX antes de transacciones reales.

Disclaimer: 
Solo valida que un número es estructuralmente válido, no que esté asociado a una cuenta real. No impide que alguien genere un número que pase la validación.

### ¿Cómo funciona?

1. Empezá desde el último dígito y avanzá hacia la izquierda.
2. Cada segundo dígito se **duplica**. Si el resultado es mayor a 9, se le **resta 9**.
3. Se suman todos los dígitos.
4. Si el total es divisible por 10, el número es válido.

Ejemplo:
```
Número: 4111 1111 1111 1111
Resultado: válido (suma total divisible por 10)
```

---

## Estructura del Proyecto

```
card-validator/
├── main.go                # Ejecución principal del programa con tarjetas de prueba
├── card/
│   ├── validator.go       # Algoritmo de Luhn y sanitización de entrada
│   └── detector.go        # Detección de tipo de tarjeta según prefijo y longitud
```

---

## Comportamiento del Sistema

- Se sanitiza el input (remueve espacios, valida que solo tenga dígitos).
- Se aplica el algoritmo de Luhn para verificar si es un número válido.
- Se analiza el prefijo para determinar la marca de la tarjeta.
- Se muestra en consola el resultado indicando si es válida y de qué tipo es.

---

## Justificación Técnica

- El algoritmo de Luhn se aplica estrictamente sobre el string de números sin espacios.
- La detección de marca se hace usando expresiones regulares y prefijos bien definidos.
- El sistema es modular: la lógica está separada entre validación, detección y entrada principal.
- Se puede extender fácilmente para agregar más tipos de tarjetas.

---

## Cómo Ejecutar

```bash
go run main.go
```

---

## Objetivo de este repositorio

Este repositorio tiene fines educativos. Es útil para aprender:

- Algoritmos clásicos aplicados a validaciones reales
- Modularización de código en Go
- Uso de expresiones regulares y sanitización de input
- Preparación de validaciones previas a gateways de pago

