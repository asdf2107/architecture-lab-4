# Lab 4

## How to build

1. Copy project folder:

```bash
cp -r . $GOPATH/github.com/asdf2107/
```

Ensure `$GOPATH` is defined

2. Install packages which are inside of each folder in src/

```bash
cd src;
for folder in $(ls); do
    cd $folder;
    go install;
    cd ..;
done;
```

3. Just `go build` and that's it!
