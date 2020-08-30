/*
Q3.3: zipファイルの書き込み
　OSのデバイスにリンクされたio.Writerやio.Readerは、
1つのファイルやデバイスと1対1に対応しています。
Go言語が提供するライブラリには、1つのファイルで
複数のio.Writerやio.Readerの仲間で構成されているものもあります。
複数ファイルを格納するアーカイブフォーマットであるtarやzipファイルや、
インターネットのマルチパート形式(ブラウザのフォームによって作られる
データやファイルを複数格納するデータ構造)をサポートするmime/multipartパッケージ
の構造体は、中に格納されるひとつひとつの要素がio.Writerやio.ReadCloserになっています。
　archive/zipパッケージを使ってzipファイルを作成してみましょう。
出力先のファイルのWriter(以下のコードのfile)をまず作って、
それをzip.NewWriter()関数に渡すと、zipファイルの書き込み用の構造体ができます。
最後にClose()を確実に呼ぶ必要がありますが、これにはGo言語のdeferという機能を使って次の
ようにすればいいでしょう。

この構造体そのものはio.Writerではありませんが、Create()メソッドを呼ぶと
個別のファイルを書き込むためのio.Writerが帰ってきます。

上記の例では、newfile.txtという実際のファイルが、最初に作った出力先の
ファイルfileへと圧縮されます。では、実際のファイルではなく、文字列strings.Reader
を使ってzipファイルを作成するにはどうすればいいでしょうか。考えてください。

*/

package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("aiueo.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(writer, strings.NewReader("aiueo"))
	// ちょっとこの問題意味不明すぎんよ
}
