package meta

import (
	"crypto/md5"
	"fmt"
	"sync"

	np "github.com/zero-os/0-fs/cap.np"
)

//File represents a file inode
type File struct {
	np.Inode
	file   np.File
	access Access

	blks []BlockInfo
	o    sync.Once
}

//ID returns file ID
func (f *File) ID() string {
	m := md5.New()
	for _, blk := range f.Blocks() {
		m.Write(blk.Key)
	}
	return fmt.Sprintf("%x", m.Sum(nil))
}

//Name return file name
func (f *File) Name() string {
	name, _ := f.Inode.Name()
	return name
}

//IsDir false for files
func (f *File) IsDir() bool {
	return false
}

//Children nil for files
func (f *File) Children() []Meta {
	return nil
}

//Info return meta info for this dir
func (f *File) Info() MetaInfo {
	return MetaInfo{
		CreationTime:     f.CreationTime(),
		ModificationTime: f.ModificationTime(),
		Size:             f.Size(),
		Type:             RegularType,
		Access:           f.access,
		FileBlockSize:    uint64(f.file.BlockSize()) * 4096,
	}
}

func (f *File) blocks() {
	var blocks []BlockInfo
	if !f.file.HasBlocks() {
		return
	}

	cblocks, _ := f.file.Blocks()
	for i := 0; i < cblocks.Len(); i++ {
		block := cblocks.At(i)

		hash, _ := block.Hash()
		key, _ := block.Key()
		blocks = append(blocks, BlockInfo{
			Key:      hash,
			Decipher: key,
		})
	}

	f.blks = blocks
}

//Blocks loads and return blocks of file
func (f *File) Blocks() []BlockInfo {
	f.o.Do(f.blocks)
	return f.blks
}
