FROM golang:1.15.0
WORKDIR /go/src/github.com/Abdelmounaim-Azz/inject-commit
COPY .git .
COPY app.go .
RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
  go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
CMD ["./git-tester", "-version"]