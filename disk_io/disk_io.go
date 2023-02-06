package disk_io

import (
	"bufio"
	"os"
	"strings"
)

const BasePath = "tmp/"

type Disk struct {
	offset uint64
}

type DiskI interface {
	ReadAt(offset uint64) (string, error)
	Write(value string) (uint64, error)
}

func NewDisk() DiskI {
	return &Disk{
		offset: func() uint64 {
			filepath := BasePath + "value.log"
			if _, err := os.Stat(BasePath); os.IsNotExist(err) {
				return 0
			}
			file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				return 0
			}
			defer file.Close()

			buf := bufio.NewReader(file)
			var i = uint64(0)
			for {
				_, err := buf.ReadString('\n')

				i++
				if err != nil {
					break
				}
			}
			return i
		}(),
	}
}

func (d *Disk) ReadAt(offset uint64) (string, error) {
	filepath := BasePath + "value.log"
	if _, err := os.Stat(BasePath); os.IsNotExist(err) {
		os.Mkdir(BasePath, 0777)
	}
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	var i = uint64(1)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if i == offset {
			return line, nil
		}
		i++
		if err != nil {
			break
		}
	}
	return "", nil
}

func (d *Disk) Write(value string) (uint64, error) {
	if _, err := os.Stat(BasePath); os.IsNotExist(err) {
		os.Mkdir(BasePath, 0777)
	}

	filepath := BasePath + "value.log"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	value = "\n" + strings.TrimSpace(value)
	if d.offset == 0 {
		value = strings.TrimSpace(value)
	}
	_, err = buf.WriteString(value)
	if err != nil {
		return 0, err
	}
	err = buf.Flush()
	if err != nil {
		return 0, err
	}
	d.offset++
	return d.offset, nil
}
