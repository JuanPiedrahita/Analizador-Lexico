//Presentado Por:
//Juan Sebastian Piedrahita 20141020036
//Juan David  Cubillos 20141020050
//Edison Peñuela 20141020018
//Grupo 81 
//Ciencias de la Computacion 3
package main 

import(
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"bufio"
	"os"
)

//pila de arboles
func NewStack() *Stack {
	return &Stack{}
}


type Stack struct {
	nodes []*Arbol
	count int
}

func (s *Stack) Push(n *Arbol) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Arbol {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

//pila de float64
func NewPila64() *Pila64 {
	return &Pila64{}
}


type Pila64 struct {
	nodes []float64
	count int
}

func (s *Pila64) Push(n float64) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Pila64) Pop() float64 {
	if s.count == 0 {
		return 0
	}
	s.count--
	return s.nodes[s.count]
}


//tipo arbol
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
	tipoNodo string
}

func NewArbol(s string) *Arbol {	
	return &Arbol{nil,s,nil,""}
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInorden(t.Izquierda)
	fmt.Print(t.Valor)
  	fmt.Print("  ")
	RecorrerInorden(t.Derecha)
}

func RecorrerExpresiones(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInordenTipos(t)
	fmt.Println("")
	RecorrerInorden(t)
	fmt.Println("")
	RecorrerExpresiones(t.Izquierda)
	RecorrerExpresiones(t.Derecha)
}

func RecorrerInordenTipos(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInordenTipos(t.Izquierda)
	fmt.Print(t.tipoNodo)
  	fmt.Print("  ")
	RecorrerInordenTipos(t.Derecha)
}

//funcion para armar arbol
func armarArbol(m map[string]Arbol, s []string ) (*Arbol){
	pila := NewStack()
	for j:=0;j<len(s);j++{
		if(esEntero(s[j])){
			pila.Push(&Arbol{nil,s[j],nil,"Valor"})
		} else if (esVariable(s[j])) {
			if(enMapArbol(m, s[j])){
				arbol:=m[s[j]]
				var subExpresion *Arbol
				if(arbol.Derecha.tipoNodo=="Variable"){
					subExpresion = arbol.Izquierda
				}else{
					subExpresion = arbol.Derecha
				}
				pila.Push(subExpresion)
			} else {
				pila.Push(&Arbol{nil,s[j],nil,"Variable"})
			}
		} else {	
			if pila.count>=2{
				h2:= pila.Pop()
				h1:= pila.Pop()
	
				pila.Push(&Arbol{h1,s[j],h2,"Expresion"})
				}else{
					panic("Error de stack, no hay suficientes elementos.")
				}
			}
		}
	return pila.Pop()
}
//en Hashmap
func enMapArbol(m map[string]Arbol, x string) bool{
	for key, _ := range m{
		if(key == x){
			return true
		}
	}
	return false
}
//en Hashmap
func enMap(m map[string]float64, x string) bool{
	for key, _ := range m{
		if(key == x){
			return true
		}
	}
	return false
}
//Funcion evaluar expresion
func evaluarExpresion(s []string, m map[string]float64) (string, float64){
	pila := NewPila64()
	var variable string
	for j:=0;j<len(s);j++{
		if(esVariable(s[j])){
			if enMap(m, s[j]){
				pila.Push(m[s[j]])
			}else{
				variable = s[j]
			}
		}else if(esEntero(s[j])){
			i, _ := strconv.ParseFloat(s[j], 64)
			pila.Push(i)
		}else{	
			if pila.count>=2{
				valor2 := pila.Pop()
				valor1 := pila.Pop()
				var res float64
				switch s[j] {
				  case "+":
		 				  res=valor1+valor2
				  case "-":
		       				  res=valor1-valor2
				  case "*":
		       				  res=valor1*valor2 
				  case "/":	
						  res=valor1/valor2	
				} 
				pila.Push(res)
			}else if(s[j]==":="){
				
				return variable, pila.Pop()
			}else{
				panic("Error de stack, no hay suficientes elementos.")
			}
		}
	}
	return "", pila.Pop()
}

//esEntero
func esEntero(s string) bool{
	entero:=regexp.MustCompile("^([0-9])+$")
	return entero.MatchString(s)
}
//esSimbolo
func esSimbolo(s string) bool{
	simbolo:=regexp.MustCompile("^[-|+|*|/]$")
	if !simbolo.MatchString(s){
		simbolo=regexp.MustCompile("^(:=)$")
	}
	return simbolo.MatchString(s)
}
//esVariable
func esVariable(s string) bool{
	variable:=regexp.MustCompile("^[A-Z]([_|A-Z|a-z|0-9]*?)$")
	return variable.MatchString(s)
}
func validarTokens(x []string){
	for j:=0;j<len(x);j++{
		if(!esSimbolo(x[j]) && !esEntero(x[j]) && !esVariable(x[j])){
			panic("Token no valido"+ x[j])
		}
	}
}



func main(){
	m := make(map[string]float64)
	mapArboles := make (map[string]Arbol)
	seguir := "s"
	for seguir == "s"{
		fmt.Println("Ingrese la expresion posfija (cada termino separado por ' '):")
		reader:=bufio.NewReader(os.Stdin)
		x, _ := reader.ReadString('\n')
		x = x[0:len(x)-1]
		s := strings.Split(x, " " )
		
		validarTokens(s)

		arbol := armarArbol(mapArboles, s)
	
		//fmt.Println("Expresion infija: ")
		//RecorrerInorden(arbol)
		//fmt.Println("")
		//fmt.Println("Expresion infija tipo nodos: ")
		//RecorrerInordenTipos(arbol)
		//fmt.Println("")
		fmt.Println("Expresiones: ")
		RecorrerExpresiones(arbol)
		fmt.Println("")

		variable, valor:= evaluarExpresion(s,m)
		fmt.Println(variable,"=",valor)
		m[variable] = valor
		mapArboles[variable] = *arbol
		fmt.Println("Desea ingresar otra expresion [s/n]:")
		fmt.Scan(&seguir)
	}	
}
