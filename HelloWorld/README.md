# Hello World

## main()

- golang は main 関数がないとエラーになる。

```go
func main(){
  fmt.Println("Hello world!")
}
```

## init()

- init 関数は main 関数よりも前に実行される。
- 初期設定などに使われる。

```go
func init(){}
```
