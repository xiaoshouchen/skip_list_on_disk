package sorted_string_table

import (
	"fmt"

	"github.com/xiaoshouchen/skip-list/disk_io"
	"github.com/xiaoshouchen/skip-list/skip_list"
)

type SortedStringTableI interface {
	Get(key string) (string, error)
	Insert(key, value string)
	Remove(key string) error
	Size() uint
}

type SortedStringTable struct {
	skipList skip_list.SkipListI
	io       disk_io.DiskI
}

func NewSST() SortedStringTableI {
	return &SortedStringTable{}
}

func (sst *SortedStringTable) Get(key string) (string, error) {
	offset, err := sst.skipList.Get(key)
	if err != nil {
		return "", skip_list.NotFoundErr
	}
	//从磁盘里进行读取数据
	return "", nil
}

func (sst *SortedStringTable) Insert(key, value string) {
	var offset uint64

	//数据写入磁盘

	sst.skipList.Insert(key, fmt.Sprintf("%d", offset))

}

func (sst *SortedStringTable) Remove(key string) error {
	return nil
}

func (sst *SortedStringTable) Size() uint {
	return sst.skipList.Size()
}
