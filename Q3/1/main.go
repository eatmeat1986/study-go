/*
Q3.1: ファイルのコピー
古いファイル(old.txt)を新しいファイル(new.txt)にコピーしてみましょう。
本章で紹介したサンプルコードを応用すれば難しくないと思います。
　さらに改造して実用的なコマンドにしてみたいと思われる方は、
コマンドラインオプションでファイル名を渡せるようにするとよいでしょう。
本書の範囲からは外れるので、省きますが、os.Argsという文字列配列にオプションが格納されます。
また、標準ライブラリにあるFlagパッケージを使うと、オプションのパース処理がより便利に行えます。
*/

package main

import (
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, oldFile)
}
