package csvHandler

import (
	"bytes"
	"encoding/csv"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type fileUpload struct {
	CSVFile *multipart.FileHeader `form:"file"`
}

func AttachCSV(c *gin.Context) {
	sample := [][]string{
		{"Kilometraje", "Sedan", "Hatchback", "SUV", "Pickup", "Furgoneta"},
		{"0 - 5_000", "0", "0", "0", "0", "0"},
		{"5_001 - 12_000", "0.03", "0.03", "0.05", "0.05", "0.05"},
		{"12_001 - 25_000", "0.07", "0.07", "0.09", "0.09", "0.09"},
		{"25_000 - 40_000", "0.12", "0.12", "0.12", "0.12", "0.12"},
		{"40_000 - mas", "0.2", "0.2", "0.2", "0.2", "0.2"},
	}

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(sample)

	_, err := c.Writer.Write(csvBuffer.Bytes())
	if err != nil {
		log.Printf("Failed writing buffer to context - [%v]", err)
		return
	}
}

func ReceiveCSV(c *gin.Context) [][]string {
	var csvfile fileUpload
	if err := c.ShouldBind(&csvfile); err != nil {
		log.Printf("Failed binding to json - [%v]", err)
		return nil
	}

	if csvfile.CSVFile == nil {
		log.Printf("Failed, CSV file missing.")
		return nil
	}

	file, err := csvfile.CSVFile.Open()
	if err != nil {
		log.Printf("Failed opening CSV file - [%v]", err)
		return nil
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Printf("Failed reading CSV file - [%v]", err)
		return nil
	}

	return records
}
