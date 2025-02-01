# count-token

`count-token` は、標準入力からテキストを読み込み、そのテキストに含まれるトークンの数を数えるシンプルなコマンドラインツールです。

## 使い方

1. [リリース](<https://github.com/thr3a/llm-token-count/releases>)ページから、あなたの環境に合った実行ファイルをダウンロードしてください。
2. ダウンロードしたファイルを解凍してください。
3. ターミナルで解凍したディレクトリに移動し、以下のコマンドを実行してください。

```bash
./count-token < input.txt
263
```

`input.txt` はトークン数を数えたいテキストファイルです。

標準入力から直接テキストを入力することもできます。

```bash
echo "Hello, world!" | ./count-token
4
```
