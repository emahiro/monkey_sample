# monkey_sample
[「Go言語でつくるインタプリタ」](https://amzn.to/2LgPeoW)の学習サンプル

# Usage
## Lexer

```shell
$ go run ./src/monkey/main.go
hello emahiro!. this is the Monkey programming language.
feel free ato type in command following.
>>
let five = 5;
{Type:LET Literal:let}
{Type:IDENT Literal:five}
{Type:= Literal:=}
{Type:INT Literal:5}
{Type:; Literal:;}
```

lexerによって字句解析されトークン化されたソースコードを取得できる。
