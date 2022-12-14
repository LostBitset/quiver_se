// <entry> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~no_header~~ TU
// REASON IS_ENTRY_POINT

#include "lib/quiver.h"

// ~~forward_decl_in_cpp~~ SEGMENT
// REASON PROPOGATE ~~no_header~~
// begin forward declarations

int main(int argc, const char** argv);

// end forward declarations

int main(int argc, const char** argv) {
    Quiver<int, int> q;
    q.insert_node(99);
    return 0;
}
