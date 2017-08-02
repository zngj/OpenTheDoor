package log4g

import (
	"fmt"
	"os"
	"log"
	"strings"
	"time"
	"path/filepath"
	"strconv"
	"bytes"
	"io"
)

func lineCounter(filename string) int {
	file, err := os.OpenFile(filename, os.O_RDONLY | os.O_CREATE, 0660)
	if err != nil && !os.IsNotExist(err) {
		return 0
	}
	defer file.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		if err == nil {
			count += bytes.Count(buf[:c], lineSep)
		} else if err == io.EOF {
			return count
		} else {
			return 0
		}
	}
}

func newFileLogger(prefix string, flag int, filename string, maxlines int, maxsize int64, maxcount int, daily bool) Logger {

	os.MkdirAll(filepath.Dir(filename), os.ModePerm)

	fileLogger := new (FileLogger)
	fileLogger.filename = filename
	fileLogger.filedir = filepath.Dir(filename)
	fileLogger.maxlines = maxlines
	fileLogger.maxsize = maxsize * 1024 * 1024 * 1024
	fileLogger.maxcount = maxcount
	fileLogger.format = "%s.%0" + strconv.Itoa(len(strconv.Itoa(maxcount-1))) +  "d"
	fileLogger.daily = daily
	fileLogger.lines  = lineCounter(filename)

	output, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Print(err)
		return nil
	}
	fileLogger.file = output
	info, err := output.Stat()
	if err != nil {
		log.Print(err)
		return nil
	}

	fileLogger.size = info.Size()
	fileLogger.lastTime = info.ModTime()
	//for test
	//fileLogger.lastTime = info.ModTime().Add(- 24 * time.Hour)

	filepath.Walk(fileLogger.filedir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(filepath.ToSlash(path), filepath.ToSlash(filename)) {
			fileLogger.count++
		}
		return nil
	})

	fileLogger.GenericLogger = newLogger(prefix, flag, output)

	return fileLogger
}

type FileLogger struct {
	*GenericLogger
	filename string
	filedir string
	file     *os.File
	maxlines int
	maxsize  int64
	maxcount int
	daily    bool
	lines    int
	size     int64
	count    int
	format   string
	lastTime time.Time
}

func (l *FileLogger) BeforeLog() {
	if l.daily {
		l.dailyBackup()
	}
}

func (l *FileLogger) dailyBackup() {
	l.now = time.Now()
	if !l.lastTime.IsZero() {
		ltYear, ltMonth, ltDay := l.lastTime.Date()
		nowYear, nowMonth, nowDay := l.now.Date()
		if ltDay != nowDay || ltMonth != nowMonth || ltYear != nowYear {

			l.mu.Lock()
			defer l.mu.Unlock()

			strDate := fmt.Sprintf("%d%02d%02d", ltYear, ltMonth, ltDay)
			dateDir := filepath.Join(l.filedir, strDate)
			err := os.MkdirAll(dateDir, os.ModePerm)
			if err == nil {
				l.Close()
				//move all file to date director
				err = filepath.Walk(l.filedir, func(path string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}
					if strings.HasPrefix(filepath.ToSlash(path), filepath.ToSlash(l.filename)) {
						os.Remove(filepath.Join(l.filedir, strDate, info.Name()))
						return os.Rename(filepath.Join(l.filedir, info.Name()), filepath.Join(l.filedir, strDate, info.Name()))
					}
					return nil
				})
				if err != nil {
					Error(err)
					return
				}
				l.count = 0
				l.newOutput()
			} else {
				Error(err)
			}
		}
	}
}


func (l *FileLogger) newOutput() {
	//create new log file
	output, err := os.OpenFile(l.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		l.stop = true
	}
	l.file = output
	l.out = output
	l.lines = 0
	l.size = 0
	l.count++
}

func (l *FileLogger) AfterLog(n int) {

	if n <= 0 {
		return
	}

	l.lastTime = l.now

	l.mu.Lock()
	defer l.mu.Unlock()

	l.lines++
	l.size += int64(n)
	if (l.maxlines > 0 && l.lines >= l.maxlines) || (l.maxsize > 0 && l.size > l.maxsize) {

		log.Printf("lines=%d,maxlines=%d,count=%d", l.lines, l.maxlines, l.count)

		//close log file
		l.Close()

		//remove the oldest log
		if l.count == l.maxcount {
			if os.Remove(fmt.Sprintf(l.format, l.filename, l.maxcount-1)) != nil {
				l.stop = true
				return
			}
			l.count--
		}

		//try to rename log files
		var err error
		for i := l.count; i > 0; i-- {
			var oldpath string
			if i == 1 {
				oldpath = l.filename
			} else {
				oldpath = fmt.Sprintf(l.format, l.filename, i-1)
			}
			newpath := fmt.Sprintf(l.format, l.filename, i)
			err = os.Rename(oldpath, newpath)
			if err != nil {
				log.Println(err)
				l.stop = true
				return
			}
		}

		l.newOutput()

	}
}

func (l *FileLogger) Close() {
	l.file.Close()
}
