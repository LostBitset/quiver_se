#include "quiver.h"

#include <algorithm>
#include <iterator>
#include <utility>

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

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNodeRef* QuiverNode<N, E, C>::follow_edge_fwd(E edge) {
    return this->edge_container.fwd_lookup(edge);
}

// *TODO* Quiver<N, E, C>::insert_node

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
void Quiver<N, E, C>::insert_edge(QuiverNodeRef src, QuiverNodeRef dst, E edge) {
    src.find_in_quiver(this)->edge_container.insert(edge, dst);
    dst.find_in_quiver(this)->parents.insert(src);
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
std::vector<std::pair<QuiverNodeRef, E>> Quiver<N, E, C>::follow_all_fwd(
    QuiverNodeRef node_ref
) {
    // *TODO* all
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
std::vector<std::pair<QuiverNodeRef, E>> Quiver<N, E, C>::follow_all_rev(
    QuiverNodeRef node_ref
) {
    std::vector<std::pair<QuiverNodeRef, E>> res;
    QuiverNode<N, E, C>* node = node_ref.find_in_quiver(this);
    for (QuiverNodeRef parent : node->parents) {
        QuiverNode parent_node = parent.find_in_quiver(this);
        std::vector<E> edges = parent_node.edge_container.rev_lookup(node_ref);
        auto xform = [=](E edge) {
            return std::make_pair(parent, edge);
        };
        std::transform(
            edges.begin(),
            edges.end(),
            edges.begin(),
            xform
        );
        res.reserve(edges.size());
        res.insert(res.end(), edges.begin(), edges.end());
    }
    return res;
}
