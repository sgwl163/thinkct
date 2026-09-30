package views

import "testing"

func TestWebModulePath(t *testing.T) {
	tests := []struct {
		name       string
		basePath   string
		sourceRoot string
		varName    string
		want       string
	}{
		{
			name:       "relative views path",
			basePath:   "../web/src/views/forum",
			sourceRoot: "views",
			varName:    "Topic",
			want:       "forum/topic",
		},
		{
			name:       "absolute Windows views path",
			basePath:   "P:\\go\\hotgo-2.0\\web\\src\\views\\forum",
			sourceRoot: "views",
			varName:    "Topic",
			want:       "forum/topic",
		},
		{
			name:       "absolute Unix API path",
			basePath:   "/opt/hotgo/web/src/api/im",
			sourceRoot: "api",
			varName:    "Conversation",
			want:       "im/conversation",
		},
		{
			name:       "source root only",
			basePath:   "api",
			sourceRoot: "api",
			varName:    "Topic",
			want:       "topic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := webModulePath(tt.basePath, tt.sourceRoot, tt.varName); got != tt.want {
				t.Fatalf("webModulePath(%q, %q, %q) = %q, want %q", tt.basePath, tt.sourceRoot, tt.varName, got, tt.want)
			}
		})
	}
}

func TestParseServFunName(t *testing.T) {
	curd := gCurd{}
	tests := []struct {
		templateGroup string
		varName       string
		want          string
	}{
		{templateGroup: "sys", varName: "Topic", want: "SysTopic"},
		{templateGroup: "forum", varName: "Category", want: "ForumCategory"},
		{templateGroup: "forum", varName: "ForumCategory", want: "ForumForumCategory"},
	}

	for _, tt := range tests {
		if got := curd.parseServFunName(tt.templateGroup, tt.varName); got != tt.want {
			t.Fatalf("parseServFunName(%q, %q) = %q, want %q", tt.templateGroup, tt.varName, got, tt.want)
		}
	}
}
