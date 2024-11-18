# GranadaInfo

## Problema propuesto

A la hora de planificar un viaje es fundamental contar con información actualizada y accesible de los horarios de los distintos lugares de interés que queremos visitar, con el objetivo de realizar rutas turísticas optimizadas. Sin embargo, en muchas ocasiones nos encontramos con una actualización irregular de estos datos, formatos de presentación inconsistentes en las distintas páginas e incluso fragmentación de la información en distintos sitios web. Contar con esta información en un formato fijo y actualizado ayudaría a poder planificar nuestras visitas en Granada de forma más óptima y satisfactoria.

## Solución propuesta

Desarrollar una aplicación que extraiga los horarios de los distintos monumentos de Granada, los actualice (si fuese necesario) semanalmente e indique aquellos monumentos que cierran en un tiempo próximo.

## Lenguaje de Programación

Se usará como lenguaje de programación principal `Go`.

## Gestor de Tareas
Como gestor de tareas se ha decidido usar la herramienta `Task`. Mediante el archivo `Taskfile.yaml` podemos esepcificar distintas tareas en lenguaje `YAML`. Actualmente podemos verificar la sintaxis del código mediante el comando `task check` y compilar el código mediante `task build`. Para elegir esta herramienta se han considerado una serie de criterios objetivos necesarios para el buen desarrollo del proyecto. Estos criterios pueden consultarse en [Gestor de Tareas](docs/gestor_tareas.md).

## Gestor de Dependencias
Como gestor de dependencias se ha optado por usar `Go modules`, el estándar recomendado por `Go`. De esta forma, en el fichero `go.mod` podremos especificar las dependencias necesarias para el correcto desarrollo del proyecto. Esta herramienta se ha elegido a partir de unos criterior objetivos especificados en [Gestor de Dependencias](docs/gestor_dependencias.md).

## Comprobación de Sintaxis: orden check

Si queremos comprobar la sintaxis del código desarrollado hasta el momento basta con ejecutar el comando `task check`.

## Documentos

[Historias de usuario](docs/user-stories.md)

[Hitos](docs/milestones.md)

[User Journeys](docs/user-journeys.md)

[Gestor de Dependencias](docs/gestor_dependencias.md)

[Gestor de Tareas](docs/gestor_tareas.md)