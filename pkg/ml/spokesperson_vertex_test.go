package ml

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfer_SpokespersonRestVertex(t *testing.T) {

	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "spokesperson",
	}

	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	spokesperson, err := vertex.NewSpokespersonVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)
	// fmt.Printf("model vertex config: %+v\n\n", spokesperson.vertex.config)

	resp, err := spokesperson.Infer(`<p><strong><!--img.1--></strong></p>
	<p><strong>SEMARANG, KOMPAS.com -</strong> Gelombang protes kaum buruh di Jawa Tengah semakin membesar terhadap penolakan Program Tabungan Perumahan Rakyat (Tapera).</p>
	<p>Ratusan buruh di Kota Semarang berunjukrasa menolak program itu di depan Kantor Gubernur Jateng, Kamis (6/6/2024).</p>
	<p>Tak hanya aksi demo, Kepala Dinas Tenaga Kerja dan Transmigrasi (Disnakertrans) Jateng Ahmad Aziz mengatakan, sejak kebijakan itu disampaikan oleh pemerintah pusat, banyak kalangan buruh dan pengusaha mendatangi Kantor Dinas Tenaga Kerja di kabupaten/kota.</p>
	<p><strong>Baca juga: <a href="http://nasional.kompas.com/read/2024/06/07/13560841/moeldoko-sebut-pemberlakuan-tapera-menunggu-aturan-3-kementerian-maksimal" target="_self">Moeldoko Sebut Pemberlakuan Tapera Menunggu Aturan 3 Kementerian, Maksimal hingga 2027</a></strong></p>
	<p>Berangkat dengan keresahan mendalam, buruh dan pengusaha melakukan audiensi dan berdiskusi terkait program Tapera yang bakal memotong 2,5 persen upah mereka.</p>
	<p>"Dua hari terakhir ini di kabupaten/kota juga ada audensi dari temen-temen apindo maupun temen-temen dari serikat pekerja,serikat buruh. Dua-duanya datang ke Dinas Tenaga Kerja terkait dengan isu Tapera ini," ungkap Aziz kepada wartawan, Jumat (7/6/2024).</p>
	<p>Dia menilai, Apindo dan serikat buruh langsung merespon munculnya kebijakan baru itu. Hanya saja Apindo tidak menolak secara tegas. Sedangkan buruh menolak tegas.</p>
	<!--video.1-->
	<p><br /><br /></p>
	<p>"Alasan mereka (buruh) itu terkait dengan pembebanan. Dimaksimalkan aja kaitannya dengan BPJS itu. Di BPJS kan sudah ada JHT ada JKP dan lain sebagainya itu," beber Aziz.</p>
	<p>Kemudian terkait penambahan iuran dari pihak pekerja dan dari pihak pengusaha dikhawatirkan terjadi penyelewengan.</p>
	<p>Menurutnya para buruh berharap agar rumah yang dijanjikan diberikan terlebih dahulu, lalu dibayar dengan iuran atau cicilan bulanan setelahnya.</p>
	<p>"Terus juga kepastian itu terkait dengan menabung. Harapannya kalo mereka kan mendapatkan rumah dulu, terus nyicil. Salah satunya itu," terangnya.</p>
	<p><strong>Baca juga: <a href="http://regional.kompas.com/read/2024/06/06/204300278/ratusan-buruh-di-semarang-tolak-tapera--program-tidak-masuk-akal" target="_self">Ratusan Buruh di Semarang Tolak Tapera: Program Tidak Masuk Akal</a></strong></p>
	<p>Sementara itu, Tapera masih akan ditindaklanjuti dengan Permenaker. Dia menjelaskan Permenaker akan membahas lebih detail mengenai mekanisme pemungutan hingga kelompok masyarakat yang wajib mengikuti program itu.</p>
	<p>Lebih lanjut, Aziz memastikan aspirasi melalui audiensi dan unjuk rasa para buruh bakal disampaikan kepada Pj Gubernur Jateng Nana Sudjana, Sekda Jateng Sumarno, hingga Menteri Ketenagakerjaan RI Ida Fauziyah.</p>
	<p>"Iya kita akan sampaikan ya, saya akan laporkan ke pimpinan dalam hal ini Pak Gubernur dan Pak Sekda, juga akan sampaikan ke Kementerian Ketenagakerjaan," tandasnya.</p>
	<p>Diberitakan sebelumnya, ratusan buruh di Kota Semarang, Jawa Tengah beramai-ramai melakukan untuk rasa menolak Program Tabungan Perumahan Rakyat (Tapera) di depan Kantor Gubernur Jateng, Kamis (6/6/2024).</p>
	<p>Koordinator Lapangan Aulia Hakim mengatakan, wacana pemotongan gaji buruh sekitar 2,5 persen untuk Tapera sangatlah tidak masuk akal. Aulia menilai, saat pensiun nilai uang yang diterima sudah berubah ketimbang sekarang.</p>
	<p>"Setelah kita iuran (sampai pensiun) jatuhnya adalah Rp 48 juta. Logikannya ketika kita nabung, ketika kita masa-masa pensiun (di usia) 58 tahun, hanya mendapat Rp 48 juta, tidak masuk logika," tutur Aulia ditemui di sela aksi.</p>`)
	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("resp: %s\n", string(jsonBytes))
	assert.NoError(t, err)
	// assert.Equal(t, 5, len(resp))
}

