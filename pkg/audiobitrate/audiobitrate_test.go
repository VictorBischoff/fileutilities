package audiobitrate

import (
	"testing"
)

func TestConvert16to24Wav(t *testing.T) {
	type args struct {
		path    string
		pathOut string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CASE 1",
			args: args{
				path:    "../../audiofortesting",
				pathOut: "../../audiofortesting/1-test-32bit.wav",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert16to24Wav(tt.args.path, tt.args.pathOut); (err != nil) != tt.wantErr {
				t.Errorf("Convert16to24() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvertFolder24to16Wav(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CASE 1",
			args: args{
				path: "../../audiofortesting",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertFolder24to16Wav(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ConvertFolder24to16Wav() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
