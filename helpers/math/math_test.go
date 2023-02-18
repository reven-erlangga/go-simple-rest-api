package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative to absolute",
			args: args{
				num: -1,
			},
			want: 1,
		},
		{
			name: "absolute to absolute",
			args: args{
				num: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Absolute(tt.args.num)
			assert.Equal(t, tt.want, c)
		})
	}
}


func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative and negative",
			args: args{
				a: -1,
				b: -1,
			},
			want: -2,
		},
		{
			name: "negative and positive",
			args: args{
				a: -3,
				b: 1,
			},
			want: -2,
		},
		{
			name: "positive and positive",
			args: args{
				a: 4,
				b: 1,
			},
			want: 5,
		},
		{
			name: "positive and negative",
			args: args{
				a: 3,
				b: -1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Add(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, c)
		})
	}
}