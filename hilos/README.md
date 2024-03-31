# Hilos

Un hilo es una tarea que puede ser ejecutada al mismo tiempo que otra tarea. Esta tarea comparte recursos del proceso padre que lo ejecuta.

Por lo que vamos a utilizar hilos cuando mi programa tenga que hacer mas de una cosa al mismo tiempo.

## Hilos de kernel

Son entidades de kernel y son administrador por el planificador del sistema operativo, por lo que el programador no tiene control directo sobre estos hilos. Un hilo de kernel se ejecuta dentro de un proceso, pero cualquier otro hilo del sistema puede hacer referencia a él.

## Hilos de usuario

Entidad utilizada por los programadores para manejar varios flujos de controles dentro de un programa. Este hilo solo existe dentro de un proceso por lo que no puede hacer referencia a un hilo de otro proceso; El sistema operativo NO sabe que existen y solo reconoce un hilo de ejecución.

# Hilos en Go

Golang como todo lenguaje permite la creación a demanda de Hilos y a estos los llama **"Rutinas"**. Estas rutinas son administradas por el motor de Golang, osea las rutinas que creemos son **Hilos de usuario**. Dentro del motor existe un **Scheduler** que se encarga de planificar y administrar dichas rutinas dentro de hilos de kernel.

Podemos iniciar un hilo usando la palabra `go` delante de la invocacion de una función.
