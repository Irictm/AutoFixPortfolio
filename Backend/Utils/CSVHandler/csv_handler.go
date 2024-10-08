package csvHandler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type fileUpload struct {
	CSVFile *multipart.FileHeader `form:"file"`
}

type CsvHandler struct{}

func (*CsvHandler) AttachCSV(c *gin.Context) error {
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
		err = fmt.Errorf("failed writing buffer to context: - %w", err)
		return err
	}
	return nil
}

func (*CsvHandler) ReceiveCSV(c *gin.Context) ([][]string, error) {
	var csvfile fileUpload
	var err error
	if err = c.ShouldBind(&csvfile); err != nil {
		err = fmt.Errorf("failed binding to json: - %w", err)
		return nil, err
	}

	if csvfile.CSVFile == nil {
		err = fmt.Errorf("failed, CSV file missing")
		return nil, err
	}

	file, err := csvfile.CSVFile.Open()
	if err != nil {
		err = fmt.Errorf("failed opening CSV file: - %w", err)
		return nil, err
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		err = fmt.Errorf("failed reading CSV file: - %w", err)
		return nil, err
	}

	return records, nil
}
