# Segunda Avaliação de Desenvolvimento Web

## Desenvolvimento de API com conexão à banco de dados

Com o intúito de desenvolver as habilidades no que se refere a desenvolvimento de software, foi proposta a seguinte atividade, que se caracterizava por desenvolver uma aplicação _Backend_ que pudesse realizar todas as operações básicas de um __CRUD__(_Create, Read, Update, Delete_) em um banco de dados do tipo _NoSQL_, em específico, utilizando a linguagem _Javascript_ com o __ODM__(_Object Document Mapping_) _Mongoose_, e o banco de dados _NoSQL_ _MongoDB_, que se trata de um banco de dados que realiza o armazenamento por meio de __JSON__(_Javascript Object Notation_) e fazendo o uso da plataforma Atlas, que é um serviço que oferece de forma gratuita e facilitada a hospedagem do Banco.

Por escolha própria, decifi fazer utilizando a linguagem de programação _Golang_, fazendo uso das bibliotecas _GORM_ para lidar com os dados de forma facilidata, funcionando como um __ORM__(_Object Relational Mapping_) e o banco de dados _Postgres_, um banco _SQL_ amplamente utilizando pela indústria de tecnologia.

Para facilitar no deploy local da aplicação, foi utilizado as ferramentas _Docker_, para a criação do container da aplicação, e _Docker Compose_ para fazer uso do banco de dados em forma de container, sem a necessidade de instalação tanto da linguagem _Golang_ nem do banco _Postgres_.

### Rodando a aplicação

Para rodar a aplicação, basta estar no diretório raiz da mesma, onde se localiza o arquivo _docker-compose.yaml_ e a respectiva _Dockerfile_ da aplicação.

Estando no diretório raiz, basta utilizar o comando:

```bash
docker-compose up --build
```

Após isso, verificado os logs do container e visto que a aplicação está rodando, basta acessar as rotas especifica para cada uma das opções na sua ferramente de requisições preferida. As rotas estão descritas no arquivo _main.go_, na pasta _cmd_

```golang
router.GET("/foods/all", foodHandler.GetAllFoods)
router.GET("/foods/:id", foodHandler.GetFoodsByID)
router.GET("/foods/", foodHandler.GetAllFoodsByCategory)
router.POST("/foods", foodHandler.CreateFood)
router.PUT("/foods/:id", foodHandler.UpdateFood)
router.DELETE("/foods/:id", foodHandler.DeleteFood)
```

Basta adicionar as urls acima ao seu endereço de _localhost_

#### Acessa todos os recursos, no caso, uma lista com as comidas no servidor é retornada

__GET:__ <http://localhost:8080/foods/all>

#### Cria um registro no servidor

__POST:__ <http://localhost:8080/foods>

O corpo da requisição deve ser passado no sequinte formato:

```javascript
{
    name: "nome-da-comida"
    category: "categoria-da-comida"
    quantity: quantidade-da-comida
    expiration_at: "data-no-padrão-ISO-8601"

}
```

Não é necessário passar o ID no corpo da requisição, pois o _GORM_ Autoincrementa o ID

#### Modifica um registro existente no servidor

__Put:__ <http://localhost:8080/foods/id-da-comida>

```javascript
{  
"name": "Nome-da-comida",
"category": "Categoria-da-comida",
"quantity": quantidade-inteira,
"price": preço-float,
"expiration_at": "data-no-padrão-ISO-8601"
}
```

#### Para deletar uma comida

__Delete:__ <http://localhost:8080/foods/id-da-comida>

É retornado o _status code_ da operação e, caso seja concluída com sucesso, é retornado o registro que foi deletado

#### Retorna o registro com o ID informado

__Get By ID:__ <http://localhost:8080/foods/id-da-comida>

#### Retorna todos os registro de determinada categoria

__Get By Category Query:__ <http://localhost:8080/foods?category=categoria-da-comida>

#### Exemplo de retorno de um registro

```javascript
{
  "ID": 1,
  "CreatedAt": "2024-04-08T18:49:50.259524Z",
  "UpdatedAt": "2024-04-08T18:49:50.259524Z",
  "DeletedAt": null,
  "name": "Lasanha",
  "category": "Massas",
  "quantity": 20,
  "price": 25.99,
  "expiration_at": "2024-05-08T18:49:50.160252Z"
}
```

__PS:__ Os campos _ID_, _CretedAt_, _UpdateAt_ e _DeletedAt_ são criados automaticamente pelo _GORM_ e são utilizados para informações e ajuda ao manipular os registros
