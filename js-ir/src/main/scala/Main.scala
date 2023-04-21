import grapple.json.{_, given}

import scala.io.Source
import scala.language.implicitConversions

@main def main: Unit =
    val input = Source.stdin.mkString("\n")
    val query = Json.parse(input).as[ConcolicQuery]
    val result = query.run
    val resultJson = Json.toJson(result)
    val outputW = JsonWriter(Console.out)
    outputW.write(resultJson.asInstanceOf[JsonObject])
