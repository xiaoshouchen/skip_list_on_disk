package sorted_string_table

import (
	"fmt"
	"strconv"

	"github.com/xiaoshouchen/skip-list/disk_io"
	"github.com/xiaoshouchen/skip-list/skip_list"
)

type SortedStringTableI interface {
	Get(key string) (string, error)
	Insert(key, value string) error
	Remove(key string) error
	Size() uint
}

type SortedStringTable struct {
	skipList skip_list.SkipListI
	io       disk_io.DiskI
}

func NewSST() SortedStringTableI {
	d := disk_io.NewDisk()
	return &SortedStringTable{
		skipList: skip_list.NewSkipList(10),
		io:       d,
	}
}

func (sst *SortedStringTable) Get(key string) (string, error) {
	offset, err := sst.skipList.Get(key)
	if err != nil {
		return "", skip_list.NotFoundErr
	}
	//从磁盘里进行读取数据
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}
	resStr, err := sst.io.ReadAt(uint64(offsetInt))
	if err != nil {
		return "", err
	}
	return resStr, nil
}

func (sst *SortedStringTable) Insert(key, value string) error {
	//数据写入磁盘
	offset, err := sst.io.Write(value)
	if err != nil {
		return err
	}

	//offset写入链表
	sst.skipList.Insert(key, fmt.Sprintf("%d", offset))
	return nil
}

func (sst *SortedStringTable) Remove(key string) error {
	return sst.skipList.Remove(key)
}

func (sst *SortedStringTable) Size() uint {
	return sst.skipList.Size()
}
