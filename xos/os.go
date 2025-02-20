package xos

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func Open(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDONLY, 0)
}

func Create(name string) (*os.File, error) {
	return os.Create(name)
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
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

func SocketCallMethod(ep any, methodName string, services map[string]reflect.Value, requestCls, responseCls string, request, empty any) (results []reflect.Value, err error) {
	var method reflect.Value
	for _, service := range services {
		method = service.MethodByName(methodName)
		if method.IsValid() {
			break
		}
	}
	if !method.IsValid() {
		s := fmt.Sprintf("no found method %s", methodName)
		if responseCls != "*socket.NoReturn" {
			err = errors.New(s)
			return
		}
	}
	if method.Kind() != reflect.Func {
		s := fmt.Sprintf("%s is not a method", methodName)
		if responseCls != "*socket.NoReturn" {
			err = errors.New(s)
			return
		}
	}

	args := []reflect.Value{
		reflect.ValueOf(ep),
		//reflect.ValueOf(context),
		//reflect.ValueOf(request),
	}

	switch requestCls {
	case "*socket.Request":
		args = append(args, reflect.ValueOf(request))
	case "*socket.Empty":
		args = append(args, reflect.ValueOf(empty))
	}

	defer func() {
		recoverVal := recover()
		if recoverVal == nil {
			return
		}
		s := fmt.Sprintf("%v", recoverVal)
		err = errors.New(s)
		panic(err)
	}()

	results = method.Call(args) // 分发到各 rpc 业务处理函数
	return
}

// 修复漏洞
//func Int2Int32(p int) (int32, error) {
//	if p < -2147483648 || p > 2147483647 {
//		return 0, fmt.Errorf("value %d out of range for int32", p)
//	}
//	return int32(p), nil
//}

func Int2Int32(p int) int32 {
	return int32(p)
}

func TLSEMail(serverName string) *tls.Config {
	return &tls.Config{
		ServerName: serverName,
	}
}

func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}
