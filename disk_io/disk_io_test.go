package disk_io

import (
	"fmt"
	"testing"
)

func TestDisk_ReadAt(t *testing.T) {
	disk := NewDisk()
	for i := 1; i < 50; i++ {
		str, _ := disk.ReadAt(uint64(i))
		fmt.Println(str)
	}
}

func TestDisk_Write(t *testing.T) {
	disk := NewDisk()

	offset, err := disk.Write(fmt.Sprintf("第%d行", 1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(offset)

}
