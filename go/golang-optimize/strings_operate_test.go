package main

import(
    "bytes"
    "fmt"
    "strings"
    //"time"
    "testing"
)

const NUM int = 100000

func AppendWithAdd(n int) {
    var s string
    for i:=0; i < n; i++{
        s = s + "string"
    }
}

func AppendWithSprintf(n int) {
    var s string
    for i:=0; i < n; i++{
        s = fmt.Sprintf("%s%s", s, "string")
    }
}

func AppendWithBytesBuffer(n int) {
    var byt bytes.Buffer
    for i:=0; i < n; i++{
        byt.WriteString("string")
    }
    byt.String()
}

func AppendWithStringBuilder(n int) {
    var sbuilder strings.Builder
    for i:=0; i < n; i++{
        sbuilder.WriteString("string")
    }
    sbuilder.String()
}


func Benchmark_AppendWithAdd(b *testing.B) {
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        AppendWithAdd(NUM)
    }
}

func Benchmark_AppendWithSprintf(b *testing.B) {
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        AppendWithSprintf(NUM)
    }
}

func Benchmark_AppendWithBytesBuffer(b *testing.B) {
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        AppendWithBytesBuffer(NUM)
    }
}

func Benchmark_AppendWithStringBuilder(b *testing.B) {
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        AppendWithStringBuilder(NUM)
    }
}