# Variable

## 基本

var ステートメントを使って変数名、型、値を記述します。

```go
// int型のiという変数名に1を代入しています。
var i int = 1
```

下記のように省略して書くことができます。

```go
var t, f bool = true, false
```

## 変数宣言のグループ化

import の時と同じように()で括って複数の変数を宣言することができます。

```go
var (
  i int = 1
  s string = "hoge"
)
```

## 値を代入しないで変数を宣言

値を代入しないで変数を宣言することもできます。その場合はデフォルト値が入ります。

```go
	var (
		i int
		f bool
	)
	fmt.Println(i, f)
  // 0 false が出力される
```

## Short variable declarations

関数内であれば、var を使わなくても変数を宣言することができます。

```go
s := "hoge"
```

Short variable declarations の書き方の場合、型も省略されます。

### var と :=　の使い分け

- Short variable declarations は関数内でしか使えない。
- 明示的に型を宣言したい場合は、var を使う。

```go
f := 1.2
fmt.Printf("%T", f)
// float64　が出力される。
```

例えば、上記の変数 f の型を float32 として扱いたい場合は明示的に var で宣言してあげる必要があります。

```go
var f float32 = 1.2
fmt.Printf("%T", f)
// float32　が出力される。
```
