package commands

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Command
	}{
		{
			name: "nil args",
			args: nil,
			want: UndefinedCommand{Err: fmt.Errorf("too few arguments")},
		},
		{
			name: "empty args",
			args: []string{},
			want: UndefinedCommand{Err: fmt.Errorf("too few arguments")},
		},
		{
			name: "undefined command",
			args: []string{"foo"},
			want: UndefinedCommand{Err: fmt.Errorf("undefined command: [foo]")},
		},
		{
			name: "fetch all lists command",
			args: []string{"list"},
			want: FetchAllListsCommand{},
		},
		{
			name: "fetch selected lists command",
			args: []string{"list", "foo", "bar"},
			want: FetchListsCommand{[]string{"foo", "bar"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.args)

			var opts cmp.Options
			if _, ok := tt.want.(UndefinedCommand); ok {
				opts = append(opts, cmp.Comparer(func(x, y UndefinedCommand) bool {
					return (x.Err == nil && y.Err == nil) || (x.Err.Error() == y.Err.Error())
				}))
			}

			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Parse(%#v) diff: %s", tt.args, diff)
			}
		})
	}
}
