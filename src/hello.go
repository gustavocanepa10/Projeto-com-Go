package main
import "fmt"
import "os"

func main() {
	exibeIntroducao()
	exibeMenu()
	
	


    
	
    comando :=leComando()

	switch comando { 
	   case 1:
		fmt.Println("Monitorando...")
	   case 2:
		fmt.Println("Exibindo Logs...")
	   case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	   default:
		fmt.Println("Não conheco essa instrução")


	}

}

func exibeIntroducao(){
	nome:="Gustavo"
	versao:=1.1
	fmt.Println("Olá,sr.",nome)
	fmt.Println("Este programa está na versao",versao)




}

func exibeMenu(){
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs ")
	fmt.Println("0- Sair do programa ")





}



func leComando()int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi",comandoLido)


	return comandoLido



}
