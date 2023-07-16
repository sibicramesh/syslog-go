package common

import "testing"

func TestCalculatePriority(t *testing.T) {
	type args struct {
		facility int
		severity int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid facility and severity",
			args: args{
				23,
				7,
			},
			want:    191,
			wantErr: false,
		},
		{
			name: "valid facility and severity 0",
			args: args{
				0,
				0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "invalid facility and valid severity",
			args: args{
				25,
				7,
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "valid facility and invalid severity",
			args: args{
				23,
				9,
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "invalid facility and invalid severity",
			args: args{
				26,
				9,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculatePriority(tt.args.facility, tt.args.severity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculatePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
