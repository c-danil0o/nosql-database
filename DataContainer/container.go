package container

type DataNode interface {
	Value() []byte
	Key() []byte
	Timestamp() int64
	Tombstone() byte
	SetValue([]byte)
	SetKey([]byte)
	SetTimestamp(int64)
	SetTombstone(byte)
}

type Container interface {
	Insert(key []byte, value []byte, timestamp int64, tombstone byte)
	Find(key []byte) DataNode
	PrefixScan(key []byte) []DataNode
	RangeScan(minKey []byte, maxKey []byte) []DataNode
	Delete(key []byte) bool
	Size() int
	GetSortedData() []DataNode
	Print()
}
