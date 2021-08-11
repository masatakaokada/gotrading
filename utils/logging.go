package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//引数: ファイルのパス, フラグ, パーミッション
	// os.O_RDWR 読み書き両方する時
	// os.O_CREATE ファイルがなかったら作成
	// os.O_APPEND 追記できるように
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}
	// Stdoutとlogfile両方にログ出力する
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
