# Imagem base
FROM golang:1.22 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o arquivo go.mod e go.sum
COPY evento-service/go.mod evento-service/go.sum ./

# Baixa as dependências do módulo
RUN go mod download

# Copia o restante do código-fonte
COPY evento-service/ .

# Compila o código
RUN go build -o main .

# Etapa de construção final
FROM golang:1.22

# Define o diretório de trabalho
WORKDIR /app

# Copia o executável compilado da etapa anterior
COPY --from=builder /app/main .

# Exponha a porta que a aplicação Go irá ouvir
EXPOSE 8080

# Comando padrão para executar o aplicativo quando o contêiner for iniciado
CMD ["./main"]