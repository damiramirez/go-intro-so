# Sincronización

Cuando hablamos de Sincronización nos estamos refiriendo a poder "coordinar" u "ordenar" dos acciones que pueden ocurrir al mismo tiempo. De esta manera, podemos ordenarlas u organizarlas de manera temporal para que solo una ocurra al mismo tiempo.

Dentro de Go vamos a tener `sync.WaitGroup`, `sync.Mutex` y `chan`.

## sync.WaitGroup

Se utiliza para esperar la finalizacion de un grupo de rutinas. Permite que una goroutine espere hasta que todas las rutinas en el grupo hayan completado su ejecución. Este metodo NO se puede utilizar para sincronizar recursos.

- `Add` -> Incrementa el contador
- `Done` -> Decrementa el contador
- `Wait` -> Bloquea hasta que el contador llegue a cero

```Go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
  wg.Add(1)
  go func(id int) {
    defer wg.Done()
    // Hacer algún trabajo
    // ...
  }(i)
}
// Esperar a que todas las goroutines completen
wg.Wait()
```

## sync.Mutex

Se utiliza para lograr exclusión mutua y evitar condiciones de carrera en secciones criticas del código. Garantiza que solo una rutina pueda acceder a una sección crítica a la vez.

```Go
var mu sync.Mutex

// Sección crítica protegida por el mutex
mu.Lock()
// Acciones en la sección crítica
mu.Unlock()
```

## chan

Los canales se utilizan para la comunicación y sincronización entre rutinas. Permite desde coordinar y sincronizar secciones criticas hasta el envío y la recepción de datos entre rutinas. Los canales tambien pueden ser usados como semáforos o para controlar el accesso a recursos compartidos.

- `make(chan Tipo)` -> creo canal
- `<- ch` -> enviar datos
- `ch <-` -> recibo datos

```Go
ch := make(chan int)

go func() {
  // Hacer algún trabajo
  // ...
  ch <- 42 // Enviar datos al canal
}()

// Recibir datos del canal
valor := <-ch
```

## Resumen

`sync.WaitGroup` se utiliza principalmente para esperar la finalización de un conjunto específico de rutinas antes de continuar la ejecución. Útil cuando el número de rutinas es conocido de antemano.

`sync.Mutex` se utiliza para lograr la exclusión mutua y evitar condiciones de carreras

`chan` se utiliza si se necesita comunicación y sincronización entre rutinas o si se necesita controlar el acceso a un recurso compartido.
