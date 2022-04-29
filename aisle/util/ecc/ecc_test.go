/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 17:19
 */

package ecc

import (
	"reflect"
	"testing"
)

func TestConvertToPubKey(t *testing.T) {
	type args struct {
		pub string
	}
	tests := []struct {
		name    string
		args    args
		want    Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToPubKey(tt.args.pub)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToPubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToPubKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name string
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_ConvertToString(t *testing.T) {
	type fields struct {
		key T
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				key: tt.fields.key,
			}
			if got := key.ConvertToString(); got != tt.want {
				t.Errorf("ConvertToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Decrypt(t *testing.T) {
	type fields struct {
		key T
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantDData []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				key: tt.fields.key,
			}
			gotDData, err := key.Decrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDData, tt.wantDData) {
				t.Errorf("Decrypt() gotDData = %v, want %v", gotDData, tt.wantDData)
			}
		})
	}
}

func TestKey_Encrypt(t *testing.T) {
	type fields struct {
		key T
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantEData []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				key: tt.fields.key,
			}
			gotEData, err := key.Encrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEData, tt.wantEData) {
				t.Errorf("Encrypt() gotEData = %v, want %v", gotEData, tt.wantEData)
			}
		})
	}
}

func TestKey_GeneratePublicKey(t *testing.T) {
	type fields struct {
		key T
	}
	tests := []struct {
		name   string
		fields fields
		want   Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				key: tt.fields.key,
			}
			if got := key.GeneratePublicKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
