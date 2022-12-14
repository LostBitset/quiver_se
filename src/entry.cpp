// <entry> -*- C++ -*-
// ~~nonstandard_tu~~ CPP_ONLY

#include "lib/quiver.hpp"

// begin forward declarations

int main(int argc, const char** argv);

// end forward declarations

int main(int argc, const char** argv) {
    Quiver<int, int> q;
    QuiverNodeRef n1 = q.insert_node(99);
    QuiverNodeRef n2 = q.insert_node(43);
    q.insert_edge(n1, n2, 8);
    q.insert_edge(n1, n2, 7);
    q.insert_edge(n2, n1, 3);
    QuiverNode<int, int>* new_node = n1.find_in_quiver(&q);
    printf(
        "node value %d at index %zu in quiver\n",
        new_node->get_value(),
        n1.index
    );
    printf(
        "n1 backing_map size is %zu\n",
        new_node->get_edge_container().backing_map.size()
    );
    auto fwd_lookup = new_node->get_edge_container().fwd_lookup(8);
    printf("{n1}->fwd_lookup(8) returns %zu\n", fwd_lookup->index);
    auto fwd_vector = q.follow_all_fwd(n1);
    printf("[list]multiline {size=%zu}\n", fwd_vector.size());
    for (auto item : fwd_vector) {
        printf("[pair](%zu, %d)\n", item.first.index, item.second);
    }
    return 0;
}
