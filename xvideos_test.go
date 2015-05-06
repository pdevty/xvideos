package xvideos

import "testing"

func TestGetTagsJp(t *testing.T) {
	got, _ := getTags("http://jp.xvideos.com/video2601078/hot3x.net")
	if got == nil {
		t.Errorf("got %q expected not nil", got)
	}
}

func TestGetTagsUs(t *testing.T) {
	got, _ := getTags("http://www.xvideos.com/video2601078/hot3x.net")
	if got == nil {
		t.Errorf("got %q expected not nil", got)
	}
}

func TestGetEmptyUrl(t *testing.T) {
	got, _ := Get("")
	if got != nil {
		t.Errorf("got %q expected nil", got)
	}
}

func TestGetInvalidateUrl(t *testing.T) {
	got, _ := Get("http://hgehogejfak.com")
	if got != nil {
		t.Errorf("got %q expected nil", got)
	}
}

func TestGetXvideosUs(t *testing.T) {
	got, _ := Get("http://www.xvideos.com/best/")
	if got == nil {
		t.Errorf("got %q expected not nil", got)
	}
}

func TestGetXvideosJp(t *testing.T) {
	got, _ := Get("http://jp.xvideos.com/c/asian_woman-32/")
	if got == nil {
		t.Errorf("got %q expected not nil", got)
	}
}
