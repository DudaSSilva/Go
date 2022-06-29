package main

import "fmt"

type Nota struct{
  nota1 float64
  nota2 float64
}

type Aluno struct{
  nome string
  nascimento int
}

type Cadastro struct{
  aluno Aluno
  data int
  notas Nota
}

func (c Cadastro) idade() int{
   return c.data - c.aluno.nascimento
}

func (n Nota) media() float64{
  return (n.nota1 + n.nota2)/2
}

func main() {
  var alunoNome string 
  var dataNascimento int
  var notaUm, notaDois float64
  var anoAtual int
  var exibir string
  var cadastrar string = "sim"

  fmt.Print("\n\n========== CADASTRO ==========\n\n")

  for cadastrar != "nao"{
    fmt.Print("\n\n---------- DADOS PESSOAIS----------\n\n")
    fmt.Println("\nInforme o nome do aluno: ")
    fmt.Scan(&alunoNome)
    fmt.Println("\nInforme o ano de nascimento do aluno: ")
    fmt.Scan(&dataNascimento)
    fmt.Println("\nInforme o ano atual: ")
    fmt.Scan(&anoAtual)
    fmt.Print("\n\n---------- NOTAS ----------\n\n")
    fmt.Println("\nDigite a primeira nota: ")
    fmt.Scan(&notaUm)
    fmt.Println("\nDigite a segunda nota: ")
    fmt.Scan(&notaDois)

    notaAluno := Nota{
      nota1: notaUm,
      nota2: notaDois,
    }

    pessoa := Aluno{
      nome: alunoNome,
      nascimento: dataNascimento,
    }

    cadastro := Cadastro{
      aluno: pessoa,
      data: anoAtual,
      notas: notaAluno,
    }

    fmt.Println("Aluno cadastrado! Deseja exibir as informações deste aluno?")
    fmt.Scan(&exibir)

    if exibir == "sim"{
      fmt.Println("\nNome:", cadastro.aluno.nome, "\nData de Nascimento: ", cadastro.aluno.nascimento, "\nIdade: ", cadastro.idade(), "\nMédia: ",notaAluno.media(), "\n")
    }else{
      fmt.Println("\nCadastro finalizado.")
    }

    fmt.Println("\nDeseja cadastrar outro aluno?")
    fmt.Scan(&cadastrar)
  }
}
