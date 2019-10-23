package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// コマンドライン引数を受け取る変数を宣言する。
	var (
		dir = flag.String("dir", "./", "target dirpath")
	)
	// コマンドライン引数をパースする。
	flag.Parse()
	fmt.Println("dir:" + *dir)
	fmt.Println(dirwalk(*dir))
}

func dirwalk(dir string) []string {

	/*
	ioutil.ReadDirは、指定ディレクトリ内のファイルコンテンツを、
	os.FileInfoのスライスで返却する。
	*/
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		/*
		filepath.Joinは、指定文字列をファイルパス区切で連結する。
		*/
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
