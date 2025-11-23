package handlers

import (
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		function func(t *testing.T)
	}{
		{
			name: "TestHandlersss",
			function: func(t *testing.T) {
				fmt.Println("Handlers test executed")
			},
		},
		{
			name: "TestAnotherHandler",
			function: func(t *testing.T) {
				fmt.Println("Another handler test executed")
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, tt.function)
	}
}
