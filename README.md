# Sistema de Stress test
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI:

- --url (-u): URL do serviço a ser testado.
- --requests (-r): Número total de requests.
- --concurrency (-c): Número de chamadas simultâneas.
- **--log (-l) : [opcional] Printar os logs em tempo real de cada tentativa.**

## Execução do Teste:

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.
- Geração de Relatório:

## Apresentar um relatório ao final dos testes contendo:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Executar Localmente 
```bash
go run main.go --url <url> --requests <requests> --concurrency <concurrency>
```
ou
```bash
go run main.go -u <url> -r <requests> -c <concurrency>
```

## Executar com Docker
### Criar a imagem com o Dockerfile
```bash
docker build -t gostress .
```
### Executar o teste
```bash
docker run gostress --url=http://google.com --requests=30 --concurrency=2
```

### Exemplo de saída do teste
```bash
Initializing test 🧐
Finishing test 🥵
Stress Report 🚀...

URL:                    http://google.com
Requests:               30
Concurrency:            2
Duration:               8.194479709 s
Total Status Code 200:  0
Total Status Code 429:  30
```