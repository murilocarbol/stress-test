# Stress-Test CLI - Sistema de Teste de Carga
Este projeto é uma ferramenta criada em GoLang para realizar testes de carga em serviços web. O usuário pode especificar a URL a ser testada, o número total de requisições e a quantidade de concorrências. Ao final, a ferramenta gera um relatório com informações detalhadas sobre o teste realizado.

## Funcionalidades  
- Enviar requisições HTTP para uma URL especificada.
- Controlar o número total de requisições e a quantidade de concorrências.
- Gerar um relatório com:
  - Tempo total gasto.
  - Número total de requisições realizadas.
  - Contagem de requisições com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (404, 429, 500, etc.).

## Como executar

### Pré-requisitos
- [Docker](https://www.docker.com/) instalado.
- Go 1.23.

### Build e execução com Docker

1. **Build da imagem Docker:**
   ```sh
   docker build -t stress-test .
   ```

2. **Executar a aplicação:**
   ```sh
   docker run stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

### Parâmetros disponíveis:
- `--url`: (Obrigatório) URL do serviço web a ser testado.  
- `--requests`: (Obrigatório) Número total de requisições a serem realizadas.
- `--concurrency`: (Obrigatório) Número de requisições simultâneas.



## Exemplo de Uso  

```sh
docker run stress-test --url=http://example.com --requests=500 --concurrency=20
```

### Saída esperada:  

```text
Iniciando teste de carga...
Tempo total: 10.5s
Total de requisições: 500
Requisições bem-sucedidas (HTTP 200): 450
Outros códigos de status:
  - 404: 30
  - 500: 20
```



## Estrutura do Projeto  

```bash
.
├── application/
│     └── client/ 
│           └── generic_client.go # Arquivo que contem a estrutura de requisição http
│     └── usecase/ 
│           └── stress_usecase.go # Arquivo que contem a lógica do processamento
├── cmd/
│     └── main.go                 # Arquivo principal da aplicação
├── config/
│     └── config.go               # Arquivo de configuração do aplicação
├── go.mod                        # Arquivo de dependências Go
├── Dockerfile                    # Configuração para criar a imagem Docker
└── README.md                     # Documentação do projeto
```
