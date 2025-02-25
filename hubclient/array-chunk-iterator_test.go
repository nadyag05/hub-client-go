package hubclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayChunkIterator_hasNext(t *testing.T) {
	type fields struct {
		chunks   []string
		position int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test has next",
			fields: fields{
				chunks:   []string{"chunk1", "chunk2"},
				position: 0,
			},
			want: true,
		},
		{
			name: "Test has no next",
			fields: fields{
				chunks:   []string{"chunk1", "chunk2"},
				position: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ArrayChunkIterator{
				chunks:   tt.fields.chunks,
				position: tt.fields.position,
			}
			assert.Equalf(t, tt.want, i.hasNext(), "hasNext()")
		})
	}
}

func TestArrayChunkIterator_next(t *testing.T) {
	type fields struct {
		chunks   []string
		position int
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test get existing next",
			fields: fields{
				chunks:   []string{"chunk1", "chunk2"},
				position: 0,
			},
			want:    "chunk2",
			wantErr: assert.NoError,
		},
		{
			name: "Test fail when there is no next",
			fields: fields{
				chunks:   []string{"chunk1", "chunk2"},
				position: 1,
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ArrayChunkIterator{
				chunks:   tt.fields.chunks,
				position: tt.fields.position,
			}
			got, err := i.next()
			if !tt.wantErr(t, err, fmt.Sprintf("next()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "next()")
		})
	}
}

func TestNewArrayChunkIterator(t *testing.T) {
	type args struct {
		chunks []string
	}
	tests := []struct {
		name string
		args args
		want ArrayChunkIterator
	}{
		{
			name: "Create new ArrayChunkIterator",
			args: args{
				chunks: []string{"chunk1", "chunk2"},
			},
			want: ArrayChunkIterator{chunks: []string{"chunk1", "chunk2"}, position: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewArrayChunkIterator(tt.args.chunks), "NewArrayChunkIterator(%v)", tt.args.chunks)
		})
	}
}
