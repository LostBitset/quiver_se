import grapple.json.{_, given}

import scala.language.implicitConversions

@main def main(filename: String): Unit =
    val input = io.Source.fromFile(filename).mkString
    val query = Json.parse(input).as[ConcolicQuery]
    val result = query.run
    val resultJson = Json.toJson(result)
    val outputW = JsonWriter(Console.out)
    outputW.write(resultJson.asInstanceOf[JsonObject])
    println()
