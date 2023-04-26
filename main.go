package main

import (
    "fmt"
    "github.com/google/uuid"
)
var (
    raza = "Saiyayin"
    tecnica = "Rayos"
)

// func suma(a1 int, a2 int) (resultado int){
//     resultado = a1 + a2
//     return
// }
func suma(a1 int, a2 int) int{
    resultado := a1 + a2
    return resultado
}
// Funciones variadic 
func sumaV(numeros ...int) int{
    resultados := 0
    for _, c := range numeros{
        resultados += c
    }
    fmt.Println(resultados)
    return resultados
}

func multiplicar(nu int) (int, int, int){
    n1 := nu * 10
    n2 := nu * 20
    n3 := nu * 30
    return n1, n2, n3
}

func main() {
    fmt.Println(uuid.New().String())
    var numero int
    numero = 25
    var nombre string
    nombre = "Facu"
    fmt.Println(nombre, numero)
    numero2, nombre2 := 1, "Ile"
    fmt.Println(numero2, nombre2)
    fmt.Println(raza)
    for j:=0; j < 10; j++{
        fmt.Println(j)
    }
    switch numero {
    case 7:
        fmt.Println(7)
    case 25:
        fmt.Println(25)
    default:
        fmt.Println("Inserte un numero")
    }

    //Arrays
    // var x [3]int => array [0 0 0]
    // var k [3][2]float64 => [[0 0][0 0 ][0 0]]
    //asignar un valor x[1] = 25 => [0 25 0]
    //creo con valores y:= [5]int{1,2,3,4,5} => [1 2 3 4 5]
    //creo con valores a determinar
    k := [...]int{1,2,3,4,}
    fmt.Println(k)
    fmt.Println(k[0]) // Imprimo el primer elemento de un array
    fmt.Println(k[len(k)-1]) //Imprimo el ultimo elemento

    //Slice son arrays dinamicos, no esta definido el tamanho
    // Recordar que el slice no se define el tamanho al crerlo
    y := make([]int, 5, 10)
    fmt.Println(y)
    fmt.Println(len(y))
    fmt.Println(cap(y))
    
    // Slice a partir de un array
    var e, t []int
    e = k[2:4]
    //Si modificamos el slice se modifica el array
    fmt.Println(e, t, k)
    e[0] = 10
    fmt.Println(e, k)
    e = append(e, 100) // El append se usa para los slice, dado que devuelve un nuevo slice 
    fmt.Println(k, e)
    for _,i := range e{ //Siempre son dos elementos que se traen _ completar el argumento mas no lo inicializa
        fmt.Println(i)
    }

    // Diccionarios
    var diccionario = map[string]int{"Mark": 10, "Sandy": 20}
    dicct := make(map[string]string, 7) // para hacer mas rapido el map se le puede definir el tamanho de entrada
    dicct["Facu"] = "Silvestre"  
    diccionario["Mark"]++
    fmt.Println(diccionario)
    fmt.Println(dicct)
    if diccionario["Mark"] == 10{
        fmt.Println("No se sumo")
    }else if  diccionario["Mark"] == 11 {
        fmt.Println(" Se sumo 1")
    } else {
        fmt.Println("Error")
    }
    fmt.Println(suma(1, 3))
    valores := []int{1, 3, 2, 3, 1}
    fmt.Println(sumaV(valores...))
    fmt.Println(multiplicar(10))

    //Punteros
    // & Lugar en la memoria donde se guarda la info
    // * Valor de la direccion del puntero
    // Sirven para indentificar dentro de una funcion una variable global al apunter con punteros

    // album represents data about a record album.
    type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float64 `json:"price"`
    }
    // Metodo de una structura
    // Para que sea un metodo recibe antes del nombre de una func entre parentesis la struct que va asociar. 
    

    var albums = []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    }
    fmt.Println(albums)
}