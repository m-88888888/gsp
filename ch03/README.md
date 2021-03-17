# ch03

io.Readerのインタフェースを満たすRead()の動き

1. 読み込むデータの格納先を確保（make関数で確保）
2. Read()実行。引数には上記のバッファを指定
3. 実際に読み込んだバイト数とエラーの有無を返す

バッファをいちいち用意するのがめんどう・・・

そこで、入力を扱う補助関数を使用する。

### 補助関数

```jsx
// すべて読み込みたいとき
buffer, err := ioutil.ReadAll(reader)

// 決まったバイト数だけ読み込みたいとき
buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)

// io.Reader→io.Writerにそのままデータをコピーしたいとき
// すべてコピー
writeSize, err := io.Copy(writer, reader)

// 指定したサイズだけコピーしたいとき
writeSize, err := io.CopyN(writer, reader, size)

// バッファを使いまわしたいとき
buffer := make([]byte, 8 * 1024)
io.CopyBuffer(writer, reader, buffer)
```

### その他入出力関連インタフェース

- io.Closer・・・使用後のファイルを閉じる
- io.Seeker・・・読み書き位置の変更
- io.ReaderAt・・・ランダムアクセスできるオブジェクトに自由にアクセス（使いみちなぞすぎる）

GolangのRead()はタイムアウトの仕組みがないから、入力待ちでブロックしてしまう。

だから、goroutineを使ってノンブロッキングな処理にする必要がある。

1. goroutineで軽量スレッド作成
2. スレッドで読み込み

### ファイル入力

異なるインタフェース間（os.Stdoutとfile）で値をコピー

```jsx
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
```

### バイナリ解析

ちょっとムズすぎるのでスキップ

### テキスト解析

```jsx
var source = `１行目
２行目
３行目`

// ①
func main() {
	// bufio.Readerを使ってテキスト分割
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

// ②
func main() {
	// bufio.Scannerでも同様のことができ、コードが短く済む
	// しかしデフォルトでは分割文字が削除される
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

// ③
// データ型を指定
var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("%v, %v, %v, %v", i, f, g, s)
}

`### パイプ（ストリーム）

io.Readerとio.Writer間のデータが流れるパイプ

```jsx
// ①引数のReaderすべて繋げる
func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}

// ②writerに書き出したあともreaderに残す
func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)

	fmt.Println(buffer.String())
}
```

①と②でも使っている`io.Pipe()`は同期処理なので、ReadとWrite両方が終わらないと次の処理が動かない

bufio.NewWriterなどを使ってバッファリングすることで対処可能

チャネルによる並列処理でも同期処理が走るから、ReadとWriteの並行処理を走らせたい場合はgoroutineを使うこと``
