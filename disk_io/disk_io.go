package disk_io

type Disk struct {
}

type DiskI interface {
	ReadAt(offset uint64) []byte
	WriteAt(offset uint64, content []byte)
}
