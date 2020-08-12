package com.aidodev.csvparquet

import org.apache.spark.sql.{SaveMode, SparkSession}
import scala.collection.JavaConverters._
import java.net.URLDecoder
import com.amazonaws.services.lambda.runtime.events.S3Event

class Main {

  def decodeS3Key(key: String): String = URLDecoder.decode(key.replace("+", " "), "utf-8")

  def getSourceBuckets(event: S3Event): java.util.List[String] = {
    val objectKey = event.getRecords.asScala.map(record => decodeS3Key(record.getS3.getObject.getKey)).asJava
    val bucketName = event.getRecords.asScala.map(record => decodeS3Key(record.getS3.getBucket.getName)).asJava
    csvToParquet(objectKey, bucketName)

    return "Done!"
  }

  def csvToParquet(objectKey: String, bucketName: String) {
    val spark: SparkSession = SparkSession.builder()
        .master("local[*]")
        .appName("Main")
        .getOrCreate()

    spark.sparkContext.setLogLevel("ERROR")
    val csvObjectPath: String = "s3a://"+ bucketName + "/csv/" + objectKey
    var parquetObjectPath = csvObjectPath.replace("csv", "parquet")
    val df = spark.read.options(Map("inferschema"->"true","delimiter"->";","header"->"true"))
        .csv(csvObjectPath)
    df.write.mode(SaveMode.Overwrite).parquet(parquetObjectPath)
    }
}