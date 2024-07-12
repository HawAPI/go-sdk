package hawapi

import (
	"testing"
)

func TestClient_buildUrl(t *testing.T) {
	type fields struct {
		options Options
	}
	type args struct {
		origin string
		query  []QueryOptions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "should build a simple url",
			fields: fields{},
			args: args{
				origin: "actors",
			},
			want: "https://hawapi.theproject.id/api/v1/actors",
		},
		{
			name: "should build url with custom endpoint and version",
			fields: fields{
				options: Options{
					Endpoint: "https://hawapi.example.id/api",
					Version:  "v3",
				},
			},
			args: args{
				origin: "actors",
			},
			want: "https://hawapi.example.id/api/v3/actors",
		},
		{
			name: "should build url with global options",
			fields: fields{
				options: Options{
					Language: "pt-BR",
					Size:     20,
				},
			},
			args: args{
				origin: "actors",
				query:  []QueryOptions{},
			},
			want: "https://hawapi.theproject.id/api/v1/actors?language=pt-BR&size=20",
		},
		{
			name:   "should build url with pageable",
			fields: fields{},
			args: args{
				origin: "actors",
				query: []QueryOptions{
					WithPage(2),
					WithSize(40),
				},
			},
			want: "https://hawapi.theproject.id/api/v1/actors?page=2&size=40",
		},
		{
			name:   "should build url with sort",
			fields: fields{},
			args: args{
				origin: "actors",
				query: []QueryOptions{
					WithSort("first_name"),
					WithOrder("DESC"),
				},
			},
			want: "https://hawapi.theproject.id/api/v1/actors?sort=first_name,DESC",
		},
		{
			name:   "should build ignore order if sort is not present",
			fields: fields{},
			args: args{
				origin: "actors",
				query: []QueryOptions{
					WithOrder("DESC"),
				},
			},
			want: "https://hawapi.theproject.id/api/v1/actors",
		},
		{
			name:   "should build overwrite filter if is already set",
			fields: fields{},
			args: args{
				origin: "actors",
				query: []QueryOptions{
					WithFilter("gender", "1"),
					WithFilter("first_name", "Finn"),
					WithFilter("gender", "0"),
				},
			},
			want: "https://hawapi.theproject.id/api/v1/actors?gender=0&first_name=Finn",
		},
		{
			name:   "should build a complete url",
			fields: fields{},
			args: args{
				origin: "actors",
				query: []QueryOptions{
					WithLanguage("fr-FR"),
					WithSize(20),
					WithFilter("gender", "1"),
					WithSort("first_name"),
					WithOrder("DESC"),
				},
			},
			want: "https://hawapi.theproject.id/api/v1/actors?language=fr-FR&gender=1&size=20&sort=first_name,DESC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithOpts(tt.fields.options)

			if got := c.buildUrl(tt.args.origin, tt.args.query); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
