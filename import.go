package main

// Auto generation, edit and go fmt it

import (
 "github.com/yuanlixg/qlang/qlang/bufio"
 "github.com/yuanlixg/qlang/qlang/bytes"
 "github.com/yuanlixg/qlang/qlang/compress/flate"
 "github.com/yuanlixg/qlang/qlang/compress/gzip"
 "github.com/yuanlixg/qlang/qlang/container/list"
 "github.com/yuanlixg/qlang/qlang/context"
 "github.com/yuanlixg/qlang/qlang/crypto"
 "github.com/yuanlixg/qlang/qlang/crypto/aes"
 "github.com/yuanlixg/qlang/qlang/crypto/cipher"
 "github.com/yuanlixg/qlang/qlang/crypto/des"
 "github.com/yuanlixg/qlang/qlang/crypto/dsa"
 "github.com/yuanlixg/qlang/qlang/crypto/ecdsa"
 "github.com/yuanlixg/qlang/qlang/crypto/elliptic"
 "github.com/yuanlixg/qlang/qlang/crypto/hmac"
 "github.com/yuanlixg/qlang/qlang/crypto/md5"
 crand "github.com/yuanlixg/qlang/qlang/crypto/rand"
 "github.com/yuanlixg/qlang/qlang/crypto/rc4"
 "github.com/yuanlixg/qlang/qlang/crypto/rsa"
 "github.com/yuanlixg/qlang/qlang/crypto/sha1"
 "github.com/yuanlixg/qlang/qlang/crypto/sha256"
 "github.com/yuanlixg/qlang/qlang/crypto/sha512"
 "github.com/yuanlixg/qlang/qlang/crypto/subtle"
 "github.com/yuanlixg/qlang/qlang/crypto/tls"
 "github.com/yuanlixg/qlang/qlang/crypto/x509"
 "github.com/yuanlixg/qlang/qlang/crypto/x509/pkix"
 "github.com/yuanlixg/qlang/qlang/encoding/asn1"
 "github.com/yuanlixg/qlang/qlang/encoding/base64"
 "github.com/yuanlixg/qlang/qlang/encoding/binary"
 "github.com/yuanlixg/qlang/qlang/encoding/hex"
 "github.com/yuanlixg/qlang/qlang/encoding/json"
 "github.com/yuanlixg/qlang/qlang/encoding/pem"
 "github.com/yuanlixg/qlang/qlang/errors"
 "github.com/yuanlixg/qlang/qlang/fmt"
 "github.com/yuanlixg/qlang/qlang/github.com/yuanlixg/cookiejar"
 "github.com/yuanlixg/qlang/qlang/github.com/yuanlixg/goreq"
 "github.com/yuanlixg/qlang/qlang/hash/adler32"
 "github.com/yuanlixg/qlang/qlang/hash/crc32"
 "github.com/yuanlixg/qlang/qlang/hash/crc64"
 "github.com/yuanlixg/qlang/qlang/hash/fnv"
 "github.com/yuanlixg/qlang/qlang/io"
 "github.com/yuanlixg/qlang/qlang/io/ioutil"
 "github.com/yuanlixg/qlang/qlang/log"
 "github.com/yuanlixg/qlang/qlang/math"
 "github.com/yuanlixg/qlang/qlang/math/big"
 "github.com/yuanlixg/qlang/qlang/math/cmplx"
 "github.com/yuanlixg/qlang/qlang/math/rand"
 "github.com/yuanlixg/qlang/qlang/mime"
 "github.com/yuanlixg/qlang/qlang/mime/multipart"
 "github.com/yuanlixg/qlang/qlang/mime/quotedprintable"
 "github.com/yuanlixg/qlang/qlang/net"
 "github.com/yuanlixg/qlang/qlang/net/http"
 "github.com/yuanlixg/qlang/qlang/net/http/cgi"
 "github.com/yuanlixg/qlang/qlang/net/http/fcgi"
 "github.com/yuanlixg/qlang/qlang/net/http/httptest"
 "github.com/yuanlixg/qlang/qlang/net/http/httptrace"
 "github.com/yuanlixg/qlang/qlang/net/http/httputil"
 "github.com/yuanlixg/qlang/qlang/net/http/pprof"
 "github.com/yuanlixg/qlang/qlang/net/mail"
 "github.com/yuanlixg/qlang/qlang/net/rpc"
 "github.com/yuanlixg/qlang/qlang/net/rpc/jsonrpc"
 "github.com/yuanlixg/qlang/qlang/net/smtp"
 "github.com/yuanlixg/qlang/qlang/net/textproto"
 "github.com/yuanlixg/qlang/qlang/net/url"
 "github.com/yuanlixg/qlang/qlang/os"
 "github.com/yuanlixg/qlang/qlang/os/exec"
 "github.com/yuanlixg/qlang/qlang/os/signal"
 "github.com/yuanlixg/qlang/qlang/os/user"
 "github.com/yuanlixg/qlang/qlang/path"
 "github.com/yuanlixg/qlang/qlang/path/filepath"
 "github.com/yuanlixg/qlang/qlang/reflect"
 "github.com/yuanlixg/qlang/qlang/runtime"
 "github.com/yuanlixg/qlang/qlang/runtime/debug"
 rpprof "github.com/yuanlixg/qlang/qlang/runtime/pprof"
 "github.com/yuanlixg/qlang/qlang/runtime/trace"
 "github.com/yuanlixg/qlang/qlang/sort"
 "github.com/yuanlixg/qlang/qlang/strconv"
 "github.com/yuanlixg/qlang/qlang/strings"
 "github.com/yuanlixg/qlang/qlang/sync"
 "github.com/yuanlixg/qlang/qlang/sync/atomic"
 "github.com/yuanlixg/qlang/qlang/syscall"
 "github.com/yuanlixg/qlang/qlang/time"
 "github.com/yuanlixg/qlang/qlang/unicode"
 "github.com/yuanlixg/qlang/qlang/unicode/utf16"
 "github.com/yuanlixg/qlang/qlang/unicode/utf8"

	"qlang.io/qlang/meta"
	qlang "qlang.io/spec"

	// qlang builtin modules
	_ "qlang.io/qlang/builtin"
	_ "qlang.io/qlang/chan"
)

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("meta", meta.Exports) // import meta package

 qlang.Import("bufio", bufio.Exports)
 qlang.Import("bytes", bytes.Exports)
 qlang.Import("flate", flate.Exports)
 qlang.Import("gzip", gzip.Exports)
 qlang.Import("list", list.Exports)
 qlang.Import("context", context.Exports)
 qlang.Import("crypto", crypto.Exports)
 qlang.Import("aes", aes.Exports)
 qlang.Import("cipher", cipher.Exports)
 qlang.Import("des", des.Exports)
 qlang.Import("dsa", dsa.Exports)
 qlang.Import("ecdsa", ecdsa.Exports)
 qlang.Import("elliptic", elliptic.Exports)
 qlang.Import("hmac", hmac.Exports)
 qlang.Import("md5", md5.Exports)
 qlang.Import("crand", crand.Exports)
 qlang.Import("rc4", rc4.Exports)
 qlang.Import("rsa", rsa.Exports)
 qlang.Import("sha1", sha1.Exports)
 qlang.Import("sha256", sha256.Exports)
 qlang.Import("sha512", sha512.Exports)
 qlang.Import("subtle", subtle.Exports)
 qlang.Import("tls", tls.Exports)
 qlang.Import("x509", x509.Exports)
 qlang.Import("pkix", pkix.Exports)
 qlang.Import("asn1", asn1.Exports)
 qlang.Import("base64", base64.Exports)
 qlang.Import("binary", binary.Exports)
 qlang.Import("hex", hex.Exports)
 qlang.Import("json", json.Exports)
 qlang.Import("pem", pem.Exports)
 qlang.Import("errors", errors.Exports)
 qlang.Import("fmt", fmt.Exports)
 qlang.Import("cookiejar", cookiejar.Exports)
 qlang.Import("goreq", goreq.Exports)
 qlang.Import("adler32", adler32.Exports)
 qlang.Import("crc32", crc32.Exports)
 qlang.Import("crc64", crc64.Exports)
 qlang.Import("fnv", fnv.Exports)
 qlang.Import("io", io.Exports)
 qlang.Import("ioutil", ioutil.Exports)
 qlang.Import("log", log.Exports)
 qlang.Import("math", math.Exports)
 qlang.Import("big", big.Exports)
 qlang.Import("cmplx", cmplx.Exports)
 qlang.Import("rand", rand.Exports)
 qlang.Import("mime", mime.Exports)
 qlang.Import("multipart", multipart.Exports)
 qlang.Import("quotedprintable", quotedprintable.Exports)
 qlang.Import("net", net.Exports)
 qlang.Import("http", http.Exports)
 qlang.Import("cgi", cgi.Exports)
 qlang.Import("fcgi", fcgi.Exports)
 qlang.Import("httptest", httptest.Exports)
 qlang.Import("httptrace", httptrace.Exports)
 qlang.Import("httputil", httputil.Exports)
 qlang.Import("pprof", pprof.Exports)
 qlang.Import("mail", mail.Exports)
 qlang.Import("rpc", rpc.Exports)
 qlang.Import("jsonrpc", jsonrpc.Exports)
 qlang.Import("smtp", smtp.Exports)
 qlang.Import("textproto", textproto.Exports)
 qlang.Import("url", url.Exports)
 qlang.Import("os", os.Exports)
 qlang.Import("exec", exec.Exports)
 qlang.Import("signal", signal.Exports)
 qlang.Import("user", user.Exports)
 qlang.Import("path", path.Exports)
 qlang.Import("filepath", filepath.Exports)
 qlang.Import("reflect", reflect.Exports)
 qlang.Import("runtime", runtime.Exports)
 qlang.Import("debug", debug.Exports)
 qlang.Import("rpprof", rpprof.Exports)
 qlang.Import("trace", trace.Exports)
 qlang.Import("sort", sort.Exports)
 qlang.Import("strconv", strconv.Exports)
 qlang.Import("strings", strings.Exports)
 qlang.Import("sync", sync.Exports)
 qlang.Import("atomic", atomic.Exports)
 qlang.Import("syscall", syscall.Exports)
 qlang.Import("time", time.Exports)
 qlang.Import("unicode", unicode.Exports)
 qlang.Import("utf16", utf16.Exports)
 qlang.Import("utf8", utf8.Exports)

}

// ---------------------------------------------------------------------------
