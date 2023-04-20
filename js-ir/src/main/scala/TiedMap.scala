case class TiedMap[H[+_]](private val inner: Map[H[Any], Any] = Map[H[Any], Any]()):
    def +[A](kv: (H[A], A)): TiedMap[H] =
        TiedMap(inner + kv)
    
    def get[A](key: H[A]): Option[A] =
        inner.get(key).asInstanceOf[Option[A]]
    
    def apply[A](key: H[A]): A =
        get(key).get
    
    def mapP[B](f: [A] => (H[A], A) => B): Iterable[B] =
        inner
            .map({
                case (k, v) => f[v.type](k.asInstanceOf[H[v.type]], v)
            })
            .map(_.asInstanceOf[B])
