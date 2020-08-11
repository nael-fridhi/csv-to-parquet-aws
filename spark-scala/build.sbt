javacOptions ++= Seq("-source", "1.8", "-target", "1.8", "-Xlint")

val sparkVersion = "3.0.0"

lazy val root = (project in file(".")).
  settings(
    name := "csv-parquet-scala",
    organization := "com.xebia.interview",
    version := "1.0",
    scalaVersion := "2.11.4",
    retrieveManaged := true,
    libraryDependencies += "org.typelevel" %% "cats-core" % "2.0.0",
    libraryDependencies += "com.amazonaws" % "aws-lambda-java-core" % "1.0.0",
    libraryDependencies += "com.amazonaws" % "aws-lambda-java-events" % "1.0.0",
    libraryDependencies ++= Seq(
        "org.apache.spark" % "spark-core_2.12" % sparkVersion,
        "org.apache.spark" % "spark-sql_2.12" %  sparkVersion
    )
  )

assemblyMergeStrategy in assembly := {
    case PathList("META-INF", xs @ _*) => MergeStrategy.discard
    case x => MergeStrategy.first
}