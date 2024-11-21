# Decisión Técnica: Gestor de Tareas

## Características Objetivas Deseadas

Para este proyecto, se han identificado las siguientes características clave que debe cumplir el gestor de tareas, con una explicación de por qué son esenciales:

### 1. **Soporte y Mantenimiento Activo**
   Es fundamental contar con una herramienta respaldada por soporte oficial o una comunidad activa para garantizar su evolución y resolución de problemas. Un gestor sin soporte podría quedar obsoleto, lo que obligaría a realizar un cambio costoso y generaría deuda técnica en el futuro.

### 2. **Ecosistema Go**
   Es clave elegir una herramienta que se integre de manera eficiente con el ecosistema Go, minimizando las dependencias externas. Esto garantiza que el proyecto sea fácil de portar y mantener, evitando complicaciones derivadas de la necesidad de otros lenguajes o plataformas. Además, al reducir las dependencias externas, se disminuye la superficie de ataque, lo que puede mejorar la seguridad del proyecto al evitar vulnerabilidades asociadas con herramientas o bibliotecas de otros lenguajes.


## Elección Basada en los Criterios Objetivos

### 1. Soporte y Mantenimiento Activo
   `Task`, `Bazel`, `Just` y `Ninja` cuentan con una comunidad activa y actualizaciones regulares, lo que garantiza su mantenimiento a largo plazo y la resolución de problemas de manera eficiente. Otras herramientas como `Rake` reciben soporte en la actualidad pero no poseen una comunidad tan activa. Otras como `Make`, `Mage` y `Realize` llevan desde el año pasado sin registrar actividad. Esto puede suponer un problema a largo plazo a nivel de deuda técnica. 

### 2. Ecosistema Go
- `Task`, `Mage` y `Realize` están escritas en `Go`, lo que las convierte en opciones atractivas para proyectos `Go`. Ambas herramientas se integran bien en el ecosistema y aprovechan sus características.
  - `Mage` utiliza sintaxis Go para definir las tareas, lo que permite escribir las tareas en el mismo lenguaje que el proyecto, manteniendo coherencia y facilitando el aprovechamiento de características de `Go`. 

  - `Task`, por otro lado, utiliza YAML para definir las tareas, lo que proporciona una configuración más declarativa y legible, aunque requiere conocer el formato YAML.

  - `Realize` se vende como el #1 task runner en `Go`. Al igual que `Task`, utiliza sintaxis `YAML` para definir sus tareas.

En cuanto a las demás herramientas:

- `Rake` está escrita en Ruby, lo que significa que se necesita tener Ruby instalado y gestionar dependencias mediante `gems` para ejecutar tareas en proyectos Go. Esto puede añadir una capa adicional de complejidad y dependencias fuera del ecosistema Go.
  
- `Just`, `Bazel`, `Make` y `Ninja` están escritos en lenguajes distintos a Go. Sin embargo, todas requieren solo la instalación de la herramienta en el sistema, sin necesidad de añadir dependencias adicionales en el proyecto `Go` mismo.

Una vez descritos los criterios y como se adaptan las distintas herramientas consideradas, obtenemos las siguientes conclusiones:

- `Mage` y `Realize` son gestores interesantes al estar escritos en `Go` y utilizar una sintaxis sencilla. Sin embargo, la falta de actividad en ambos casos puede suponer un problema de deuda técnica futura, por lo que decidimos descartarlos.

- `Make` y `Ninja` son herramientas potentes que no requieren de dependencias externas y cuya comunidad es activa (más en `Ninja` que en `Make`). Sin embargo, la sintaxis que estas herramientas utilizan me hace descartarlas, ya que la curva de aprendizaje puede ser más pronunciada al principio.

- `Rake` requiere dependencias externas y su uso se recomienda en proyectos de `Ruby`, por lo que también la descartamos como opción.

- Por último, `Bazel`, `Just` y `Task` son herramientas que cumplen los requisitos impuestos anteriormente. Sin embargo, decido decantarme por `Task` ya que la sintaxis `YAML` me resulta atractiva, y el estar escrito en `Go` aporta coherencia al proyecto.




