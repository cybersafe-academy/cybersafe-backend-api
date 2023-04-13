# CyberSafe Academy API

Este projeto faz parte do Cybersafe Academy, uma empresa especializada em segurança da informação e fornecedora de soluções exclusivas para empresas do setor financeiro. Nosso objetivo é fornecer soluções que garantam a segurança e proteção das informações confidenciais dos nossos clientes, ao mesmo tempo em que oferecemos insights valiosos sobre possíveis vulnerabilidades na segurança.


## Instruções de instalação do Swaggo

Para gerar documentação da API com o Swaggo, é necessário instalar a ferramenta em seu ambiente de desenvolvimento. Para isso, siga as instruções abaixo:

Execute o seguinte comando para instalar o Swaggo:

```
go get -u github.com/swaggo/swag/cmd/swag

```

Verifique se a instalação foi bem-sucedida executando o seguinte comando:

```
swag --version
```

Se tudo estiver configurado corretamente, o comando deverá retornar a versão atual do Swaggo.

## Iniciando a API
Certifique-se de ter instalado em sua máquina o Docker, Docker Compose `Go 1.20` e [Make](https://www.gnu.org/software/make/ "Make") .
Para iniciar a API, execute o seguinte comando:

```
make run
```

Este comando vai iniciar o container Docker com as dependências do projeto. 

Para visualizar a documentação da API gerada pelo Swaggo, execute o seguinte comando:

```
make swag
```

Isso irá gerar a documentação e disponibilizá-la em http://localhost:8080/swagger/index.html.

## Autores
- João Victor Marques - https://www.linkedin.com/in/joaovmvale/
- Vinicius Taborda - https://www.linkedin.com/in/viniciustaborda/

## Licença
Este projeto é licenciado sob a Licença MIT - consulte o arquivo [LICENSE](https://www.mit.edu/~amini/LICENSE.md "LICENSE") para obter mais detalhes.
