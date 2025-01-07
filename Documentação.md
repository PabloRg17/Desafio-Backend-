Abaixo está a documentação do desafio técnico para backend em Golang da Taghos.

 ## Primeiramente deixar claro, que eu fui perceber que o desafio era em Golang depois que aceitei, não cursei Go e tive que aprender o básico de Go e PostGresql para poder fazer o desafio. Tive que aceitar alguma ajuda do Chatgpt, mas foi divertido.

 ## Funcionalidades

- **Criar Livro**: Adicionar um novo livro ao banco de dados.
- **Buscar Livro**: Consultar informações de um livro específico pelo seu ID.
- **Listar Todos os Livros**: Obter uma lista de todos os livros cadastrados.
- **Atualizar Livro**: Atualizar informações de um livro existente.
- **Excluir Livro**: Remover um livro do banco de dados.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para o back-end.
- **Gin**: Framework web para Go utilizado para criar a API RESTful.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar as informações dos livros.
- **pgx**: Driver PostgreSQL para Go.
- **godotenv**: Carrega variáveis de ambiente a partir de um arquivo `.env`.

## Estrutura de Diretórios

 ├── db.go # Conexão com o banco de dados PostgreSQL.
 ├── handlers.go # Funções responsáveis pelas rotas e lógica de negócios. 
 ├── main.go # Arquivo principal que inicializa o servidor e as rotas. 
 ├── .env # Arquivo de configuração com variáveis de ambiente.
 ├── go.mod # Arquivo de gerenciamento de dependências Go. 
  └── Documentação.md # Documentação do projeto.

  ################################################


## Detalhamento dos Arquivos

### 1. **db.go**

Este arquivo contém a função responsável pela conexão ao banco de dados PostgreSQL.

- **Função `ConnectToDB()`**: 
  - Carrega as variáveis de ambiente do arquivo `.env` usando a biblioteca `godotenv`.
  - Cria uma string de conexão utilizando as variáveis carregadas e estabelece a conexão com o banco de dados PostgreSQL.
  - Se a conexão falhar, um erro é registrado e o processo é encerrado.

### 2. handlers.go
Este arquivo contém as funções responsáveis por manipular as requisições HTTP para criar, ler, atualizar e excluir livros.

Estrutura Book: Define o modelo de dados para um livro, com os campos ID, Title, Category, Author e Synopsis. O json tag é usado para garantir que o JSON seja serializado corretamente.

**Função CreateBook():** Insere um novo livro no banco de dados, retornando o livro com o ID gerado.

**Função GetBook():** Recupera as informações de um livro específico a partir do seu ID. Retorna um erro se o livro não for encontrado.

**Função GetAllBooks():** Retorna uma lista de todos os livros cadastrados.

**Função UpdateBook():** Atualiza as informações de um livro existente no banco de dados.

**Função DeleteBook():** Exclui um livro do banco de dados a partir do seu ID.

### 3. main.go
Este é o arquivo principal que inicializa o servidor e define as rotas da aplicação.

**Função main():**
Chama a função ConnectToDB() para conectar ao banco de dados.
Define as rotas da API utilizando o framework Gin.
Inicia o servidor na porta 8080.

### 4. .env
Este arquivo contém as variáveis de ambiente para a conexão com o banco de dados.

__________________________ ### _______________________________

 ### Como Rodar o Projeto
 **1. Configuração do Banco de Dados**:

 - Crie um banco de dados no PostgreSQL com o nome definido na variável DB_NAME no arquivo .env.

 - Execute as migrações necessárias para criar a tabela books.

 **2. Instalar Dependências:** 

 - Execute o comando abaixo para instalar as dependências:
         go mod tidy 


 **3. Rodar a Aplicação:** 

 - Execute o comando abaixo para rodar o servidor:
        go run main.go db.go handlers.go 

 **4. Acessar a API:** 

 - A API estará disponível em http://localhost:8080/.
 - Você pode testar as rotas utilizando ferramentas como Postman ou curl.
