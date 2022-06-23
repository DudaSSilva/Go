package main

import "fmt"

func main() {

  var n1, n2 float32
  var operacao string
  var sair string = "sim"

  fmt.Print("\n\n========== CALCULADORA ==========\n\n")
  
  for sair != "nao"{
    fmt.Println("\nDigite o primeiro número: ")
    fmt.Scan(&n1)
    fmt.Println("Digite o segundo número: ")
    fmt.Scan(&n2)
    fmt.Println("\nInforme o sinal correspondente à operação que deseja efetuar: ")
    fmt.Scan(&operacao)
    if operacao == "+"{
      fmt.Printf("\nA soma de %.f e %.f é: %.f \n", n1, n2, n1 + n2)
    }else if operacao == "-"{
      fmt.Printf("\nA subtração de %.f e %.f é: %.f \n", n1, n2, n1 - n2)
    }else if operacao == "*"{
      fmt.Printf("\nO produto de %.f e %.f é: %.f \n", n1, n2, n1 + n2)
    }else if operacao == "/"{
      fmt.Printf("\nA divisão entre %.1f e %.1f resulta em: %.2f \n", n1, n2, n1 / n2)
    }else{
      fmt.Println("\nOperação inválida!")
    }    
    fmt.Println("\nQuer realizar outra operação? ")
    fmt.Scan(&sair)
  }
  fmt.Print("\n\nATÉ A PRÓXIMA!")
}
