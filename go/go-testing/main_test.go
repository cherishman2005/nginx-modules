package main

import (
    "testing"
    "strconv"
)

func BenchmarkMap_Int_Int(b *testing.B) {
    b.ReportAllocs()
    m := make(map[int]int)
    for i := 0; i < b.N; i++ {
        //k := strconv.FormatInt(int64(i), 10)
        m[i] = i
    }
}

func BenchmarkMap_Int_Bool(b *testing.B) {
    b.ReportAllocs()
    m := make(map[int]bool)
    for i := 0; i < b.N; i++ {
        //k := strconv.FormatInt(int64(i), 10)
        m[i] = true
    }
}

func BenchmarkMap_String_Int(b *testing.B) {
    b.ReportAllocs()
    m := make(map[string]int)
    for i := 0; i < b.N; i++ {
        k := strconv.FormatInt(int64(i), 10)
        m[k] = i
    }
}

func BenchmarkMap_String_Bool(b *testing.B) {
    b.ReportAllocs()
    m := make(map[string]bool)
    for i := 0; i < b.N; i++ {
        k := strconv.FormatInt(int64(i), 10)
        m[k] = true
    }
}


```
go test -run=xxx -bench=. -benchtime="10s" -cpuprofile profile_cpu.out
```
