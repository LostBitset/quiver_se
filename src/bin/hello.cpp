// <hello> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~noheader~~ TU
// REASON target("bincode")

#include <cstdio>

// ~~forward_def_main~~ NEXT_ITEM
// REASON target("bincode")
int main(int argc, const char** argv);

void hello_world() {
    printf("Hello, world!\n");
}

int main(int argc, const char** argv) {
    hello_world();
    return 0;
}
