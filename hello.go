package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("exibindo logs...")
			imprimeLog()
		} else if comando == 0 {
			fmt.Println("saindo do programa irmao...")
			os.Exit(0)
		} else {
			fmt.Println("nao reconheço esse comando")
			os.Exit(-1)
		}
	}
}
func exibeIntroducao() {
	nome := "thalao"
	idade := 18
	versao := 1.5
	fmt.Println("ola, sr.", nome, "sua idade é:", idade)
	fmt.Println("este programa esta na versao:", versao)
}
func exibeMenu() {
	fmt.Println("1-iniciar monitoramento")
	fmt.Println("2-exibir logs")
	fmt.Println("0-sair do programa")
}
func leComando() int {
	var comandolido int
	fmt.Scan(&comandolido)
	return comandolido
}
func iniciarMonitoramento() {
	const delay = 15
	const monitoramentos = 3
	fmt.Println("monitorando...")
	sites := leArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("site", site, "nao esta carregando, erro:", resp.StatusCode)
		registraLog(site, false)
	}
}
func leArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "- " + site + " -online: " + strconv.FormatBool(status) + "\n")
	fmt.Println(arquivo)
	arquivo.Close()
}
func imprimeLog() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))

}
