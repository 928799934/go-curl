package curl

import "net/http"

func Http1Get(options []string, header http.Header, url string) ([]byte, error) {
	options = append(options, "--request", "GET")
	bCompressed := false
	bHttp11 := false
	for _, v := range options {
		if v == "--compressed" {
			bCompressed = true
			break
		}
		if v == "--http1.1" {
			bHttp11 = true
		}
	}
	// 补充 --compressed
	if !bCompressed {
		options = append(options, "--compressed")
	}
	// 补充 --http11
	if !bHttp11 {
		options = append(options, "--http1.1")
	}
	return Exec(options, header, nil, url)
}

func Http2Get(options []string, header http.Header, url string) ([]byte, error) {
	options = append(options, "--request", "GET")

	bCompressed := false
	bHttp2 := false
	bInsecure := false
	for _, v := range options {
		if v == "--compressed" {
			bCompressed = true
		}
		if v == "--http2" {
			bHttp2 = true
		}
		if v == "--insecure" {
			bInsecure = true
		}
		if bCompressed && bHttp2 && bInsecure {
			break
		}
	}
	// 补充 --compressed
	if !bCompressed {
		options = append(options, "--compressed")
	}
	// 补充 --http2
	if !bHttp2 {
		options = append(options, "--http2")
	}
	// 补充 --insecure
	if !bInsecure {
		options = append(options, "--insecure")
	}
	return Exec(options, header, nil, url)
}

func Http1GetWithStream(options []string, header http.Header, url string, fn func([]byte) error) error {
	options = append(options, "--request", "GET")
	bCompressed := false
	bHttp11 := false
	for _, v := range options {
		if v == "--compressed" {
			bCompressed = true
			break
		}
		if v == "--http1.1" {
			bHttp11 = true
		}
	}
	// 补充 --compressed
	if !bCompressed {
		options = append(options, "--compressed")
	}
	// 补充 --http11
	if !bHttp11 {
		options = append(options, "--http1.1")
	}
	return Stream(options, header, nil, url, fn)
}

func Http2GetWithStream(options []string, header http.Header, url string, fn func([]byte) error) error {
	options = append(options, "--request", "GET")
	bCompressed := false
	bHttp2 := false
	bInsecure := false
	for _, v := range options {
		if v == "--compressed" {
			bCompressed = true
		}
		if v == "--http2" {
			bHttp2 = true
		}
		if v == "--insecure" {
			bInsecure = true
		}
		if bCompressed && bHttp2 && bInsecure {
			break
		}
	}
	// 补充 --compressed
	if !bCompressed {
		options = append(options, "--compressed")
	}
	// 补充 --http2
	if !bHttp2 {
		options = append(options, "--http2")
	}
	// 补充 --insecure
	if !bInsecure {
		options = append(options, "--insecure")
	}
	return Stream(options, header, nil, url, fn)
}
