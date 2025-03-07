package ml

import (
	"fmt"
)

type SpokespersonVertexRest struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewSpokespersonVertexRest(projectID, location string, projectLabel ProjectLabel) (*SpokespersonVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		SetSystemInstruction(`You are a tools for extracting spokesperson from indonesian articles, it's quotes, and sentiment from news articles. each quotes has its own sentiment. some quotes may be indirect, but only extract if the subject is known. If there is nothing, return empty
		
		extract following information using Bahasa Indonesia: Name of the spokerperson (column name), job title of the spokerperson, including position, name of company/organization that is the spokesperson belong (column job_title), sentiment of the spokerperson (column sentiment), quotes or statements of the spokerperson with its each sentiment  (column statements).
		
		example-1
		input: 
		ID: 123, Content: Kali ini, ajang pemeran otomotif yang diadakan di JIExpo Kemayoran, Jakarta Pusat, tersebut bakal dimeriahkan oleh penyanyi ternama Tanah Air, Iwan Fals.
		Keseruan itu akan dibalut dalam agenda bertajuk IIMS Invinite Live 2025. Kehadiran Iwan Fals diyakini akan menambah keramaian di IIMS tahun ini. 
		ID: 234, Content: TEMPO.CO, Jakarta - Ahli gizi Tri Mutiara Ramdani, mengingatkan pentingnya upaya memenuhi kebutuhan nutrisi mikro dan omega-3 bagi ibu hamil dan menyusui. Menurutnya, zat gizi makro dan mikro sangat penting untuk dipenuhi agar pertumbuhan bayi optimal. Hal itu disampaikannya pada acara Kumpul ASIK (Asupan Ibu Berkualitas Karena) Blackmores pada 2 Februari 2025 di Kota Kasablanka, Jakarta, yang diselenggarakan mengedukasi kaum ibu tentang pentingnya nutrisi anak yang optimal untuk mendapatkan kehamilan sehat dan ASI berkualitas.
		Menurutnya, jika kebutuhan nutrisi makro cenderung gampang terpenuhi setiap hari, tidak demikian halnya dengan nutrisi mikro. Jumlah nutrisi mikro yang terkandung dalam makanan tergolong sedikit. Di samping itu, kualitas dan kuantitas nutrisi mikro rentan rusak karena makanan kerap diolah dengan beragam teknik memasak. Padahal, fungsi nutrisi tersebut sangat krusial untuk mendukung tumbuh kembang janin serta menghasilkan ASI berkualitas.
		Itulah sebabnya konsumsi suplemen secara teratur sangat penting untuk membantu memenuhi kebutuhan nutrisi mikro pada masa kehamilan dan menyusui. Blackmores Pregnancy & Breast-Feeding Gold (PBFG) hadir sebagai solusi efektif untuk mewujudkan pertumbuhan optimal janin serta produksi ASI berkualitas karena mengandung 17 nutrisi mikro, termasuk omega-3 yang diformulasikan khusus bagi ibu hamil dan menyusui.
		Pentingnya omega-3 bagi kecerdasanSelain nutrisi mikro, omega-3 juga penting untuk tumbuh kembang otak anak agar kecerdasannya maksimal. Kombinasi nutrisi mikro dan omega-3 akan membantu memenuhi kebutuhan nutrisi anak pada 1.000 hari pertama kehidupan, tepatnya sejak masa awal kehamilan, agar pertumbuhannya optimal.
		“Kami percaya setiap ibu berhak mendapatkan nutrisi terbaik untuk dirinya dan buah hatinya dan kami sangat berkomitmen untuk memberikan awal yang baik untuk buah hati dan mendukung perjalanan kehamilan dan menyusui yang sehat melalui produk-produk kami.” ujar Juliana Nur Wulan, Brand Representative Kalbe Blackmores Nutrition, lewat keterangannya.

		output: [{"id": "123", "name": "", "job_title": "", "sentiment": "neutral", "statements":[]}, {"id": "234", "name": "Tri Mutiara Ramdani", "job_title": "Ahli Gizi", "sentiment": "positive", "statements": [{"Quote": "Zat gizi makro dan mikro sangat penting untuk dipenuhi agar pertumbuhan bayi optimal.", "Sentiment": "positive"}, {"Quote": "Jika kebutuhan nutrisi makro cenderung gampang terpenuhi setiap hari, tidak demikian halnya dengan nutrisi mikro.", "Sentiment": "neutral"}, {"Quote": "Jumlah nutrisi mikro yang terkandung dalam makanan tergolong sedikit.", "Sentiment": "negative"}, {"Quote": "Kualitas dan kuantitas nutrisi mikro rentan rusak karena makanan kerap diolah dengan beragam teknik memasak.", "Sentiment": "negative"}, {"Quote": "Fungsi nutrisi tersebut sangat krusial untuk mendukung tumbuh kembang janin serta menghasilkan ASI berkualitas.", "Sentiment": "positive"}, {"Quote": "Konsumsi suplemen secara teratur sangat penting untuk membantu memenuhi kebutuhan nutrisi mikro pada masa kehamilan dan menyusui.", "Sentiment": "positive"}]},{"id": "234", "name": "Juliana Nur Wulan", "job_title": "Brand Representative Kalbe Blackmores Nutrition", "sentiment": "positive", "statements": [{"Quote": "“Kami percaya setiap ibu berhak mendapatkan nutrisi terbaik untuk dirinya dan buah hatinya dan kami sangat berkomitmen untuk memberikan awal yang baik untuk buah hati dan mendukung perjalanan kehamilan dan menyusui yang sehat melalui produk-produk kami.", "Sentiment": "positive"}]}]

		example-2
		input: JAKARTA, KOMPAS.TV - Presiden Prabowo Subianto menerima delapan konglomerat di Istana Kepresidenan Jakarta, Kamis (6/3/2025) sore.
		Bersama 8 pengusaha besar ini, Prabowo membahas pelaksanaan program makan bergizi gratis hingga BPI Danantara.
		Delapan konglomerat yang diterima Prabowo, yakni Anthony Salim dari Salim Group, Sugianto Kusuma alias Aguan dari Agung Sedayu Group, Prajogo Pangestu dari PT Barito Pacific Tbk, Garibaldi Thohir alias Boy Thohir dari Adaro, Franky Widjaja dari Sinar Mas Group, Dato Sri Tahir dari Mayapada Group, James Riady dari Lippo Group, serta Tomy Winata dari Artha Graha Group.
		Dalam pertemuan itu, Prabowo menyinggung situasi global dan kondisi dalam negeri, serta program prioritas pemerintah, seperti makan bergizi gratis, infrastruktur, industri tekstil, swasembada pangan dan energi, industrialisasi, hingga BPI Danantara.
		Menurut Sekretaris Kabinet, Teddy Indra Widjaja, Prabowo mengapresiasi dukungan para pengusaha terhadap berbagai kebijakan pemerintah.

		#prabowo #konglomeratindonesia #danantara #pengusaha

		output: [{"id": "0", "name": "Teddy Indra Widjaja", "job_title": "Sekretaris Kabinet", "sentiment": "neutral", "statements":["Quote":"Prabowo mengapresiasi dukungan para pengusaha terhadap berbagai kebijakan pemerintah", "Sentiment":"positive"]}]
		`).
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "string",
					},
					"name": map[string]interface{}{
						"type": "string",
					},
					"job_title": map[string]interface{}{
						"type": "string",
					},
					"sentiment": map[string]interface{}{
						"type": "string",
						"enum": []string{"positive", "negative", "neutral", "mixed"},
					},
					"statements": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"quote": map[string]interface{}{
									"type": "string",
								},
								"sentiment": map[string]interface{}{
									"type": "string",
									"enum": []string{"positive", "negative", "neutral"},
								},
							},
							"required": []string{"quote", "sentiment"},
						},
					},
				},
				"required": []string{"id", "name", "statements", "job_title", "sentiment"},
			},
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model))

	return &SpokespersonVertexRest{vertex: vertex}, nil
}

func (s SpokespersonVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SpokespersonVertexRest.(%v)(%v) %w", method, params, err)
}

func (s *SpokespersonVertexRest) Infer(content string) ([]Spokesperson, error) {

	s.vertex.SetContent(content, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}
	result := []Spokesperson{}
	err = ParseSingleResponseVertex(resp, &result)
	if err != nil {
		return nil, s.error(err, "Infer")
	}
	return result, nil
}

func (s *SpokespersonVertexRest) InferBatch(content map[string]string) (map[string][]Spokesperson, error) {
	s.vertex.ResetContentsParts()
	for i, c := range content {
		s.vertex.AddContent(fmt.Sprintf("ID %v: , Content: %v\n", i, c), "USER")
	}

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}
	var spokespersons []Spokesperson
	if err = ParseSingleResponseVertex(resp, &spokespersons); err != nil {
		return nil, s.error(err, "InferBatch - ParseSingleResponseVertex")
	}

	finalResult := make(map[string][]Spokesperson, len(content))
	for _, spokesperson := range spokespersons {
		finalResult[spokesperson.ID] = append(finalResult[spokesperson.ID], spokesperson)
	}

	return finalResult, nil
}
