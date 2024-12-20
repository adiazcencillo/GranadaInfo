# Decisión Técnica: Test runner

## Características Objetivas Deseadas

Tal y como se indica en [#23](https://github.com/adiazcencillo/GranadaInfo/issues/23), será importante que el _test runner_ elegido cumpla las siguientes consideraciones generales:

- **Soporte y comunidad activa**, asegurando ausencia de deuda técnica en el medio plazo.
- **Cumplir con los estándares y mejores prácticas del lenguaje**.

En cuanto a requisitos particulares para el _test runner_, destacamos los siguientes:

### **Capacidad para realizar subtests**
Para nuestro cometido será recomendable que el _test runner_ elegido nos permita generar subtests para poder rastrear mejor los errores en el código.

### **Limitar el número de dependencias**
Será recomendable contar con un _test runner_ que no aumente el número de dependencias de nuestro proyecto.

---

## _Test runner_ considerados

### **`go test`**
Es el _test runner_ estándar de `Go`. Se encarga de ejecutar las pruebas definidas en archivos `_test.go` e informar de los resultados de dichas pruebas. Permite la realización de subtests y cuenta con una CLI que permite la ejecución de pruebas especifícas o el uso de _flags_ en la salida.

### **`Ginkgo`**
Es otro _test runner_ compatible con `Go`. Está enfocado en la creación de tests siguiendo una filosofía _BDD_, usándose normalmente junto a la biblioteca de aserciones `gomega`. Al igual que `go test`, permite la creación de subtests, y cuenta con una CLI personalizada que permite entre otras cosas obtener salidas detalladas de las pruebas. Es un _test runner_ orientado a tareas complejas


---

## Conclusión y Elección

Ambas herramientas son dos buenas opciones para nuestro proyecto y cumplen en general con los requisitos especificados. La elección de `ginkgo` supondría cargar con más dependencias al proyecto (recordemos que vamos a usar `gomega` para las aserciones), mientras que optar por `go test` nos aseguraría usar la herramienta estándar del lenguaje. Por este motivo he decidido usar `go test`.