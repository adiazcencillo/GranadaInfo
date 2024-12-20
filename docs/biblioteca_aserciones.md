# Decisión Técnica: Biblioteca de aserciones

## Características Objetivas Deseadas

Tal y como se indica en [#23](https://github.com/adiazcencillo/GranadaInfo/issues/23), será importante que la bilioteca elegida cumpla las siguientes consideraciones generales:

- **Soporte y comunidad activa**, asegurando ausencia de deuda técnica en el medio plazo.

En cuanto a requisitos particulares para la biblioteca de aserciones, destacamos lss siguientes:

### **Capacidades avanzadas de aserción**
Queremos contar con una herramienta de aserciones que nos permita escribir tests legibles y detallados. Será importante contar con:
- **Variedad de tipos de aserciones**, como comparaciones, verificaciones de errores y validaciones complejas.
- **Mensajes descriptivos** personalizables que expliquen claramente los fallos.
- Opcionalmente, una herramienta que siga la filosofía **_BDD_** (Behavior-Driven Development) para hacer más legibles y entendibles los tests.

Recalcar que `Go` no maneja de forma nativa aserciones ya que realiza un manejo de errores explícito.

---

## Bibliotecas consideradas


### **Testing**
Es el framework de pruebas en `Go`. Si bien no soporta el uso de aserciones (ya que como se indica más arriba `Go` no soporta estas expresiones), se apoya en los métodos `t.Error`y `t.Fatalf` para generar los tests y reportar los errores. Esto limita la capacidad de por ejemplo personalizar los mensajes de las aserciones, ya que sigue un enfoque distinto.

### **`Testify`**
Biblioteca que proporciona aserciones expresivas como `Equal`, `Contains`, y `Error`. Esto añade una capa de abstracción sobre los métodos `t.Error` y `t.Fatalf` de `Testing`, proporcionando una sintaxis más clara, alineada con las llamadas aserciones funcionales. 

### **`Gomega`**
Biblioteca de aserciones estilo _BDD_ diseñada para integrarse con `Ginkgo`, pero también compatible con `go test`. Su enfoque _BDD_ permite escribir aserciones expresivas cercanas al lenguaje humano.

### **`Goblin`**
Aunque no es únicamente una biblioteca de aserciones (se considera un "_testing framework_") ya que también incluye funcionalidad para organizar las pruebas. He decidido añadirlo por su filosofía _BDD_ en las aserciones.

---

## Conclusión y Elección

Las tres herramientas están bastante actualizadas y cuentan con una comunidad activa. Si bien `Testing` es la biblioteca estándar de `Go`, presenta algunas limitaciones. Cuenta únicamente con los métodos `t.Error`y `t.Fatalf` para generar las aserciones, por lo que la sintaxis de los tests podría complicarse, además de no seguir un enfoque _BDD_.

`Goblin` es una opción interesante porque además es capaz de organizar los tests, pero no parece contar con soporte actualizado.

En cuanto a `Testify` y `Gomega`, he decidido decantarme por `Gomega` debido a su enfoque _BDD_, que nos va a permitir escribir tests legibles.




