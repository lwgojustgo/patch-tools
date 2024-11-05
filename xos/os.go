package xos

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"reflect"
)

// 替换exec.CommandContext
func CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, arg...)
}

// 替换exec.CommandContext
func Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// 替换filepath.Join
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func MethodByName(value *reflect.Value, method string) reflect.Value {
	return value.MethodByName(method)
}

func TLSConfig() *tls.Config {
	return &tls.Config{InsecureSkipVerify: true}
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

func ListenAndServe(address string, handler http.Handler) error {
	return http.ListenAndServe(address, handler)
}
