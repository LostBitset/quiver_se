case class TiedMap[H[_]](inner: Map[H[?], ?]):
    def +[A](kv: (H[A], A)): TiedMap[H] =
        TiedMap(inner + kv)
    
    def get[A](key: H[A]): Option[A] =
        inner.get(key).asInstanceOf[Option[A]]
    
    def apply[A](key: H[A]): A =
        get(key).get
