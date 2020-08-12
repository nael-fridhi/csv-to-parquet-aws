package csvparquet

import org.apache.spark.sql.{SaveMode, SparkSession}
import scala.collection.JavaConverters._
import scala.collection.mutable._
import java.net.URLDecoder
import com.amazonaws.services.lambda.runtime.events.S3Event

class Main {

  def decodeS3Key(key: String): String = URLDecoder.decode(key.replace("+", " "), "utf-8")

  def getSourceBuckets(event: S3Event): Unit = {
    val objectKey = event.getRecords.asScala.map(record => decodeS3Key(record.getS3.getObject.getKey))
    val bucketName = event.getRecords.asScala.map(record => decodeS3Key(record.getS3.getBucket.getName))
    csvToParquet(objectKey(0), bucketName(0))
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