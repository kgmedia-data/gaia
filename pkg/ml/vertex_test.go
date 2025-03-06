package ml

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

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

	resp, err := model.ProcessAndBatchSummarize("English", 2, 6, input)
	fmt.Println("resp", resp)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(resp))

	contents := `Create a summary of given indonesian news article title. Ignore the html tag and unrelated characters. For each group_id, ONLY generate 1 summary of 1 paragraph (e.g. main-1 only generated in 1 summary), which SHOULD contain only between %d and %d sentences. if there are 5 unique group_id, then also return summary of 5 group_id. Write the summarization in %s. Here are the input:

	group_id: main-1, content 1: IHSG dan Rupiah Menguat di Akhir Sesi group_id: main-1, content 2: IHSG Ditutup Turun, Rupiah Menguat 
	group_id: main-1, content 3: Imbas Boikot, Penjualan Unilever di Indonesia Turun 15 Persen 
	group_id: main-1, content 4: Kejar Target Penerbitan, CGS-CIMB Sekuritas Terbitkan 6 Waran Terstruktur 
	group_id: main-2, content 5: Foreign Direct Investment (FDI): Pengertian, Jenis, dan Contohnya 
	group_id: main-2, content 6: Perluas Kemitraan, Unilever Teken Kerja Sama dengan GP Ansor 
	group_id: comp-3, content 7: 5 Saham Paling Cuan Pekan Ini, Ada BMRI, hingga SRTG 
	group_id: comp-4, content 8: Laba Bersih Turun 10,4 Persen, Bos Unilever: Kami Terdampak Sentimen Konsumen Negatif... 
	group_id: comp-4, content 9: Lima Saham Paling "Boncos" Sepekan, dari CUAN hingga BTPS 
	group_id: comp-4, content 10: Simak Rekomendasi Saham Perbankan untuk Pemburu Dividen 
	group_id: comp-5, content 11: Dukung Pemberdayaan Ekonomi, Industri FMCG dan Kemenag Teken MoU 
	group_id: comp-5, content 12: Rasakan Dampak Boikot karena Dukung Israel, Unilever Sebut Penjualan di Indonesia Menurun`
	resp, err = model.BatchSummarize(contents)
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
	err = os.WriteFile("vertex_test_rag.json", jsonData, 0644)
	assert.NoError(t, err)
}

func TestInferEntitySentimentVertex(t *testing.T) {
	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "sentiment",
	}

	model, err := vertex.NewEntitySentimentVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)

	text := `extract entity-based sentiment from this texts
			: text-1: Kasus Pemakaian Narkoba naik 10%% di Indonesia. 

			entity-keywords-1: Paminal, entityID-1: 1, entity-name: kepolisian
			entity-keywords-2: Polisi, entityID-2: 2, entity-name: kepolisian
			entity-keywords-3: Pagar Laut, entityID-3: 3, entity-name: kepolisian
			entity-keywords-4: Kepolisian Republik Indonesia (Polri), entityID-4: 4, entity-name: kepolisian
			entity-keywords-5: #kasuspolisi, entityID-5: 5, entity-name: kepolisian
			entity-keywords-6: #polisiviral, entityID-6: 6, entity-name: kepolisian
			entity-keywords-7: Parcok, entityID-7: 7, entity-name: kepolisian
			entity-keywords-8: Partai Cokelat, entityID-8: 8, entity-name: kepolisian
			entity-keywords-9: polisi, entityID-9: 9, entity-name: kepolisian
			entity-keywords-10: polisi tembak sopir mobil, entityID-10: 10, entity-name: kepolisian
			entity-keywords-11: Maut Di Ujung Pistol Sang Oknum Brigpol, entityID-11: 11, entity-name: kepolisian
			entity-keywords-12: Pembunuhan Sadis Oknum Polisi, entityID-12: 12, entity-name: kepolisian
			entity-keywords-13: #parcok, entityID-13: 13, entity-name: kepolisian
			entity-keywords-14: #PolisiJahat, entityID-14: 14, entity-name: kepolisian
			entity-keywords-15: #SeragamCokelat, entityID-15: 15, entity-name: kepolisian
			entity-keywords-16: Seragam Cokelat, entityID-16: 16, entity-name: kepolisian
			entity-keywords-17: #ReformasiPOLRI, entityID-17: 17, entity-name: kepolisian
	`
	resp, err := model.Infer(text)
	fmt.Println("resp", resp)
	assert.NoError(t, err)
}

func TestRenewToken(t *testing.T) {
	vertex, err := NewVertexRest()
	vertex.tokenExpiration = time.Now().Add(1 * time.Second)
	assert.NoError(t, err)

	// sleep until expired
	time.Sleep(2 * time.Second)
	assert.True(t, time.Now().After(vertex.tokenExpiration))

	tokenBefore := vertex.request.Header.Get("Authorization")

	err = vertex.RenewToken()
	assert.NoError(t, err)

	tokenAfter := vertex.request.Header.Get("Authorization")
	assert.NotEqual(t, tokenBefore, tokenAfter)
	assert.True(t, time.Now().Before(vertex.tokenExpiration))
}
