FROM golang:alpine
RUN mkdir /knowledge && mkdir /app
COPY . /knowledge
WORKDIR /knowledge
RUN apk add git && go build -o knowledge ./cmd/knowledge && cp ./knowledge /app/knowledge &&\
chmod +x /app/knowledge && rm -rf /knowledge
CMD ["/app/knowledge"]