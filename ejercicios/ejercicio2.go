package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	for i := 1; i < len(tareas); i++ {
		key := tareas[i]
		j := i - 1

		for j >= 0 && tareas[j].Tiempo > key.Tiempo {
			tareas[j+1] = tareas[j]
			j = j - 1
		}
		tareas[j+1] = key

	}
}
