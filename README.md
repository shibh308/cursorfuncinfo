# cursorfuncinfo
ファイル名とカーソル位置(オフセット)から対象の関数を取得し、テスト関数ならテスト先の関数を、それ以外の関数ならテスト元の関数のデータ(パッケージ、ファイルパス、オフセット等)をjson形式で返します

## Install
```sh
$ go get github.com/shibh308/cursorfuncinfo
```

## Flags
`-f`, `--file`: filepath  
`-p`, `--pos`: cursor offset

## Example
```sh
$ cursorfuncinfo -f run/run_test.go -p 300
```
