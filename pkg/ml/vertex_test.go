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
	// input := []Summary{
	// 	{GroupID: "main-1", Content: "IHSG dan Rupiah Menguat di Akhir Sesi"},
	// 	{GroupID: "main-1", Content: "IHSG Ditutup Turun, Rupiah Menguat"},
	// 	{GroupID: "main-1", Content: "Imbas Boikot, Penjualan Unilever di Indonesia Turun 15 Persen"},
	// 	{GroupID: "main-1", Content: "Kejar Target Penerbitan, CGS-CIMB Sekuritas Terbitkan 6 Waran Terstruktur"},
	// 	{GroupID: "main-2", Content: "Foreign Direct Investment (FDI): Pengertian, Jenis, dan Contohnya"},
	// 	{GroupID: "main-2", Content: "Perluas Kemitraan, Unilever Teken Kerja Sama dengan GP Ansor"},
	// 	{GroupID: "comp-3", Content: "5 Saham Paling Cuan Pekan Ini, Ada BMRI, hingga SRTG"},
	// 	{GroupID: "comp-4", Content: "Laba Bersih Turun 10,4 Persen, Bos Unilever: Kami Terdampak Sentimen Konsumen Negatif..."},
	// 	{GroupID: "comp-4", Content: "Lima Saham Paling \"Boncos\" Sepekan, dari CUAN hingga BTPS"},
	// 	{GroupID: "comp-4", Content: "Simak Rekomendasi Saham Perbankan untuk Pemburu Dividen"},
	// 	{GroupID: "comp-5", Content: "Dukung Pemberdayaan Ekonomi, Industri FMCG dan Kemenag Teken MoU"},
	// 	{GroupID: "comp-5", Content: "Rasakan Dampak Boikot karena Dukung Israel, Unilever Sebut Penjualan di Indonesia Menurun"},
	// }
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

	// resp, err := model.ProcessAndBatchSummarize("English", 2, 6, input)
	// fmt.Println("resp", resp)
	// assert.NoError(t, err)
	// assert.Equal(t, 5, len(resp))

	contents := `Create a summary of given indonesian news article title. Ignore the html tag and unrelated characters. For each group_id, ONLY generate 1 summary of 1 paragraph (e.g. main-1 only generated in 1 summary), which SHOULD contain only between %d and %d sentences. if there are 5 unique group_id, then also return summary of 5 group_id. Write the summarization in %s. Here are the input:

	group_id: main-1, content 1: 
	<img src="https://asset-2.tstatic.net/manado/foto/bank/images/Daftar-Bupati-di-Sumatera-Utara-dilantik-20-februari-2025.jpg"/>
<p><strong>TRIBUNMANADO.CO.ID </strong>- Daftar lengkap kepala daerah terpilih pada Pilkada 2024 di Sumatera Utara (Sumut) yang telah dipastikan dilantik besok, Kamis 20 Februari 2025.</p>
<p>Sebanyak 33 kabupaten/kota di Sumut menggelar Pilkada 2024.</p>
<p>Namun, paslon terpilih yang bakal dilantik besok hanya sebanyak 32 paslon dari semua kabupaten/kota, ditambah paslon terpilih gubernur-wakil gubernur.</p>
<p>Kepastian itu setelah pascaputusan Mahkamah Konstitusi (MK) dalam sidang dismissal pada 4-5 Februari lalu.</p>
<p>Sedangkan 1 paslon terpilih, yakni Bupati - Wakil Bupati Kabupaten Mandailing Natal belum akan dilantik.</p>
<p>Sebagaimana, gugatan Pilkada Kabupaten Mandailing Natal dengan perkara bernomor 32/PHPU.BUP-XXIII/2025, menjadi satu-satunya dari Sumatera Utara yang diputus MK berlanjut ke tahap pembuktian.</p>
<p>Lantas siapa saja 33 paslon kepala daerah terpilih Pilkada 2024 di Sumatera Utara yang dipastikan dilantik pada Kamis, 20 Februari 2025?</p>
<p>Berikut daftarnya:</p>
<p>1. Provinsi Sumatera Utara (Pilgub)</p>
<p>Muhammad Bobby Afif Nasution - Surya</p>
<p>2. Kota Medan</p>
<p>Rico Tri Putra Bayu - Zakiyuddin Harahap</p>
<p>3. Kota Gunungsitoli</p>
<p>Sowa’a laoli - Martinus Lase</p>
<p>4. Kota Binjai</p>
<p>Amir Hamzah - Hasanul Jihadi</p>
<p>5. Kota Padang Sidempuan</p>
<p>Letnan - Harry Phlevi</p>
<p>6. Kota Pematang Siantar</p>
<p>Wesly Silalahi - Herlina</p>
<p>7. Kota Sibolga</p>
<p>Akhmad Syukri Nazry Penarik - Pantas Maruba Lumbantobing</p>
<p>8. Kota Tanjung Balai</p>
<p>Mahyaruddin Salim - Muhammad Fadly Abdina</p>
<p>9. Kota Tebing Tinggi</p>
<p>Iman Irdian Saragih - Chairil Mukmin Tambunan</p>
<p>10. Kabupaten Asahan</p>
<p>Taufik Zainal Abidin - Rianto</p>
<p>11. Kabupaten Batubara</p>
<p>Baharuddin Siagian - Syafrizal</p>
<p>12. Kabupaten Dairi</p>
<p>Vickner Sinaga - Wahyu Daniel Sagala</p>
<p>13. Kabupaten Deli Serdang</p>
<p>Asri Ludin Tambunan - Lom Lom Suwondo</p>
<p>14. Kabupaten Humbang Hasundutan</p>
<p>Oloan P Nababan - Junita Rebeka Marbun</p>
<p>15. Kabupaten Karo</p>
<p>Antonius Ginting - Komando Tarigan</p>
<p>16. Kabupaten Labuhan Batu</p>
<p>Maya Hasmita - Jamri</p>
<p>17. Kabupaten Labuhan Batu Selatan</p>
<p>Fery Sahputa Simatupang - Syahdian Purba Siboro</p>
<p>18. Kabupaten Labuhan Batu Utara</p>
<p>Hendri Yanto Sitorus - Samsul Tanjung</p>
<p>19. Kabupaten Langkat</p>
<p>Syah Afandin -Tiorita BR Surbaksti</p>
<p>20. Kabupaten Nias</p>
<p>Ya’atulo Gulo - Aorta Lase</p>
<p>21. Kabupaten Nias Barat</p>
<p>Eliyunus Waruwu - Sozishkhi Hia</p>
<p>22. Kabupaten Nias Selatan</p>
<p>Sokhiatulo Laila - Yusuf Nache</p>
<p>23. Kabupaten Nias Utara</p>
<p>Amizoro  Waruwu - Yusman Zega</p>
<p>24. Kabupaten Padang Lawas</p>
<p>Puta Mahkota Alam - Achmad Fauzan Nasution</p>
<p>25. Kabupaten Padang Lawas Utara</p>
<p>Reski Basyah Harahap - Basri Harahap</p>
<p>26. Kabupaten Pakpak Bharat</p>
<p>Franc Bernhard - Mutsyuhito Solin</p>
<p>27. Kabupaten Samosir</p>
<p>Vandiko Timotius Gultom - Ariston Tua Sidauruk</p>
<p>28. Kabupaten Serdang Bedagai</p>
<p>Darma Wijaya - Adlin Umar Yusri Tambunan</p>
<p>29. Kabupaten Simalungun</p>
<p>Anton Achmad Saragih - Benny Gusman Sinaga</p>
<p>30. Kabupaten Tapanuli Selatan</p>
<p>Gus Irawan Pasaribu - Jafar Syahbuddin Sitonga</p>
<p>31. Kabupaten Tapanuli Tengah</p>
<p>Masinton Pasaribu - Mahmud Sitompul</p>
<p>32. Kabupaten Tapanuli Utara</p>
<p>Jonius Tripar Parsaoran Hutabarat - Deni Parlindungan</p>
<p>33. Kabupaten Toba</p>
<p>Effendi Napitupulu - Audi Murphy Sitorus</p>
<p>-</p>
<p><em><strong>(TribunManado.co.id)</strong></em></p>`
	resp, output, err := model.BatchSummarize(contents[:2000])
	fmt.Println("resp", resp)
	fmt.Println("output", output)
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
	resp, output, err := model.Infer(text)
	fmt.Println("resp", resp)
	fmt.Println("output", output)
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
