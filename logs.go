package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	//"strings"
	"time"
)

type configuration struct {
	Path          string `json:"Path"`
	MaxSize       uint64 `json:"MaxSize"`
	Level         int32  `json:"Level"`
	Ldate         bool   `json:"Ldate"`
	Ltime         bool   `json:"Ltime"`
	Lmicroseconds bool   `json:"Lmicroseconds"`
	LUTC          bool   `json:"LUTC"`
	LogToFile     bool   `json:"LogToFile"`
	LogToStdErr   bool   `json:"LogToStdErr"`
	Total         int64
}

const (
	info int32 = iota
	warn
	erro
	fata
)

var (
	program = filepath.Base(os.Args[0]) // program name

	conf     configuration
	flog     *log.Logger
	now      time.Time
	curFile  *os.File
	fileinfo *os.FileInfo
)
/*
func loadConf() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Error occur when open conf.json")
	}
	conf.Path = filepath.Join(conf.Path, program)
	conf.MaxSize *= 1024 * 1024
	fmt.Println(conf)
} */
func loadConf() {

	conf.Path = filepath.Join("/var/log/opensds", strings.Split(program, ".")[0])
	conf.Level = -1
	conf.Ldate = true
	conf.Lmicroseconds = true
	conf.LogToFile = true
	conf.LogToStdErr = true
	conf.MaxSize = 1
}
func exits(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func mkdir() {
	if exits(conf.Path) {
		return
	}
	e := os.Mkdir(conf.Path, os.ModePerm)
	if e != nil {
		Fatal(e)
	}
}
func init() {
	loadConf()
	mkdir()
	var tmp int = 0
	if conf.Ldate {
		tmp |= log.Ldate
	}
	if conf.Ltime {
		tmp |= log.Ltime
	}
	if conf.Lmicroseconds {
		tmp |= log.Lmicroseconds
	}
	if conf.LUTC {
		tmp |= log.LUTC
	}
	log.SetFlags(tmp)
	curFile, fileinfo = initFile(conf.Path)
}
func initFile(path string) (*os.File, *os.FileInfo) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		Error(err)
		return nil, nil
	}
	l := len(files)
	if l == 0 {
		return nil, nil
	}
	if uint64(files[l-1].Size()) >= conf.MaxSize {
		return nil, nil
	}
	tmpfile, e1 := os.OpenFile(filepath.Join(conf.Path, files[l-1].Name()), os.O_WRONLY|os.O_APPEND, 0666)
	if e1 != nil {
		return nil, nil
	}
	tmpinfo, e2 := os.Stat(tmpfile.Name())
	if e2 != nil {
		return nil, nil
	}
	return tmpfile, &tmpinfo
}
func open() (*os.File, *os.FileInfo) {
	file, err := os.OpenFile(filepath.Join(conf.Path, (*fileinfo).Name()), os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error("can not open log file :", err)
		return nil, nil
	}
    file.WriteString(fmt.Sprintf("Log file Append at: %v\n\n", time.Now().Format("2018-01-02 15:04:05.000000")))
	tmpinfo, _ := os.Stat(file.Name())
	return file, &tmpinfo
}
func create() (*os.File, *os.FileInfo) {
	name := fmt.Sprintf("%s %04d-%02d-%02d %02d.%02d.%02d.log",
		filepath.Join(conf.Path, program),
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second())

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("create file failed %v\n", err)
		return nil, nil
	}
	file.WriteString(fmt.Sprintf("Log file Create at: %v\n\n", time.Now().Format("2018-01-02 15:04:05.000000")))
	tmpinfo, e := os.Lstat(name)
	if e != nil {
		Error("get fileinfo failed ", e)
	}
	return file, &tmpinfo
}

//
//
func doPrint(s string) {

	if fileinfo == nil || (uint64)((*fileinfo).Size()) >= conf.MaxSize {
		curFile, fileinfo = create()

	} else {
		curFile, fileinfo = open()
	}
	flog = log.New(curFile, "", log.Flags())
	flog.Println(s)
}
func doInfo(v string) {
	if info < conf.Level {
		return
	}
	if conf.LogToStdErr {
		log.Println(v)
	}
	if conf.LogToFile {
		doPrint(v)
	}
}
func Info(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Info]: ", file, " ", link, v)
	doInfo(s)
}
func Infof(format string, v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Info]: ", file, " ", link, v)
	doInfo(s)
}
func Infoln(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprintln("[Info]:", file, link, v)
	doInfo(s)
}
func doWarn(v string) {
	if warn < conf.Level {
		return
	}
	if conf.LogToStdErr {
		log.Println(v)
	}
	if conf.LogToFile {
		doPrint(v)
	}
}

func Warning(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Warn]: ", file, " ", link, v)
	doWarn(s)
}
func Warningf(format string, v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Warn]: ", file, " ", link, v)
	doWarn(s)
}
func Warningln(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprintln("[Warn]:", file, link, v)
	doWarn(s)
}
func doError(s string) {
	if erro < conf.Level {
		return
	}
	if conf.LogToStdErr {
		log.Println(s)
	}
	if conf.LogToFile {
		doPrint(s)
	}
}

func Error(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Erro]: ", file, " ", link, v)
	doError(s)
}
func Errorf(format string, v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Erro]: ", file, " ", link, v)
	doError(s)
}
func Errorln(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprintln("[Erro]:", file, link, v)
	doError(s)
}
func doFatal(s string) {
	if fata < conf.Level {
		return
	}
	if conf.LogToStdErr {
		log.Println(s)
	}
	if conf.LogToFile {
		doPrint(s)
	}
	os.Exit(1)
}
func Fatal(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Fata]: ", file, " ", link, v)
	doFatal(s)
}
func Fatalf(format string, v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprint("[Fata]: ", file, " ", link, v)
	doFatal(s)
}
func Fatalln(v ...interface{}) {
	_, file, link, _ := runtime.Caller(1)
	s := fmt.Sprintln("[Fata]:", file, link, v)
	doFatal(s)
}

func FlushLogs() {

}

type Verbose bool

func V(level int32) Verbose {
	if level >= conf.Level {
		return Verbose(true)
	}
	return Verbose(false)
}
func (v Verbose) Info(args ...interface{}) {
	if v {
		_, file, link, _ := runtime.Caller(1)
		s := fmt.Sprint("[Info]: ", file, " ", link, args)
		doInfo(s)
	}
}
func (v Verbose) Infof(format string, args ...interface{}) {
	if v {
		_, file, link, _ := runtime.Caller(1)
		s := fmt.Sprint("[Info]: ", file, " ", link, args)
		doInfo(s)
	}
}
func (v Verbose) Infoln(args ...interface{}) {
	if v {
		_, file, link, _ := runtime.Caller(1)
		s := fmt.Sprintln("[Info]: ", file, " ", link, args)
		doInfo(s)
	}
}

