package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files2 := flag.Args()
	flag.Parse()

	exe, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exe)
	path := filepath.Dir(exe)

	fmt.Println(path)
	files := "words.txt"
	fmt.Println(files2)
	fmt.Println(filepath.Join(exe, files))
	words, err := openfile_return_array_char(filepath.Join(exe, files))
	fmt.Println(words)
}

/*
Parameter
	file_path: 対象のファイルまで含めたフルパス
Response
	[]string: ファイルの中身を行ごとに格納した配列
	error: エラー
*/
func openfile_return_array_char(file_path string) ([]string, error) {
	var words []string

	/*
		pathを使用してファイルを開く。正常処理でファイルのポインタをレスポンスし、
		例外を第二引数としてレスポンスする。
		sf = 開いたファイルのポインタ
		err = 処理中の例外
	*/
	sf, err := os.Open(file_path)

	if err != nil {
		return nil, err
	}

	//読み取り対象のポインタを引数として渡して、ioインスタンスを作成≒読み取り準備。
	scanner := bufio.NewScanner(sf)

	/*
		Scan(): インスタンスが読み取りの結果でboolを返す。
		一回の読み取りで次の行に行くため、行が残っている場合はtrueで
		継続処理。読み取りの対象が無ければ、falseで終了
	*/
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, err
}
