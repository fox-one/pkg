package columnize

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestForm_String(t *testing.T) {
	var form Form
	form.Append("  Column A", " Column B", "Column C  ")
	form.Append("x", "y", " z")
	output := form.String()

	expected := "Column A Column B Column C \n"
	expected += "x        y        z        \n"

	assert.Equal(t, expected, output)
}

func Test_formatLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one item",
			args: args{
				line: []string{"a"},
			},
			want: "a\t",
		},
		{
			name: "multiple items",
			args: args{
				line: []string{"a", "b", "c"},
			},
			want: "a\tb\tc\t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatLine(tt.args.line); got != tt.want {
				t.Errorf("formatLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
