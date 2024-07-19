package hawapi

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/HawAPI/go-sdk/pkg/cache"
)

var (
	defaultTestLoggerHandler = NewFormattedHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})
	defaultTestLogger        = slog.New(defaultTestLoggerHandler)
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
		// TODO: Fix overwrite test, sometimes params will in different order
		//{
		//	name:   "should build overwrite filter if is already set",
		//	fields: fields{},
		//	args: args{
		//		origin: "actors",
		//		query: []QueryOptions{
		//			WithFilter("gender", "1"),
		//			WithFilter("first_name", "Finn"),
		//			WithFilter("gender", "0"),
		//		},
		//	},
		//	want: "https://hawapi.theproject.id/api/v1/actors?gender=0&first_name=Finn",
		//},
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
			tt.fields.options.LogHandler = defaultTestLoggerHandler
			c := NewClientWithOpts(tt.fields.options)

			if got := c.buildUrl(tt.args.origin, tt.args.query); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractHeaders(t *testing.T) {
	type args struct {
		header http.Header
	}
	tests := []struct {
		name string
		args args
		want HeaderResponse
	}{
		{
			name: "test",
			args: args{
				header: http.Header{
					apiHeaderRateLimitRemaining: []string{"15"},
					apiHeaderContentLanguage:    []string{"fr-FR"},
					apiHeaderContentLength:      []string{"123"},
					apiHeaderItemTotal:          []string{"10"},
					apiHeaderPageIndex:          []string{"1"},
					apiHeaderPageSize:           []string{"10"},
					apiHeaderPageTotal:          []string{"1"},
				},
			},
			want: HeaderResponse{
				Quota:     Quota{Remaining: 15},
				Language:  "fr-FR",
				Length:    123,
				ItemSize:  10,
				Page:      1,
				PageSize:  10,
				PageTotal: 1,
				NextPage:  2,
				PrevPage:  -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractHeaders(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_doRequest(t *testing.T) {
	type fields struct {
		options Options
	}
	type args struct {
		reqMethod  string
		mockStatus int
		wantStatus int
		out        any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should do request successfully",
			fields: fields{},
			args: args{
				reqMethod:  "GET",
				mockStatus: http.StatusOK,
				wantStatus: http.StatusOK,
				out:        nil,
			},
			wantErr: false,
		},
		{
			name:   "should return error if status is not as expected",
			fields: fields{},
			args: args{
				reqMethod:  "GET",
				mockStatus: http.StatusInternalServerError,
				wantStatus: http.StatusOK,
			},
			wantErr: true,
		},
		{
			name:   "should return error if out is not a pointer",
			fields: fields{},
			args: args{
				reqMethod: "GET",
				out:       Actor{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.options.LogHandler = defaultTestLoggerHandler
			c := NewClientWithOpts(tt.fields.options)

			sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.args.mockStatus)
			}))
			defer sv.Close()

			req, err := http.NewRequest(tt.args.reqMethod, sv.URL, nil)
			if err != nil {
				t.Fatal(err)
			}

			_, err = c.doRequest(req, tt.args.wantStatus, tt.args.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_handlePagination(t *testing.T) {
	type args struct {
		page     int
		increase bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return -1 with page value being 0 (decrease)",
			args: args{
				page:     0,
				increase: false,
			},
			want: -1,
		},
		{
			name: "should return -1 with page value being 0 (increase)",
			args: args{
				page:     0,
				increase: true,
			},
			want: -1,
		},
		{
			name: "should return -1 with page value being -1 (decrease)",
			args: args{
				page:     -1,
				increase: false,
			},
			want: -1,
		},
		{
			name: "should return 1 with page value being 2 (decrease)",
			args: args{
				page:     2,
				increase: false,
			},
			want: 1,
		},
		{
			name: "should return -1 with page value being 1 (decrease)",
			args: args{
				page:     1,
				increase: false,
			},
			want: -1,
		},
		{
			name: "should return 2 with page value being 1 (increase)",
			args: args{
				page:     1,
				increase: true,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handlePagination(tt.args.page, tt.args.increase); got != tt.want {
				t.Errorf("handlePagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_doGetRequest(t *testing.T) {
	type fields struct {
		options Options
		cache   cache.Cache
	}
	type args struct {
		origin string
		query  []QueryOptions
		out    any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BaseResponse
		wantErr bool
	}{
		{
			name: "should call get request successfully",
			fields: fields{
				options: DefaultOptions,
				cache:   cache.NewMemoryCache(),
			},
			args: args{
				origin: "actors",
				out:    &Actor{},
			},
			want: BaseResponse{
				HeaderResponse: HeaderResponse{
					Page:      -1,
					PageSize:  -1,
					PageTotal: -1,
					ItemSize:  -1,
					NextPage:  -1,
					PrevPage:  -1,
					Language:  "",
					Quota: Quota{
						Remaining: -1,
					},
					Etag:   "",
					Length: 45,
				},
				Cached: true,
				Status: 200,
			},
			wantErr: false,
		},
		{
			name: "should return error if out is not a pointer",
			fields: fields{
				options: DefaultOptions,
				cache:   cache.NewMemoryCache(),
			},
			args: args{
				origin: "actors",
				out:    Actor{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte(`{"first_name": "Lorem", "last_name": "Ipsum"}`))
			}))
			defer server.Close()

			tt.fields.options.Endpoint = server.URL
			c := &Client{
				options: tt.fields.options,
				client:  server.Client(),
				cache:   tt.fields.cache,
				logger:  defaultTestLogger,
			}

			got, err := c.doGetRequest(tt.args.origin, tt.args.query, tt.args.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("doGetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doGetRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
