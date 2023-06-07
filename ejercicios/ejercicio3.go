package ejercicios

import (
	"fmt"
	"sort"
)

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	fmt.Println(objetos)
	sort.Slice(objetos, func(i, j int) bool {
		return objetos[i].Valor/objetos[i].Peso > objetos[j].Valor/objetos[j].Peso
	})
	fmt.Println(objetos)

	resultado := make(map[Objeto]float32)
	pesoActual := 0

	for _, objeto := range objetos {
		if pesoActual+objeto.Peso <= capacidad {
			resultado[objeto] = 1
			pesoActual += objeto.Peso
		} else {
			porcion := float32(capacidad-pesoActual) / float32(objeto.Peso)
			resultado[objeto] = porcion
			break
		}
	}

	return resultado
}
