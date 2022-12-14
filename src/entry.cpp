// <entry> -*- C++ -*-
// ~~nonstandard_tu~~ CPP_ONLY

#include "lib/quiver.hpp"

// begin forward declarations

int main(int argc, const char** argv);

// end forward declarations

int main(int argc, const char** argv) {
    Quiver<int, int> q;
    QuiverNodeRef new_node_ref = q.insert_node(99);
    QuiverNode<int, int>* new_node = new_node_ref.find_in_quiver(&q);
    printf(
        "node value %d at index %zu in quiver\n",
        new_node->get_value(),
        new_node_ref.index
    );
    return 0;
}