func TestInferBatch_SpokespersonRestVertex(t *testing.T) {

	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "spokesperson",
	}

	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	spokesperson, err := vertex.NewSpokespersonVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)
	// fmt.Printf("model vertex config: %+v\n\n", spokesperson.vertex.config)

	texts := make(map[string]string, 0)
	texts["0"] = `Pengurus Dewan Pimpinan Pusat Front Pemuda Muslim Maluku (DPP FPMM) menyelenggarakan syukuran tahun baru 2025, di Restoran Bandar Jakarta, Summarecon, Kota Bekasi, Sabtu (4/1). 

	Acara yang digelar Dewan Pimpinan Pusat (DPP) Front Pemuda Muslim Maluku (FPMM) dikemas dalam bentuk syukuran dan hiburan dengan mengambil tema “Selamat tahun baru 2025, jadikan tahun lalu sebagai pelajaran dan tahun ini sebagai jalan mewujudkan impian”. 

	Tampak hadir dalam acara syukuran, Ketua Umum DPP FPMM, H. Umar Key dan jajaran pengurus DPP, jajaran pembina, jajaran penasehat, para pengurus DPD FPMM, Keluarga Besar FPMM, para ulama, laskar FPMM, Srikandi FPMM Al-Muluk dan para tokoh. Acara tersebut juga dihadiri Calon Walikota Kota Bekasi terpilih, Tri Adhianto Tjahyono. 

	Acara dimulai dengan rangkaian pembukaan, pembacaan Ayat Suci Alquran, sambutan-sambutan, potong tumpeng, doa hiburan dan ramah tamah. 

	Ketua Umum FPMM, H. Umar Key mengungkapkan, syukuran ini bukan sekadar ajang silaturahmi biasa, melainkan juga upaya konkret untuk mempersatukan seluruh organisasi Maluku. 

	“Syukuran tahun baru 2025 ini mencerminkan semangat solidaritas dan kesatuan. Kami mewakili keluarga besar bina lindung juga mohon maaf lahir dan batin. Kedepan semoga semuanya akan menjadi lebih baik lagi. Setelah kita melakukan pilkada kita kembali hidup rukun damai jangan sampai ada pertikaian. Kepada bapak Tri Adhianto yang kami pilih semoga bisa menepati janji-janjinya dan tetap amanah,” ungkap H. Umar. 

	Syukuran tahun baru 2025 juga disertakan ucapan selamat dan potong tumpeng serta doa untuk ulang tahun istri pertama H. Umar Key, Mami Ayu (Sri Rahayu) yang ke 44 tahun sekaligus ucapan ulang tahun untuk tamu kehormatan Tri Adhianto, calon Wali Kota terpilih Kota Bekasi. 

	Sementara itu, Dewan Penasehat DPP FPMM, Irjen. Pol (Purn.) Drs. Murad Ismail dan pembinan DPP FPMM H Muhammad Ongen Sangaji menyampaikan bahwa tahun 2025 adalah refleksi dan harapan dan momentum perkuat kerukunan dan pemberdayaan umat.`
	texts["1"] = `Di tempat sama, pegurus DPD FPMM Kota Bekasi Ahmad Gusti Ohoitenan menegaskan bahwa berbagai permasalahan dan pencapaian yang silih berganti kian mendewasakan dan menguatkan. 

	“Terima kasih kepada semuanya dari jajaran pembina, penasehat dan semua pengurus serta tokoh yang hadir. Atas dedikasi, pengabdian, dan kerja kerasnya selama ini sehingga banyak capaian diraih dan dirasakan manfaatnya,” ungkapnya. 

	“Kami masih ada kekurangan dan keterbatasan. Hal ini menjadi refleksi kami kita untuk memperbaiki ke depannya. Kami ucapkan selamat dan terimaksih atas terselengarakanya acara ini. Semoga FPMM semakin maju dan berkembang. Kami ucapkan juga kepada Ibu Sri Rahayu yang berulang tahun semoga tetap sehat dan panjang umur,” ucapnya. 

	Tri Adhianto menyampaikaikan ucapan terima kasih atas undangan dan dukungan saat pilkada. Mas Tri, sapaan akrabnya, juga berbicara tentang Toleransi. 

	“Kota Bekasi menduduki peringkat kedua sebagai Kota Toleran versi Setara Institute. Ini harus dipertahankan, bukan nanti kita bisa peringkat pertama. Kita memiliki toleransi yang kuat,” katanya. 

	“Tentunya apa yang kita raih ini adalah hasil dari kerja semua pihak, termasuk masyarakat. Karena itu, saya sampaikan terima kasih atas komitmen menjaga nilai toleransi dan keberagaman. Kedepan semoga Kota Bekasi semakin maju dan bahagia warganya. Kepada Front Pemuda Muslim Maluku, semoga semakin maju dan tambah bermanfaat, tetap menjaga Persatuan dan Kesatuan Bangsa, sehingga Kehidupan Toleransi dan hidup dalam Kedamaian akan tetap dijaga dalam interaksi sosial Kemasyarakatan,” pungkas Tri Adhianto.`

	resp, err := spokesperson.InferBatch(texts)
	assert.NoError(t, err)
	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("resp: %s\n", string(jsonBytes))
	assert.NoError(t, err)
	// assert.Equal(t, 5, len(resp))
}
