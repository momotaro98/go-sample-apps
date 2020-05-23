
# 2章 構文解析

## 2.3 Monkey言語のための構文解析器を書く

> プログラミング言語の構文解析には、大きく分けて2つの戦略がある。トップダウン構文解析とボトムアップ構文解析だ。どちらの戦略にも細かなバリエーションがある。例えば、再帰下降構文解析(recursive descent parsing)やアーリー方(Early parsing)、予測的構文解析(predictive parsing) はいずれもトップダウン構文解析法のバリエーションだ。

> これから書こうとしている構文解析器は再帰下降構文解析器だ。より詳しく言うと「トップダウン演算子優先順位」構文解析器で、考案者Vaughan Prattの名前を取って「Pratt構文解析器」とも呼ばれる。

## 2.4 構文解析の第一歩: let文

```
lex x = 10;
let y = 15;

let add = fn(a, b) {
  return a + b;
};
```

以下のように表せる。

```
let <identifier> = <expression>;
```

## 2.6.2 トップダウン演算子順位解析(Pratt構文解析)

> 論文"Top Down Operator Precedence"において、Vaughan Prattは式を構文解析する手法を提案している。

> > とてもシンプルで理解しやすく、実装も容易で、使いやすく、理論上はともかく実用上は非常に効率的であり、それでいて、利用者のほとんどの合理的な構文の要求に耐えられるほど柔軟である。

> この論文は1973年に出版された。以来、長年に渡って、Prattの提案した考えはそこまで大きな支持を集めたわけではなかった。ごく近年になって、他のプログラマたちがPrattの論文を再発見し、それについて記事を書いたことでPrattの構文解析のアプローチが注目されるようになった。
