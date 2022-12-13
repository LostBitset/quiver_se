// <hello> -*- C++ -*-

// ~~noheader~~ reason: target("bincode")

#include <cstdio>

// ~~forward_def_main~~ reason: target("bincode")
int main(int argc, const char** argv);

void hello_world() {
    printf("Hello, world!\n");
}

int main(int argc, const char** argv) {
    hello_world();
    return 0;
}
