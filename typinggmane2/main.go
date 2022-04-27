package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()
	files := flag.Args()

	exe, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "指定できるファイル数は1つだけです")
		os.Exit(0)
	}

	file_path := filepath.Join(exe, files[0])

	words, err := openfile_return_array_char(file_path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		os.Exit(0)
	}

	ch_rcv := myinput(os.Stdin)
	fmt.Println("タイピングゲームを始めます。制限時間は10秒。1語一点", len(words), "点満点")
	score := 0
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Print(question)
		fmt.Print(">")
		select {
		case x := <-ch_rcv:
			if question == x {
				score++
			}
		case <-time.After(10 * time.Second):
			fmt.Println("制限時間を過ぎました")
			i = false
		}
	}
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

func myinput(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
	}()
	return ch1
}
