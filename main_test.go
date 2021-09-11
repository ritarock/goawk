package main

import (
	"reflect"
	"testing"
)

func Test_read(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "read file",
			args:       args{file: "./test.txt"},
			wantResult: []string{"1 2 3 4 5", "6 7 8 9 10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := read(tt.args.file); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("read() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_parseCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "parse command",
			args: args{command: "{print $1}"},
			want: 0,
		},
		{
			name:    "error command",
			args:    args{command: "print $1}"},
			wantErr: true,
		},
		{
			name:    "error command",
			args:    args{command: "{prin $1}"},
			wantErr: true,
		},
		{
			name:    "error command",
			args:    args{command: "{print 1}"},
			wantErr: true,
		},
		{
			name:    "error command",
			args:    args{command: "{print $q}"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseField(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "parse field",
			args: args{field: "[:]"},
			want: ":",
		},
		{
			name: "parse field",
			args: args{field: "[/]"},
			want: "/",
		},
		{
			name:    "error command",
			args:    args{field: ":"},
			wantErr: true,
		},
		{
			name:    "error command",
			args:    args{field: "[:"},
			wantErr: true,
		},
		{
			name: "error command",
			args: args{field: "[]"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseField(tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseField() = %v, want %v", got, tt.want)
			}
		})
	}
}
