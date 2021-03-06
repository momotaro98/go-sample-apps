# 4.7 ファンアウト、ファンイン

時折、パイプラインのあるステージで特に計算量が大きくなることがあります。
ファンアウトはパイプライからの入力を扱うために複数のゴルーチンを起動するプロセスを説明する用語です。そしてファンインは複数の結果を1つのチャネルに結合するプロセスを説明する用語です。

ファンアウトの利用状況

* そのステージ内で並列処理される他の計算結果に依存していない。
* 実行が長時間におよぶ。

# Run

### Fast primeClassifier

```
$ go run bad.go stage_funcs.go
Primes:
        19727887
        43516159
        38043721
        45071563
        49509107
        16711297
        6403981
        48898981
        30599183
        15889513
Search took: 858.674µs
```

### Slow primeClassifier

```
$ go run bad.go stage_funcs.go
Primes:
        19727887
        43516159
        38043721
        45071563
        49509107
        16711297
        6403981
        48898981
        30599183
        15889513
Search took: 2.436096632s
```

### Fan-out Fan-in program

```
$ go run good.go fan_in.go stage_funcs.go
Spinning up 4 prime finders.
Primes:
        19727887
        38043721
        43516159
        45071563
        6403981
        16711297
        49509107
        15889513
        30599183
        48898981
Search took: 1.209089947s
```
