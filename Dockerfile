FROM golang:1.15

ENV REPO_URL=github.com/DataPsycho/portfolio
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src

COPY src $WORKPATH

WORKDIR $WORKPATH

RUN go build -o portfolio .

# Expose Port
EXPOSE 8080 

CMD ["./demoapp"]