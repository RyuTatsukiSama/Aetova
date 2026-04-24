package api_download

import (
	"context"
	"encoding/binary"
	"errors"
	"os"
	"sync/atomic"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

type ChunkBitmap struct {
	words []atomic.Uint64
	total int
}

func NewChunkBitmap(chunkCount int) *(ChunkBitmap) {
	// +63 for round to superior unit64
	wordCount := (chunkCount + 63) / 64
	return &ChunkBitmap{
		words: make([]atomic.Uint64, wordCount),
		total: chunkCount,
	}
}

func (b *ChunkBitmap) Mark(chunkID int) {
	word := chunkID / 64
	mask := uint64(1) << (chunkID % 64)
	b.words[word].Or(mask)
}

func (b *ChunkBitmap) IsSet(chunkID int) bool {
	word := chunkID / 64
	mask := uint64(1) << (chunkID % 64)
	return b.words[word].Load()&mask != 0
}

func (b *ChunkBitmap) LoadFromDisk(path string) error {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}

	for i := range b.words {
		offset := i * 8
		if offset+8 > len(data) {
			break
		}
		val := binary.LittleEndian.Uint64(data[offset : offset+8])
		b.words[i].Store(val)
	}
	return nil
}

func (b *ChunkBitmap) SaveToDisk(path string) error {
	buf := make([]byte, len(b.words)*8)
	for i := range b.words {
		val := b.words[i].Load()
		binary.LittleEndian.PutUint64(buf[i*8:], val)
	}
	return os.WriteFile(path, buf, 0644)

	// Code to avoid corruption when crash
	/*tmp := path + ".tmp"
	err := os.WriteFile(tmp, buf, 0644)
	if err != nil {
		return err
	}
	return os.Rename(tmp, path)*/
}

func (b *ChunkBitmap) MissingChunks() []int {
	var missing []int
	for i := range b.total {
		if !b.IsSet(i) {
			missing = append(missing, i)
		}
	}
	return missing
}

func (b *ChunkBitmap) StartAutoSave(ctx context.Context, path string, interval time.Duration) {
	dLog := docLogger.NewLoggerWithGOpts("Client/ChunkBitmap")
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				err := b.SaveToDisk(path)
				if err != nil {
					dLog.Error("Error 12: " + err.Error())
				}
				return
			case <-ticker.C:
				err := b.SaveToDisk(path)
				if err != nil {
					dLog.Error("Error 12: " + err.Error())
				}
			}
		}
	}()
}

func (b *ChunkBitmap) IsComplete() bool {
	fullWords := b.total / 64
	for i := range fullWords {
		if b.words[i].Load() != ^uint64(0) { // ^uint64(0) == all bits to 1
			return false
		}
	}

	// Check remaining bit in the last word
	remaining := b.total % 64
	if remaining > 0 {
		mask := (uint64(1) << uint64(remaining)) - 1
		if b.words[fullWords].Load()&mask != mask {
			return false
		}
	}
	return true
}
