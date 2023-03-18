package rename

import (
	"testing"

	"github.com/victorbischoff/fileutilities/pkg/utils"
)

func TestRenameFilesInDirectory(t *testing.T) {
	type args struct {
		cleanPath string
		prefix    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE 1",
			args: args{
				cleanPath: "./.././../filesfortesting",
				prefix:    utils.RandomPrefix(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenameFilesInDirectory(tt.args.cleanPath, tt.args.prefix)
		})
	}
}
