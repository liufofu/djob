// Port from TiDB
// Copyright 2015 PingCAP, Inc.
//
// Copyright 2015 Wenbin Xiao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.
package store

import (
	"sync/atomic"

	"github.com/juju/errors"
	"github.com/pingcap/goleveldb/leveldb"
	"github.com/pingcap/goleveldb/leveldb/comparer"
	"github.com/pingcap/goleveldb/leveldb/iterator"
	"github.com/pingcap/goleveldb/leveldb/memdb"
	"github.com/pingcap/goleveldb/leveldb/util"

	derrors "github.com/HZ89/djob/errors"
)

// Those limits is enforced to make sure the transaction can be well handled by TiKV.
var (
	// TxnEntrySizeLimit is limit of single entry size (len(key) + len(value)).
	TxnEntrySizeLimit = 6 * 1024 * 1024
	// TxnEntryCountLimit  is limit of number of entries in the MemBuffer.
	TxnEntryCountLimit uint64 = 300 * 1000
	// TxnTotalSizeLimit is limit of the sum of all entry size.
	TxnTotalSizeLimit = 100 * 1024 * 1024
)

type memDbBuffer struct {
	db              *memdb.DB
	entrySizeLimit  int
	bufferLenLimit  uint64
	bufferSizeLimit int
}

type memDbIter struct {
	iter    iterator.Iterator
	reverse bool
}

type Key []byte

// NewMemDbBuffer creates a new memDbBuffer.
func NewMemDbBuffer() *memDbBuffer {
	return &memDbBuffer{
		db:              memdb.New(comparer.DefaultComparer, 4*1024),
		entrySizeLimit:  TxnEntrySizeLimit,
		bufferLenLimit:  atomic.LoadUint64(&TxnEntryCountLimit),
		bufferSizeLimit: TxnTotalSizeLimit,
	}
}

// Seek creates an Iterator.
func (m *memDbBuffer) Seek(k Key) (*memDbIter, error) {
	var i *memDbIter
	if k == nil {
		i = &memDbIter{iter: m.db.NewIterator(&util.Range{}), reverse: false}
	} else {
		i = &memDbIter{iter: m.db.NewIterator(&util.Range{Start: []byte(k)}), reverse: false}
	}
	i.Next()
	return i, nil
}

func (m *memDbBuffer) SeekReverse(k Key) (*memDbIter, error) {
	var i *memDbIter
	if k == nil {
		i = &memDbIter{iter: m.db.NewIterator(&util.Range{}), reverse: true}
	} else {
		i = &memDbIter{iter: m.db.NewIterator(&util.Range{Limit: []byte(k)}), reverse: true}
	}
	i.iter.Last()
	return i, nil
}

// Get returns the value associated with key.
func (m *memDbBuffer) Get(k Key) ([]byte, error) {
	v, err := m.db.Get(k)
	if err == leveldb.ErrNotFound {
		return nil, derrors.ErrNotExist
	}
	return v, nil
}

// Set associates key with value.
func (m *memDbBuffer) Set(k Key, v []byte) error {
	if len(v) == 0 {
		return derrors.ErrCannotSetNilValue
	}
	if len(k)+len(v) > m.entrySizeLimit {
		return derrors.ErrEntryTooLarge
	}

	err := m.db.Put(k, v)
	if m.Size() > m.bufferSizeLimit {
		return derrors.ErrTxnTooLarge
	}
	if m.Len() > int(m.bufferLenLimit) {
		return derrors.ErrTxnTooLarge
	}
	return errors.Trace(err)
}

// Delete removes the entry from buffer with provided key.
func (m *memDbBuffer) Delete(k Key) error {
	err := m.db.Put(k, nil)
	return errors.Trace(err)
}

// Size returns sum of keys and values length.
func (m *memDbBuffer) Size() int {
	return m.db.Size()
}

// Len returns the number of entries in the DB.
func (m *memDbBuffer) Len() int {
	return m.db.Len()
}

// Next implements the Iterator Next.
func (i *memDbIter) Next() error {
	if i.reverse {
		i.iter.Prev()
	} else {
		i.iter.Next()
	}
	return nil
}

// Valid implements the Iterator Valid.
func (i *memDbIter) Valid() bool {
	return i.iter.Valid()
}

// Key implements the Iterator Key.
func (i *memDbIter) Key() Key {
	return i.iter.Key()
}

// Value implements the Iterator Value.
func (i *memDbIter) Value() []byte {
	return i.iter.Value()
}

// Close Implements the Iterator Close.
func (i *memDbIter) Close() {
	i.iter.Release()
}
