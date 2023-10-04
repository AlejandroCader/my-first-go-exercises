package main

import (
	"fmt"
	"sync"
)

// Definimos una estructura Task para representar una tarea
type Task struct {
	ID     int
	Nombre string
}

// Función para buscar una tarea en la lista de tareas
func buscarTarea(tareas []Task, nombre string, i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	if tareas[i].Nombre == nombre {
		c <- tareas[i].ID // Enviamos el ID de la tarea al canal si se encuentra
		return
	}
	c <- -1 // Enviamos -1 si la tarea no se encuentra
}

func main() {
	// Creamos una lista de tareas
	tareas := []Task{
		{1, "Hacer compras"},
		{2, "Lavar el coche"},
		{3, "Estudiar para el taller de SO"},
		{4, "Pasear al perro"},
	}

	// Creamos un canal para recibir resultados
	resultadoChan := make(chan int)

	// Creamos un grupo WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup

	// Definimos la tarea que queremos buscar
	tareaABuscar := "Pasear al perro"

	// Iniciamos una goroutine para buscar la tarea en paralelo
	for i := range tareas {
		wg.Add(1)
		go buscarTarea(tareas, tareaABuscar, i, &wg, resultadoChan)
	}

	// Usamos una goroutine adicional para cerrar el canal después de que todas las búsquedas hayan terminado
	go func() {
		wg.Wait()
		close(resultadoChan)
	}()

	// Recibimos resultados de las goroutines y buscamos la tarea
	for id := range resultadoChan {
		if id != -1 {
			fmt.Printf("Tarea encontrada: %s (ID: %d)\n", tareaABuscar, id)
			break
		}
	}
}
