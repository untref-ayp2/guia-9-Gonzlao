package ejercicios

import "sort"

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {
	// Ordenar las actividades por tiempo de finalización
	sort.SliceStable(actividades, func(i, j int) bool {
		return actividades[i].Fin < actividades[j].Fin
	})

	return SelectorRecursivo2(actividades)
}

func SelectorRecursivo2(actividades []Actividad) []Actividad {
	if len(actividades) == 1 {
		return actividades
	}

	primeraActividad := actividades[len(actividades)-1]
	restoActividades := actividades[:len(actividades)-1]

	actividadesSeleccionadas := SelectorRecursivo2(restoActividades)

	if len(actividadesSeleccionadas) == 0 || primeraActividad.Inicio >= actividadesSeleccionadas[len(actividadesSeleccionadas)-1].Fin {
		return append(actividadesSeleccionadas, primeraActividad)
	}

	return actividadesSeleccionadas
}
