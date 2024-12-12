# Decisión Técnica: Herramientas de Testing

## Características Objetivas Deseadas

Tal y como se indica en [#23](https://github.com/adiazcencillo/GranadaInfo/issues/23), será importante que las herramientas elegidas cumplan las siguientes consideraciones generales:

- **Soporte y comunidad activa**, asegurando ausencia de deuda técnica en el medio plazo.
- **Cumplir con los estándares y mejores prácticas del lenguaje**, fomentando la interoperabilidad y el uso de patrones comunes.

En cuanto a requisitos particulares de herramientas de test, destacamos los siguientes:

### **Capacidades avanzadas de aserción**
Queremos contar con una herramienta de aserciones rica y expresiva, que nos permita escribir tests legibles y detallados. Será importante contar con:
- **Variedad de tipos de aserciones**, como comparaciones, verificaciones de errores y validaciones complejas.
- **Mensajes descriptivos** que expliquen claramente los fallos.
- Opcionalmente, una herramienta que siga la filosofía **_BDD_** (Behavior-Driven Development) para hacer más legibles y entendibles los tests, tanto para desarrolladores como para interesados no técnicos.

### **Herramientas de reporte y visualización**
Al ejecutar nuestros tests, queremos obtener una representación comprensible y concisa de los resultados obtenidos. Esto incluye:
- **Generar informes detallados** que contengan información como el número de tests ejecutados, fallados y su tiempo de ejecución.
- **Interfaz CLI** amigable que facilite la interacción con los resultados.

---

## Elección Basada en los Criterios Objetivos


### **`go test`**
Es el `test_runner` estándar para `Go`. Permite ejecutar tests unitarios con soporte básico para aserciones, subtests y perfiles de cobertura.
- **Ventajas**:
  - Es parte del ecosistema oficial de `Go`, sin necesidad de dependencias externas.
  - Es el _framework_ recomendado por la comunidad.
- **Limitaciones**:
  - Carece de capacidades avanzadas de aserción. Utiliza únicamente los métodos `t.Error`y `t.Fatalf` del paquete `Testing` a modo de "aserciones", por lo que la sintaxis de los test puede volverse compleja.
  - No proporciona herramientas de reporte avanzadas.

### **`Testify`**
Un paquete que complementa `go test` con un marco de aserciones funcionales y funcionalidad de mocking. Utiliza `go test` como _test runner_.
- **Ventajas**:
  - Proporciona aserciones expresivas como `Equal`, `Contains`, y `Error`. Esto añade una capa de abstracción sobre los métodos `t.Error` y `t.Fatalf` de `Testing`, proporcionando una sintaxis más clara y legible.
  - Cuenta con una amplia comunidad activa y mantenimiento regular.
  - Presenta reportes detallados con mensajes de error expresivos.
- **Limitaciones**:
  - Carece de funcionalidad nativa para tests BDD.

### **`Goblin`**
Un _framework_ BDD que introduce una sintaxis parecida a Mocha/Chai en JavaScript. Utiliza `go test` como _test runner_.
- **Ventajas**:
  - Sintaxis amigable para escribir tests orientados a comportamientos.
  - Permite agrupar tests jerárquicamente, lo que mejora la organización.
  - Permite realizar reportes claros e informativos.
- **Limitaciones**:
  - Comunidad más pequeña en comparación con `Testify`. Además, la comunidad parece poco activa (último _commit_ hace tres años).

### **`Ginkgo`**
Un _framework_ BDD completo para `Go`, diseñado para crear tests organizados y legibles.
- **Ventajas**:
  - Se integra fácilmente con `Gomega` para aserciones expresivas.
  - Genera reportes informativos gracias a su _CLI_ específica.
- **Limitaciones**:
  - Estructura compleja.

### **`GoConvey`**
Una herramienta visual para ejecutar y reportar tests en `Go`, orientada a facilitar la legibilidad y comprensión. Utiliza `go test` como _test runner_.
- **Ventajas**:
  - Incluye una interfaz web para visualizar resultados en tiempo real.
  - Incluye aserciones estilo _BDD_ que favorecen la legibilidad de la sintaxis.
- **Limitaciones**:
  - Dependencias adicionales para habilitar la interfaz visual.
  - Comunidad poco activa.

### **`Gomega`**
Un paquete de aserciones estilo _BDD_ diseñado para integrarse con `Ginkgo`, pero también compatible con `go test`.
- **Ventajas**:
  - Proporciona una sintaxis expresiva y flexible para aserciones gracias a su enfoque BDD.
  - Genera reportes informativos con mensajes de error expresivos.
- **Limitaciones**:

---

## Conclusión y Selección

Como se ha discutido, contamos únicamente con dos opciones para el _test runner_. Por un lado podemos decantarnos por la herramienta estándar `go test`, o utilizar `Ginkgo`. En mi caso y para cumplir con las mejores prácticas y estándares del lenguaje, he decidido optar por el uso de `go test` como _test runner_. En cuanto a la biblioteca de aserciones, podría optar también por usar el paquete `Testing`, estándar del lenguaje, pero como se ha indicado en los requisitos, queremos contar con aserciones ricas y expresivas que nos permitan redactar tests legibles, así como representaciones informativas de los resultados. Podemos elegir entre `GoConvey`, `Goblin`, `Gomega` y `Testify`. Debido a la poca actividad aparente de su comunidad, descarto el uso de `GoConvey` y `Goblin`. Las dos opciones restantes para la biblioteca de aserciones son `Testify`y `Gomega`. Ambas bibliotecas, junto con `go test` como _test runner_, cubirirían sobradamente nuestras necesidades. He optado por elegir `Gomega` para contar con aserciones _BDD_ que harán más legibles nuestros tests.




