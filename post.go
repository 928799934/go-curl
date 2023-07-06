package curl

import "net/http"

func Http1Post(options []string, header http.Header, body []byte, url string) ([]byte, error) {
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
	// 补充 --http1
	if !bHttp11 {
		options = append(options, "--http1.1")
	}
	return Exec(options, header, body, url)
}

func Http2Post(options []string, header http.Header, body []byte, url string) ([]byte, error) {
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
	return Exec(options, header, body, url)
}

func Http1PostWithStream(options []string, header http.Header, body []byte, url string, fn func([]byte) error) error {
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
	// 补充 --http1
	if !bHttp11 {
		options = append(options, "--http1.1")
	}
	return Stream(options, header, body, url, fn)
}

func Http2PostWithStream(options []string, header http.Header, body []byte, url string, fn func([]byte) error) error {
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
	return Stream(options, header, body, url, fn)
}
