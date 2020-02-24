# backend core 

#### Env config
Make sure you export this variables in your ~/.shrc or ~/.zshrc
```
echo $GOPATH
/Users/USERNAME/go/

echo $GOROOT
/usr/local/go/

echo $PATH
...:/usr/local/go/bin:/Users/USERNAME/go//bin
:/usr/local/go//bin
```

#### Get deps
``` 
~ cd
~ go get -u github.com/golang/dep/cmd/dep
~ go get -u github.com/pilu/fresh
~ go get -u github.com/labstack/echo/...

```

#### Start

```
~ dep ensure -update
~ fresh
```

#### Testing

```
~ go test -v ./... -cover
```

#### Config

```

```


