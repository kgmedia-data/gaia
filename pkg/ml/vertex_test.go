package ml

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInferSummaryRestVertex(t *testing.T) {
	input := []Summary{
		{GroupID: "main-1", Content: "IHSG dan Rupiah Menguat di Akhir Sesi"},
		{GroupID: "main-1", Content: "IHSG Ditutup Turun, Rupiah Menguat"},
		{GroupID: "main-1", Content: "Imbas Boikot, Penjualan Unilever di Indonesia Turun 15 Persen"},
		{GroupID: "main-1", Content: "Kejar Target Penerbitan, CGS-CIMB Sekuritas Terbitkan 6 Waran Terstruktur"},
		{GroupID: "main-2", Content: "Foreign Direct Investment (FDI): Pengertian, Jenis, dan Contohnya"},
		{GroupID: "main-2", Content: "Perluas Kemitraan, Unilever Teken Kerja Sama dengan GP Ansor"},
		{GroupID: "comp-3", Content: "5 Saham Paling Cuan Pekan Ini, Ada BMRI, hingga SRTG"},
		{GroupID: "comp-4", Content: "Laba Bersih Turun 10,4 Persen, Bos Unilever: Kami Terdampak Sentimen Konsumen Negatif..."},
		{GroupID: "comp-4", Content: "Lima Saham Paling \"Boncos\" Sepekan, dari CUAN hingga BTPS"},
		{GroupID: "comp-4", Content: "Simak Rekomendasi Saham Perbankan untuk Pemburu Dividen"},
		{GroupID: "comp-5", Content: "Dukung Pemberdayaan Ekonomi, Industri FMCG dan Kemenag Teken MoU"},
		{GroupID: "comp-5", Content: "Rasakan Dampak Boikot karena Dukung Israel, Unilever Sebut Penjualan di Indonesia Menurun"},
	}
	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "summarization",
	}

	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	model, err := vertex.NewSummaryVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)
	fmt.Printf("model vertex config: %+v\n", model.vertex.config)

	resp, err := model.BatchSummarize("English", 2, 6, input)
	fmt.Println("resp", resp)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(resp))
}

func TestInferRAGVertex(t *testing.T) {
	projectID := "kgdata-aiml"
	location := "global"
	dataStores := "insighthub-article-data-testing_1729824816557"
	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "insighthub",
	}
	model, err := NewRAGVertexRest(projectID, location, dataStores, projectLabel, vertex)
	assert.NoError(t, err)
	fmt.Printf("model vertex config: %+v\n", model.vertex.config)

	resp, err := model.Infer("siapa competitor yang sedang masif campaignnya?", "USER")
	assert.NoError(t, err)

	jsonData, err := json.MarshalIndent(resp, "", "  ")
	assert.NoError(t, err)
	err = os.WriteFile("test_rag.json", jsonData, 0644)
	assert.NoError(t, err)
}