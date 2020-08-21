FROM golang

WORKDIR /app
COPY . ./

EXPOSE 80

CMD [ "go", "run", "main.go" ]