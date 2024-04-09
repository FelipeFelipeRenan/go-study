FROM golang:1.22 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o restante do código-fonte
COPY . .

# Baixa as dependências do módulo
RUN go mod download
# Compila o código
RUN go build -o main ./cmd


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