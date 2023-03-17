# go-frequency-analysis
古典暗号の換字式暗号に有効な頻度分析を行う。アルファベットのみ対応。

## Usage

### 引数を渡さなかった場合は `Hello, World!` の頻度分析を行う
```
$ go run .
Hello, World!
L: 30.000%
O: 20.000%
H: 10.000%
E: 10.000%
W: 10.000%
R: 10.000%
D: 10.000%
```

### 引数で文字列を渡すことで任意の英文の頻度分析が可能。([例: 日本国憲法 第一章 第一条](https://www.japaneselawtranslation.go.jp/ja/laws/view/174))
```
$ go run . "The Emperor shall be the symbol of the State and of the unity of the people, deriving his position from the will of the people with whom resides sovereign power."
The Emperor shall be the symbol of the State and of the unity of the people, deriving his position from the will of the people with whom resides sovereign power.
E: 16.154%
O: 10.769%
T:  9.231%
H:  8.462%
I:  7.692%
S:  6.154%
R:  5.385%
P:  5.385%
L:  5.385%
N:  3.846%
F:  3.846%
W:  3.077%
M:  3.077%
D:  2.308%
A:  2.308%
V:  1.538%
G:  1.538%
B:  1.538%
Y:  1.538%
U:  0.769%
```
