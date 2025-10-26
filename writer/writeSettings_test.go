package writer

import (
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"testing"
)

func TestWriteToJSON(t *testing.T) {
	type args struct {
		objectSchema *data.SyntheticSettings
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteToJSON(tt.args.objectSchema); (err != nil) != tt.wantErr {
				t.Errorf("WriteToJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
