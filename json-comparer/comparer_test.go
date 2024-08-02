package jsoncomparer

import (
	"encoding/json"
	"testing"
)

func parseJSON(str string) interface{} {
	var content interface{}
	_ = json.Unmarshal([]byte(str), &content)
	return content
}

func TestCompare(t *testing.T) {
	type args struct {
		expected interface{}
		actual   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "bool",
			args:    args{true, false},
			wantErr: true,
		},
		{
			name: "bool",
			args: args{true, true},
		},
		{
			name:    "number",
			args:    args{1, 2},
			wantErr: true,
		},
		{
			name: "number",
			args: args{1, 1},
		},
		{
			name:    "string",
			args:    args{"happy new year", "no lah"},
			wantErr: true,
		},
		{
			name: "string",
			args: args{"happy new year", "happy new year"},
		},
		{
			name: "a map",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
			},
		},
		{
			name: "actual is not a map",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1,
					"name": []
				}`),
			},
			wantErr: true,
		},
		{
			name: "actual map missing key",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1
				}`),
			},
			wantErr: true,
		},
		{
			name: "a map contains slice",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1
				}`),
			},
			wantErr: true,
		},
		{
			name: "a map contains slice but the actual is not a slice",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": {"name": "vuong"}
					}
				}`),
			},
			wantErr: true,
		},
		{
			name: "a map contains slice but the actual missing item at expected index",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20}
						]
					}
				}`),
			},
			wantErr: true,
		},
		{
			name: "actual slice is mismatched at expected index",
			args: args{
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 20}
						]
					}
				}`),
				parseJSON(`{
					"abc": 1,
					"name": {
						"huy": "khila",
						"name": ["a", "b", "c"],
						"student": [
							{"name": "Binh", "age": 20},
							{"name": "An", "age": 21}
						]
					}
				}`),
			},
			wantErr: true,
		},
		{
			name: "slice",
			args: args{
				parseJSON(`[
					{"name": "Binh", "age": 20},
					{"name": "An", "age": 20}
				]`),
				parseJSON(`[
					{"name": "Binh", "age": 20},
					{"name": "An", "age": 21}
				]`),
			},
			wantErr: true,
		},
		{
			name: "slice",
			args: args{
				parseJSON(`[
					{"name": "Binh", "age": 20},
					{"name": "An", "age": 20}
				]`),
				parseJSON(`[
					{"name": "Binh", "age": 20},
					{"name": "An", "age": 20}
				]`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compare(tt.args.expected, tt.args.actual); (err != nil) != tt.wantErr {
				t.Errorf("Comparer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
