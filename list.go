//el paquete main tiene un ejemplo de listas
package main

//importar los paquetes fmt y container list
import (
	"container/list"
	"fmt"
)

//metodo main
func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
