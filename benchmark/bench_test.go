package main
import (
"testing"
"fmt"
)
func BenchmarkHash10(b *testing.B) {
for n := 0; n < b.N; n++ {
key := "1234567890"
Hash("6368616e676520746869732070617373", key)
}
}
func BenchmarkHash20(b *testing.B) {
for n := 0; n < b.N; n++ {
key := "12345678901234567890"
Hash("6368616e676520746869732070617373", key)
}
}