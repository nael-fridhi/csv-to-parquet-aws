package com.aidodev.csvparquet

import org.apache.spark.sql.{SaveMode, SparkSession}

object CsvToParquet extends App {
  val spark: SparkSession = SparkSession.builder()
    .master("local[*]")
    .appName("CsvToParquet")
    .getOrCreate()

  spark.sparkContext.setLogLevel("ERROR")

  val df = spark.read.options(Map("inferschema"->"true","delimiter"->";","header"->"true"))
    .csv("./data/titanic.csv")
  df.write.mode(SaveMode.Overwrite).parquet("/tmp/parquet/titanic_csv_parquet.parquet")

}