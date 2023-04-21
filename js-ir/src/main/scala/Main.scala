import grapple.json.{_, given}

@main def main: Unit =
    val input = Console.in.readLine
    val query = Json.parse(input).as[ConcolicQuery]
    val result = query.run
    ???
