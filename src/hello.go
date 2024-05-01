package main
import "fmt"
import "os"
import "net/http"


func main() {

	exibeIntroducao()
	exibeMenu()
	
	
	
	
    comando :=leComando()

	switch comando { 
	   case 1:
		iniciarMonitoramento()
	   case 2:
		fmt.Println("Exibindo Logs...")
	   case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	   default:
		fmt.Println("Não conheco esse comando")
		os.Exit(-1)


	}

}

func devolveNome () string {
	nome := "Gustavo"
	return nome



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
	fmt.Scanln(&comandoLido)
	fmt.Println("O comando escolhido foi",comandoLido)
	


	return comandoLido



}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	site :="http://www.alura.com.br"
	resp,_:=http.Get(site)
	fmt.Println(resp)

    if resp.StatusCode ==200{
		fmt.Println("Site:",site,"foi carregado com sucesso!")
	
    
	}else{ 
		fmt.Println("Site:",site,"Está com problemas. Status code ",resp.StatusCode)
		
		
		



	}
	
    




	





}
