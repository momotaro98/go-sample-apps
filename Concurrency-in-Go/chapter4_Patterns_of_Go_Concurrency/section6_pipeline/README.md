# パイプラインステージの性質

* ステージは受け取るものと返すものが同じ型である。
* ステージは引き回せるように具体化されていなければならない。Goにおいて関数は具体化されているため(※1)、この目的にうまく適合している。

※1: Goでは関数シグネチャの型を持つ変数を定義することができるという意味で言っている。またこれは関数をプログラムの中で変数として渡せることも意味する。

> パイプラインのステージ関数は関数プログラミングと密接に関係していて、モナドのセブセットと"考える"こともできます。