list:
    @just --list

build_cmd:
    cd cmd/bounded-motions && go build .

purge:
    rm -rf cmd/bounded-motions/bounded-motions

bootstrap:
    go get github.com/rivo/tview@master