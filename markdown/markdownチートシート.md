# Markdownチートシート
- [見出し](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#見出し)
- [リスト](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#リスト)
- [番号付きリスト](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#番号付きリスト)
- [空行・改行](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#空行・改行)
- [インライン表示](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#インライン表示)
- [コードの挿入](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#コードの挿入)
- [リンクの挿入](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#リンクの挿入)
- [引用](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#引用)
- [画像の挿入](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#画像の挿入)
- [画像の挿入(サイズ指定)](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#画像の挿入(サイズ指定))
- [テーブル](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#テーブル)
- [文字色](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#文字色)
- [太字](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#太字)
- [斜体](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#斜体)
- [打消し線](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#打消し線)
- [水平線](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#水平線)
- [注釈](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#注釈)
- [マークダウンのエスケープ](https://github.com/sbtosh/tips/blob/main/markdown/markdownチートシート.md#マークダウンのエスケープ)

# 見出し
- 書式
```
# テスト
## テスト１
### テスト２
```
- 表示
># テスト
>## テスト１
>### テスト２
# リスト
- 書式
```
* テキスト
    * テキスト
```
または
```
- テキスト
    - テキスト
```
- 表示
>* テキスト
>    * テキスト
# 番号付きリスト
- 書式
```
1. テキスト
2. テキスト
```
- 表示
>1. テキスト
>2. テキスト
# 空行・改行
- 書式
```
文末にスペース2つ。または改行2回。
```
- 表示
>テキスト  
>テキスト1
>
>テキスト2
# インライン表示
- 書式
```
`テキスト`
```
- 表示
>変数は`int num = 0`で定義されている
# コードの挿入
- 書式

```
```コード等```
```
- 表示
>```
>テキスト
>```
# リンクの挿入
- 書式
```
[タイトル](URL)
```
- 表示
>[github.com](https://github.com/)
# 引用
- 書式
```
> テキスト
>> テキスト1
```
- 表示
> テキスト
>> テキスト
# 画像の挿入
- 書式
```
![空でもOK](画像URL "タイトル付け可能")
```
- 表示
> ![test.png](https://github.com/sbtosh/tips/blob/main/img/test.png "テスト")
# 画像の挿入(サイズ指定)
- 書式
```
<img width="数値" alt="テキスト" src="画像URL">
```
- 表示
> <img width="100" alt="テスト" src="https://github.com/sbtosh/tips/blob/main/img/test.png">
# テーブル
- 書式
```
|a|b|c|
|:--|--:|:--:|
|hoge|hoge|hoge|
|d|d|d|
```
- 表示
>|a|b|c|
>|:--|:--:|--:|
>|tmp|tmp|tmp|
>|left|center|right|
# 文字色
- 書式
```
<font color="Red">テキスト</font>
```
- 表示
> <font color="Red">テキスト</font>
# 太字
- 書式
```
**テキスト**
__テキスト__
```
- 表示
> **テキスト**
>
> __テキスト__
# 斜体
- 書式
```
*テキスト*
_テキスト_
```
- 表示
> *テキスト*
>
> _テキスト_
# 打消し線
- 書式
```
~~打ち消し~~
```
- 表示
> ~~打ち消し~~
# 水平線
- 書式
```
***
```
- 表示
>
> ***
>
# 注釈
- 書式
```
テキスト[^1]
[^1]: 注釈内容
```
- 表示
> テキスト[^1]

> [^1]: 注釈内容
# マークダウンのエスケープ
- 書式
```
\記号
```
- 表示
> \`インライン表示無効`