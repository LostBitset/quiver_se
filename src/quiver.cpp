#include "quiver.h"

#include <algorithm>
#include <iterator>

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNode<N, E, C>* QuiverNodeRef::find_in_quiver(Quiver<N, E, C> *quiver) {
    quiver->get_node(this->index);
}

template <typename E, typename R>
SimpleQuiverEdge<E, R>::SimpleQuiverEdge() {}

template <typename E, typename R>
void SimpleQuiverEdge<E, R>::insert(E edge, R node_ref) {
    if (!this->backing_map.contains(edge)) {
        this->backing_map[edge] = node_ref;
    }
}

template <typename E, typename R>
R* SimpleQuiverEdge<E, R>::fwd_lookup(E edge) {
    return this->backing_map.at(&edge);
}

template <typename E, typename R>
std::vector<E> SimpleQuiverEdge<E, R>::rev_lookup(R node_ref) {
    std::vector<E> res;
    auto pred = [=](std::pair<E, R> kv) {
        return kv->second.index == node_ref.index;
    };
    auto xform = [=](std::pair<E, R> kv) {
        return kv->first;
    };
    std::copy_if(
        this->backing_map.begin(),
        this->backing_map.end(),
        std::back_inserter(res),
        pred
    );
    std::transform(
        res.begin(),
        res.end(),
        res.begin(),
        xform
    );
    return res;
}
