package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type integer int

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {

	// vamos aprender a utlizar interfaces funciones variaticas
	PrintList("hola", 13, true, 3.14, "juan")
	fmt.Println("###############################")

	// vamos aprender a utlizar any en funciones variaticas
	PrintListAny(13, 3.14, false, "hola", "daniel")
	fmt.Println("###############################")

	//vamos a crear uba funcion que sume y explicar los parametros de tipos y restricciones
	total := sum(12, 35, 12, 20, 30)
	fmt.Printf("la suma de los numeros es: %d \n", total)
	fmt.Println("###############################")

	totalRestriccion := sumTypeRestriction(3.13, 14.5, 14, 20.5)
	fmt.Printf("la suma de los numeros en la funcion sumTypeRestriction es: %f \n", totalRestriccion)
	fmt.Println("###############################")

	//vamos realizar la sumas de 2 datos con tipo de datos creado por uno mismo ocupando la misma funcion sumTypeRestriction
	var num1 integer = 100
	var num2 integer = 300
	total1 := sumTypeRestriction(num1, num2)
	fmt.Printf("la suma de los numeros de tipo de datos creado es: %d \n", total1)
	fmt.Println("###############################")

	//vamos realizar la sumas de 2 datos con tipo de datos creado por uno mismo ocupando la funcion cin interface
	var num3 float64 = 100.39
	var num4 float64 = 300.49
	total2 := sumTypeRestrictionInteface(num3, num4)
	fmt.Printf("la suma de los numeros de tipo de datos creado y utlizando la funcion con interface es: %f \n", total2)
	fmt.Println("###############################")

	//vamos realizar la sumas de 2 datos con tipo de datos creado por uno mismo ocupando golang.org/x/exp/constraints
	var num5 float64 = 89.39
	var num6 float64 = 123.49
	total3 := sumTypeRestrictionConstraints(num5, num6)
	fmt.Printf("la suma de los numeros de tipo de datos creado y utlizando la funcion con Constraints es: %f \n", total3)
	fmt.Println("###############################")

	// vamos a revisar como trabajar con las restricciones y operadores.
	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}
	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 3))
	fmt.Println(Includes(numbers, 8))
	fmt.Println("###############################")

	// vamos a pasar las mismas listas anteriores y donde a la funcion le pasaremos una lista y una funcion.
	fmt.Println(Filter(numbers, func(value int) bool { return value > 2 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))
	fmt.Println("###############################")

	//vamos a revisar y crear una estructura generica.
	product1 := Product[uint]{1, "Zapatos", 50}
	fmt.Println(product1)
	product2 := Product[string]{"dfghjkluytrdoknbgujngjcdsxcv", "Polerones", 200}
	fmt.Println(product2)

}

func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}
func PrintListAny(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// misma funcion que la "Sum" pero la funcion sumTypeRestriction es para trabajar con numero entero y decimales
func sumTypeRestriction[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// misma funcion que la "Sum" pero la funcion sumTypeRestriction es para trabajar con numero entero y decimales
func sumTypeRestrictionInteface[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// aca vamos ocupar la misma función aplicando el paquete constraints
func sumTypeRestrictionConstraints[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// funcion esta ocupando el comparable donde se pueden ocupar los operadores ( == y != )
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// funcion donde vamos a poder incluir otros operadores.
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}
