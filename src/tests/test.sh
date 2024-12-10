#!/bin/bash

test_name=$(basename "$0" .sh)

result=out/tests/$test_name

mkdir -p "$result"

cat <<EOF | riscv64-linux-gnu-gcc -o "$result"/a.o -c -xc -
#include <stdio.h>

int main() {
    printf("Hello, World\n");
    return 0;
}
EOF

./rvld-learn "$result"/a.o