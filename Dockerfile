# Use uma imagem base do Go
FROM golang:1.23.1

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos do projeto para o contêiner
COPY . .

# Baixe as dependências
RUN go mod tidy

# Construa o binário
RUN go build -o .cmd/main ./cmd

# Exponha a porta usada pela aplicação
EXPOSE 8080

# Comando para rodar a aplicação
CMD [".cmd/main"]
