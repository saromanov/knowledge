FROM golang:alpine
RUN mkdir /knowledge
COPY . /knowledge
WORKDIR /knowledge
RUN apk add git && go build -o main ./cmd/knowledge 
CMD ["/app/knowledge"]