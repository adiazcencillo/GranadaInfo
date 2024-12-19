# Decisión Técnica: Bibliotecas de aserciones

## Características Objetivas Deseadas

Tal y como se indica en [#23](https://github.com/adiazcencillo/GranadaInfo/issues/23), será importante que la biliobteca elegida cumpla las siguientes consideraciones generales:

- **Soporte y comunidad activa**, asegurando ausencia de deuda técnica en el medio plazo.
- **Cumplir con los estándares y mejores prácticas del lenguaje**.

En cuanto a requisitos particulares para la biblioteca de aserciones, destacamos las siguientes:

### **Capacidades avanzadas de aserción**
Queremos contar con una herramienta de aserciones que nos permita escribir tests legibles y detallados. Será importante contar con:
- **Variedad de tipos de aserciones**, como comparaciones, verificaciones de errores y validaciones complejas.
- **Mensajes descriptivos** que expliquen claramente los fallos.
- Opcionalmente, una herramienta que siga la filosofía **_BDD_** (Behavior-Driven Development) para hacer más legibles y entendibles los tests, tanto para desarrolladores como para interesados no técnicos.

### **Capacidad para realizar subtests**
Para nuestro cometido será recomendable que la herramienta elegida nos permita generar subtests para poder rastrear mejor los errores en el código.

Recalcar que `Go` no maneja de forma nativa aserciones ya que realiza un manejo de errores explícito.

---

## Elección Basada en los Criterios Objetivos


### **Testing**
Es la biblioteca de aserciones por defecto en `Go`. Se apoya en los métodos `t.Error`y `t.Fatalf` para generar las "aserciones" en los tests. Esto limita la capacidad de por ejemplo personalizar los mensajes de las aserciones, ya que sigue un enfoque distinto.

### **`Testify`**
Biblioteca que proporciona aserciones expresivas como `Equal`, `Contains`, y `Error`. Esto añade una capa de abstracción sobre los métodos `t.Error` y `t.Fatalf` de `Testing`, proporcionando una sintaxis más clara, alineada con las llamadas aserciones funcionales. 

### **`Gomega`**
Biblioteca de aserciones estilo _BDD_ diseñada para integrarse con `Ginkgo`, pero también compatible con `go test`. Su enfoque _BDD_ permite escribir aserciones expresivas cercanas al lenguaje humano.

---

## Conclusión y Selección

Las tres herramientas están bastante actualizadas y cuentan con una comunidad activa. Si bien `Testing` es la biblioteca estándar de `Go`, presenta algunas limitaciones que me hacen descartarla: 

1. Cuenta únicamente con los métodos `t.Error`y `t.Fatalf` para generar las aserciones, por lo que la sintaxis de los tests podría complicarse, además de no seguir un enfoque _BDD_.
2. No permite la creación de subtests.

En cuanto a `Testify` y `Gomega`, he decidido decantarme por `Gomega` debido a su enfoque _BDD_, que nos va a permitir escribir tests legibles.




