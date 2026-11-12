package sitein

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestHomeModelEmptyCarouselsMarshalsAsArray(t *testing.T) {
	b, err := json.Marshal(&HomeModel{Carousels: make([]*CarouselItem, 0)})
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}
	if !strings.Contains(string(b), `"carousels":[]`) {
		t.Fatalf("empty carousels should marshal as [], got %s", b)
	}
}

func TestCarouselItemExposesOnlyDisplayFields(t *testing.T) {
	b, err := json.Marshal(&CarouselItem{Id: 1, Title: "t", AspectRatio: "16/9"})
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}
	var got map[string]interface{}
	if err = json.Unmarshal(b, &got); err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}

	want := []string{"id", "title", "description", "ctaText", "ctaLink", "color", "image", "alt", "aspectRatio"}
	if len(got) != len(want) {
		t.Fatalf("expected %d display fields, got %d: %s", len(want), len(got), b)
	}
	for _, key := range want {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing display field %q in %s", key, b)
		}
	}
	for _, key := range []string{"status", "sort", "deletedAt", "createdAt", "updatedAt"} {
		if _, ok := got[key]; ok {
			t.Fatalf("management field %q must not be exposed: %s", key, b)
		}
	}

	typ := reflect.TypeOf(CarouselItem{})
	for _, name := range []string{"Status", "Sort", "DeletedAt"} {
		if _, ok := typ.FieldByName(name); ok {
			t.Fatalf("CarouselItem must not declare management field %s", name)
		}
	}
}
