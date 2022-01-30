FROM golang:alpine
RUN mkdir /knowledge
COPY . /knowledge
WORKDIR /knowledge
RUN go build -o main ./cmd/knowledge 
CMD ["/app/knowledge"]