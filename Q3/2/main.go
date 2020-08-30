/*
Q3.2: テスト用の適当なサイズのファイルを作成
　ファイルを作成してランダムな内容で埋めてみましょう。
crypto/randパッケージ（本来は付属Aで紹介するように暗号用の機能）
をインポートすると、rand.Readerというio.Readerが使えます。
このReaderは、ランダムなバイトを延々と出力し続ける
無限長のファイルのような動作をします。
これを使って、1024バイトの長さの
バイナリファイルを作ってみましょう。
ヒントですが、io.Copy()を使ってはいけません。
io.Copy()はReaderの終了まですべて愚直にコピー使用とします。
そしてrand.Readerには終わりはありません。
あとはわかりますよね？
*/

package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	generateRandBinary1024()
	fmt.Println("DONE")
}

func generateRandBinary1024() {
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	var rb1024 []byte
	rb1024 = make([]byte, 1024)

	reader1024 := io.LimitReader(rand.Reader, 1024)
	reader1024.Read(rb1024)

	_, err = newFile.Write(rb1024)
	if err != nil {
		panic(err)
	}
}
