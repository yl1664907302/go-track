package mysql

import (
	"fmt"
	"go-track/pojo"
	"testing"
)

func TestGetReceivers(t *testing.T) {
	receivers, err := GetReceivers()
	fmt.Println(err)
	fmt.Println(receivers)
}

func TestGetReceivers1(t *testing.T) {
	tests := []struct {
		name    string
		want    []pojo.Receiver
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetReceivers()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(got)
		})
	}
}
