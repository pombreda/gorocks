package gorocks

/*
#cgo LDFLAGS: -lrocksdb
#include "rocksdb/c.h"
*/
import "C"

func GetLevelDBMajorVersion() int {
	return int(C.leveldb_major_version())
}

func GetLevelDBMinorVersion() int {
	return int(C.leveldb_minor_version())
}
