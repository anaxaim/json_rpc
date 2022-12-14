package main

import (
	"net/http"
	"testing"
)

func TestGreetingService_Greeting(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *GreetingArgs
		reply *string
	}
	tests := []struct {
		name    string
		s       *GreetingService
		args    args
		wantErr bool
	}{
		{
			name:    "EmptyName",
			args:    args{args: &GreetingArgs{Name: ""}},
			wantErr: true,
		},
		{
			name:    "CorrectName",
			s:       new(GreetingService),
			args:    args{args: &GreetingArgs{Name: "Vova"}, reply: new(string)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GreetingService{}
			if err := s.Greeting(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("GreetingService.Greeting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
