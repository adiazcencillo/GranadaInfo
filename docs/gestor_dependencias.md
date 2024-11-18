# Decisión Técnica: Gestor de Dependencias

## Características Objetivas Deseadas en un Gestor de Dependencias

Para este proyecto, se han identificado las siguientes características clave que debe cumplir el gestor de dependencias, con una explicación de por qué son esenciales:

### 1. **Soporte y Mantenimiento Activo**
   La herramienta debe contar con soporte oficial o una comunidad activa que garantice su evolución y resolución de problemas. Un gestor sin soporte oficial podría ser abandonado, lo que implicaría un cambio  forzado en el futuro, aumentando la deuda técnica y los costos de mantenimiento.

### 2. **Flexibilidad**
   El gestor debe adaptarse a diferentes tamaños de proyectos y permitir tanto configuraciones simples como avanzadas. Una solución flexible asegura que el gestor de dependencias pueda manejar desde proyectos pequeños hasta iniciativas más complejas sin necesidad de cambiar de herramienta.

### 3. **Control de dependencias**    
Será importante contar con una herramienta que realice un control de versiones adecuado, permitiendo la resolución de conflictos entre dependencias y la eliminación de dependencias no usadas. Todo esto nos permitirá garantizar la estabilidad y reproducibilidad del proyecto.

## Elección de `Go Modules` como Gestor de Dependencias

Tras evaluar las necesidades del proyecto, se ha seleccionado **`Go Modules`** como gestor de dependencias por los siguientes motivos:

### 1. **Soporte y Mantenimiento Activo**
`Go Modules` es el gestor oficial recomendado desde la versión `1.13` de Go, lo que garantiza soporte continuo, actualizaciones regulares y compatibilidad con futuras versiones del lenguaje. Su adopción como estándar evita la dependencia de herramientas obsoletas como `Dep` o `Glide`, que ya no reciben mantenimiento. Este respaldo oficial minimiza el riesgo de problemas a largo plazo y asegura la estabilidad del proyecto sin generar deuda técnica innecesaria.

### 2. **Flexibilidad**
`Go Modules` destaca por su capacidad de adaptarse tanto a proyectos pequeños como a iniciativas más complejas. En proyectos sencillos, su configuración es mínima, ya que basta con el archivo `go.mod`, lo que permite un inicio rápido. Para proyectos más avanzados, ofrece características como la gestión de múltiples módulos en un repositorio, control detallado de versiones y escalabilidad para manejar dependencias transitivas complejas. Esta flexibilidad asegura que la herramienta seguirá siendo útil incluso si las necesidades del proyecto crecen en el futuro.

### 3. **Control de Dependencias**
`Go Modules` proporciona un control exhaustivo y confiable para la gestión de dependencias. Su sistema de versionado explícito bloquea versiones específicas, garantizando que las actualizaciones se realicen solo cuando sean deseadas. Además, utiliza la técnica de `Minimum Version Selection` para resolver conflictos automáticamente, asegurando compatibilidad entre dependencias. En casos complejos, permite ajustes manuales en el archivo `go.mod`, facilitados por herramientas como `go mod graph`, y su comando `go mod tidy` mantiene el entorno limpio al eliminar dependencias no utilizadas. Estas capacidades garantizan un entorno de desarrollo estable y reproducible.




