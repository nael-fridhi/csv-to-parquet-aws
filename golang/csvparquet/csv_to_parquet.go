package csvparquet

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

// Titanic schema of the titanic csv file which will be used to convert file to parquet
type Titanic struct {
	Pclass   string `parquet:"name=titanic_pclass, type=UTF8"`
	Survived string `parquet:"name=shoe_name, type=UTF8"`
	Name     string `parquet:"name=titanic_name, type=UTF8"`
	Sex      string `parquet:"name=titanic_sex, type=UTF8"`
	Age      string `parquet:"name=titanic_age, type=UTF8"`
	Sibsp    string `parquet:"name=titanic_sibsp, type=UTF8"`
	Parch    string `parquet:"name=titanic_parch, type=UTF8"`
	Ticket   string `parquet:"name=titanic_ticket, type=UTF8"`
	Fare     string `parquet:"name=titanic_fare, type=UTF8"`
	Cabin    string `parquet:"name=titanic_cabin, type=UTF8"`
	Embarked string `parquet:"name=titanic_embarked, type=UTF8"`
	Boat     string `parquet:"name=titanic_boat, type=UTF8"`
	Body     string `parquet:"name=titanic_body, type=UTF8"`
	HomeDest string `parquet:"name=titanic_home_dest, type=UTF8"`
}

// ConvertCsvToParquet used to convert the csv file to parquet file and write it in fs
func ConvertCsvToParquet() {
	var err error

	fw, err := local.NewLocalFileWriter("/tmp/parquetFile.parquet")
	if err != nil {
		log.Println("Can't create the parquet file check that the folder tmp exist", err)
		return
	}

	pw, err := writer.NewParquetWriter(fw, new(Titanic), 14)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	csvFile, _ := os.Open("/tmp/csvFile.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		titanic := Titanic{
			Pclass:   line[0],
			Survived: line[1],
			Name:     line[2],
			Sex:      line[3],
			Age:      line[4],
			Sibsp:    line[5],
			Parch:    line[6],
			Ticket:   line[7],
			Fare:     line[8],
			Cabin:    line[9],
			Embarked: line[10],
			Boat:     line[11],
			Body:     line[12],
			HomeDest: line[13],
		}
		if err = pw.Write(titanic); err != nil {
			log.Println("Write error", err)
		}
	}

	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}

	log.Println("Write Finished")
	fw.Close()
}
