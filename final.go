package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


type Bidang struct {
	ID          int
	Nama        string
	Deskripsi   string
	IconASCII   string
}

type Mentor struct {
	ID           int
	Nama         string
	Bidang       string
	Rating       float64
	Pengalaman   int
	JumlahSiswa  int
	Biaya        int
	Keahlian     string
	Lulusan      string
	Aktif        bool
}

type Jadwal struct {
	Tanggal string
	Hari    string
	Jam     string
	Durasi  int
}

type Pelatihan struct {
	ID          string
	Nama        string
	Bidang      string
	Level       string
	Kategori    string
	MentorID    int
	MentorNama  string
	Rating      float64
	Harga       int
	Durasi      int
	Kuota       int
	Terdaftar   int
	Jadwal      Jadwal
	Deskripsi   string
	Skill       string
	Target      string
}

type Peserta struct {
	ID       int
	Nama     string
	Email    string
	NoHP     string
	Kota     string
	Level    string
	Bidang   string
}

type Sertifikat struct {
	ID             string
	PesertaID      int
	PesertaNama    string
	PelatihanID    string
	PelatihanNama  string
	MentorNama     string
	TanggalSelesai string
	Grade          string
	Valid          bool
}

type RiwayatPendaftaran struct {
	ID             int
	PesertaID      int
	PesertaNama    string
	PelatihanID    string
	PelatihanNama  string
	MentorNama     string
	Kategori       string
	TanggalDaftar  string
	TanggalMulai   string
	TanggalSelesai string
	TotalBiaya     int
	Status         string
}



var daftarBidang = []Bidang{
	{1, "Cyber Security", "Keamanan siber & proteksi data", ""},
	{2, "Web Development", "Pengembangan web modern", ""},
	{3, "UI/UX Design", "Desain antarmuka pengguna", ""},
	{4, "Artificial Intelligence", "Kecerdasan buatan & ML", ""},
	{5, "Data Science", "Analisis data & visualisasi", ""},
	{6, "Cloud Computing", "Komputasi awan & infrastruktur", ""},
	{7, "Networking", "Jaringan komputer & protokol", ""},
	{8, "Blockchain", "Teknologi blockchain & crypto", ""},
	{9, "Mobile Development", "Pengembangan aplikasi mobile", ""},
	{10, "DevOps", "Development & Operations terpadu", ""},
}

var daftarMentor []Mentor
var daftarPelatihan []Pelatihan
var daftarPeserta []Peserta
var daftarSertifikat []Sertifikat
var daftarRiwayat []RiwayatPendaftaran

var pesertaIDCounter int = 1
var riwayatIDCounter int = 1
var reader = bufio.NewReader(os.Stdin)

// Warna ANSI
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorBold   = "\033[1m"
)


func inisialisasiMentor() {
	daftarMentor = []Mentor{
		{1, "Dr. Budi Santoso", "Cyber Security", 4.9, 12, 3200, 500000, "Ethical Hacking, CISSP", "ITB", true},
		{2, "Rina Wijayanti, M.Kom", "Web Development", 4.8, 8, 2800, 400000, "React, Node.js, Go", "UI", true},
		{3, "Kevin Hartono", "UI/UX Design", 4.7, 6, 1900, 350000, "Figma, Adobe XD", "Binus", true},
		{4, "Dr. Siti Rahayu", "Artificial Intelligence", 4.9, 15, 4100, 600000, "TensorFlow, PyTorch", "ITS", true},
		{5, "Ahmad Fauzi, M.T", "Data Science", 4.6, 9, 2400, 420000, "Python, R, Tableau", "UGM", true},
		{6, "Lina Permata", "Cloud Computing", 4.8, 7, 2100, 450000, "AWS, GCP, Azure", "UNPAD", true},
		{7, "Denny Kurniawan", "Networking", 4.5, 10, 1800, 380000, "Cisco, CCNA", "UNDIP", true},
		{8, "Mega Puspita, M.Sc", "Blockchain", 4.7, 5, 1300, 500000, "Solidity, Ethereum", "ITS", true},
		{9, "Rizal Firdaus", "Mobile Development", 4.6, 7, 2200, 400000, "Flutter, React Native", "Telkom", true},
		{10, "Hendra Saputra", "DevOps", 4.8, 11, 2700, 470000, "Docker, K8s, CI/CD", "UI", true},
		{11, "Sarah Johnson", "Cyber Security", 4.7, 9, 2500, 480000, "Penetration Testing", "Stanford", true},
		{12, "Michael Chen", "Web Development", 4.9, 13, 3500, 550000, "Vue.js, TypeScript", "MIT", true},
		{13, "Dewi Anggraini", "UI/UX Design", 4.8, 8, 2300, 370000, "User Research", "FSRD ITB", true},
		{14, "Prof. James Williams", "Artificial Intelligence", 5.0, 20, 5200, 750000, "Deep Learning", "Harvard", true},
		{15, "Andika Pratama, M.Sc", "Data Science", 4.5, 6, 1700, 390000, "Machine Learning", "IPB", true},
		{16, "Nadia Kusuma", "Cloud Computing", 4.6, 5, 1500, 420000, "Serverless", "Binus", true},
		{17, "Yudi Wibowo", "Networking", 4.7, 12, 2100, 400000, "Network Security", "PENS", true},
		{18, "Clara Situmorang", "Blockchain", 4.6, 4, 980, 460000, "Web3, NFT", "UPH", true},
		{19, "Fajar Nugroho", "Mobile Development", 4.9, 10, 3100, 500000, "Kotlin, iOS", "ITS", true},
		{20, "Bella Anindita", "DevOps", 4.7, 8, 2000, 450000, "Jenkins, GitOps", "Telkom", true},
		{21, "Reza Hakim, M.Kom", "Web Development", 4.5, 5, 1600, 360000, "Laravel, MySQL", "UB", true},
		{22, "Putri Maharani", "Data Science", 4.8, 7, 2200, 410000, "NLP", "UI", true},
		{23, "Eko Budiyanto", "Cyber Security", 4.6, 8, 2000, 460000, "OSINT, Forensic", "STTI", true},
		{24, "Vanessa Olivia", "UI/UX Design", 4.9, 9, 2800, 400000, "Design System", "NUS", true},
		{25, "Dr. Irfan Maulana", "Cloud Computing", 4.8, 14, 3300, 520000, "Multi-cloud", "ITB", true},
	}
}

func inisialisasiPelatihan() {
	daftarPelatihan = []Pelatihan{
		//Cyber Security
		{"PLT001", "Ethical Hacking Dasar",            "Cyber Security", "Beginner",     "Regular", 1,  "Dr. Budi Santoso",      4.9, 350000, 20, 30, 18, Jadwal{"2026-07-10", "Kamis",  "19:00", 120}, "Pengenalan ethical hacking & keamanan dasar", "Kali Linux, Nmap, Wireshark", "Pemula IT"},
		{"PLT002", "Web Security Fundamentals",        "Cyber Security", "Intermediate", "Regular", 11, "Sarah Johnson",          4.7, 450000, 30, 25, 12, Jadwal{"2026-08-12", "Sabtu",  "09:00", 180}, "Keamanan aplikasi web modern", "OWASP Top 10, Burp Suite", "Web Developer"},
		{"PLT003", "Advanced Penetration Testing",     "Cyber Security", "Advanced",     "Regular", 1,  "Dr. Budi Santoso",      4.9, 700000, 60, 15,  9, Jadwal{"2026-09-18", "Jumat",  "08:00", 240}, "Penetrasi sistem & jaringan lanjutan", "Metasploit, Burp, OWASP", "Security Analyst"},
		{"PLT004", "Ethical Hacking Professional",     "Cyber Security", "Professional", "Regular", 23, "Eko Budiyanto",          4.6, 900000, 80, 10,  6, Jadwal{"2026-10-17", "Kamis",  "08:00", 300}, "Persiapan sertifikasi CEH & OSCP", "CEH, OSCP, Kali Linux", "Security Professional"},

		//WEB DEVELOPMENT
		{"PLT005", "HTML CSS JavaScript Dasar",        "Web Development", "Beginner",     "Regular", 21, "Reza Hakim, M.Kom",     4.5, 250000, 20, 50, 45, Jadwal{"2026-07-10", "Kamis",  "16:00",  90}, "Fondasi web development modern", "HTML5, CSS3, JavaScript ES6", "Pemula"},
		{"PLT006", "React JS Masterclass",             "Web Development", "Intermediate", "Regular", 2,  "Rina Wijayanti, M.Kom", 4.8, 400000, 40, 40, 35, Jadwal{"2026-08-11", "Jumat",  "18:00", 150}, "React modern dengan hooks & state management", "React, Redux, TypeScript", "Frontend Developer"},
		{"PLT007", "Vue.js Full Stack",                "Web Development", "Advanced",     "Regular", 12, "Michael Chen",           4.9, 600000, 50, 20, 16, Jadwal{"2026-09-19", "Sabtu",  "09:00", 180}, "Full stack dengan Vue 3 & Nuxt", "Vue 3, Nuxt, GraphQL, PostgreSQL", "Senior Developer"},
		{"PLT008", "Web Dev Expert & Architecture",    "Web Development", "Professional", "Regular", 12, "Michael Chen",           4.9, 950000, 80, 10,  5, Jadwal{"2026-10-14", "Senin",  "08:00", 300}, "Arsitektur sistem web skala enterprise", "Microservices, DDD, System Design", "Tech Lead"},

		// UI/UX DESIGN
		{"PLT009", "Figma UI Design Bootcamp",         "UI/UX Design", "Beginner",     "Regular", 3,  "Kevin Hartono",           4.7, 300000, 25, 35, 30, Jadwal{"2026-07-15", "Selasa", "10:00", 120}, "Desain antarmuka dengan Figma dari nol", "Figma, Auto Layout, Components", "Pemula Desainer"},
		{"PLT010", "Motion Design & Animation",        "UI/UX Design", "Intermediate", "Regular", 13, "Dewi Anggraini",          4.8, 420000, 35, 25, 20, Jadwal{"2026-08-12", "Sabtu",  "10:00", 120}, "Animasi UI yang profesional & menarik", "After Effects, Lottie, Principle", "UI Designer"},
		{"PLT011", "UX Research & Testing",            "UI/UX Design", "Advanced",     "Regular", 24, "Vanessa Olivia",          4.9, 550000, 45, 20, 15, Jadwal{"2026-09-14", "Senin",  "09:00", 150}, "Riset pengguna mendalam & usability testing", "Usability Testing, A/B Testing, Hotjar", "UX Researcher"},
		{"PLT012", "Design System & Product Design",   "UI/UX Design", "Professional", "Regular", 24, "Vanessa Olivia",          4.9, 880000, 75, 10,  4, Jadwal{"2026-10-12", "Sabtu",  "08:00", 300}, "Membangun design system level produk", "Storybook, Tokens, Design Ops", "Product Designer"},

		//ARTIFICIAL INTELLIGENCE 
		{"PLT013", "Pengenalan Artificial Intelligence","Artificial Intelligence", "Beginner",     "Regular", 4,  "Dr. Siti Rahayu",       4.9, 380000, 30, 35, 28, Jadwal{"2026-07-16", "Rabu",   "19:00", 120}, "Konsep dasar AI & machine learning", "Python, Jupyter, Scikit-learn", "Pemula Data"},
		{"PLT014", "Machine Learning dengan Python",   "Artificial Intelligence", "Intermediate", "Regular", 4,  "Dr. Siti Rahayu",       4.9, 550000, 50, 25, 20, Jadwal{"2026-08-10", "Kamis",  "08:00", 180}, "Membangun model ML end-to-end", "Python, Scikit-learn, XGBoost", "Data Enthusiast"},
		{"PLT015", "Deep Learning & Neural Networks",  "Artificial Intelligence", "Advanced",     "Regular", 14, "Prof. James Williams",  5.0, 800000, 60, 15, 11, Jadwal{"2026-09-20", "Minggu", "08:00", 240}, "Deep learning dengan TensorFlow & PyTorch", "TensorFlow, Keras, CNN, RNN", "ML Engineer"},
		{"PLT016", "AI Engineer Professional",         "Artificial Intelligence", "Professional", "Regular", 14, "Prof. James Williams",  5.0, 980000, 90, 10,  4, Jadwal{"2026-10-20", "Minggu", "08:00", 360}, "Persiapan sertifikasi AI profesional", "MLOps, LLM, AI System Design", "AI Professional"},

		// DATA SCIENCE 
		{"PLT017", "Python for Data Science",          "Data Science", "Beginner",     "Regular", 15, "Andika Pratama, M.Sc",  4.5, 280000, 25, 45, 40, Jadwal{"2026-07-11", "Jumat",  "08:00", 120}, "Python dasar untuk analisis data", "Python, Pandas, NumPy, Jupyter", "Pemula Data"},
		{"PLT018", "Data Analysis dengan Pandas",      "Data Science", "Intermediate", "Regular", 5,  "Ahmad Fauzi, M.T",      4.6, 320000, 30, 40, 33, Jadwal{"2026-08-16", "Rabu",   "19:00", 120}, "Analisis & visualisasi data profesional", "Pandas, Matplotlib, Seaborn, Plotly", "Data Analyst"},
		{"PLT019", "Tableau & Power BI",               "Data Science", "Advanced",     "Regular", 22, "Putri Maharani",         4.8, 430000, 35, 30, 24, Jadwal{"2026-09-16", "Rabu",   "08:00", 150}, "Visualisasi & business intelligence lanjutan", "Tableau, Power BI, DAX, SQL", "BI Analyst"},
		{"PLT020", "Data Scientist Professional",      "Data Science", "Professional", "Regular", 22, "Putri Maharani",         4.8, 870000, 80, 10,  5, Jadwal{"2026-10-16", "Rabu",   "08:00", 300}, "Sertifikasi Data Scientist profesional", "Statistics, ML, BigData, Spark", "Data Scientist"},

		// CLOUD COMPUTING 
		{"PLT021", "AWS Cloud Practitioner",           "Cloud Computing", "Beginner",     "Regular", 6,  "Lina Permata",          4.8, 480000, 40, 35, 29, Jadwal{"2026-07-12", "Sabtu",  "13:00", 150}, "Fondasi cloud computing dengan AWS", "AWS, EC2, S3, IAM, VPC", "IT Generalist"},
		{"PLT022", "Google Cloud Platform",            "Cloud Computing", "Intermediate", "Regular", 16, "Nadia Kusuma",          4.6, 520000, 45, 25, 19, Jadwal{"2026-08-17", "Kamis",  "19:00", 150}, "Deploy & kelola layanan di GCP", "GCP, BigQuery, GKE, Cloud Run", "Cloud Engineer"},
		{"PLT023", "Terraform Infrastructure as Code", "Cloud Computing", "Advanced",     "Regular", 25, "Dr. Irfan Maulana",     4.8, 650000, 55, 15, 11, Jadwal{"2026-09-15", "Selasa", "08:00", 180}, "Infrastruktur cloud sebagai kode", "Terraform, Ansible, Pulumi", "Cloud Architect"},
		{"PLT024", "Cloud Architect Professional",     "Cloud Computing", "Professional", "Regular", 25, "Dr. Irfan Maulana",     4.8, 960000, 90, 10,  4, Jadwal{"2026-10-15", "Selasa", "08:00", 360}, "Sertifikasi Cloud Architect profesional", "Multi-cloud, FinOps, SRE, Security", "Cloud Professional"},

		// ── NETWORKING ───────────────────────────────────────────────────────────
		{"PLT025", "Jaringan Komputer Dasar",          "Networking", "Beginner",     "Regular", 7,  "Denny Kurniawan",         4.5, 300000, 25, 35, 28, Jadwal{"2026-07-17", "Kamis",  "08:00", 120}, "Konsep dasar jaringan komputer", "OSI Model, TCP/IP, Cisco Packet Tracer", "Pemula Jaringan"},
		{"PLT026", "CCNA Networking Fundamental",      "Networking", "Intermediate", "Regular", 7,  "Denny Kurniawan",         4.5, 380000, 35, 30, 25, Jadwal{"2026-08-17", "Kamis",  "08:00", 180}, "Persiapan sertifikasi CCNA Cisco", "Cisco IOS, Routing, Switching, VLAN", "Network Admin"},
		{"PLT027", "Network Security & Firewall",      "Networking", "Advanced",     "Regular", 17, "Yudi Wibowo",             4.7, 480000, 40, 20, 17, Jadwal{"2026-09-13", "Minggu", "09:00", 180}, "Keamanan jaringan & manajemen firewall", "pfSense, Snort, IDS/IPS, VPN", "Network Security"},
		{"PLT028", "Network Professional CCIE",        "Networking", "Professional", "Regular", 17, "Yudi Wibowo",             4.7, 920000, 90, 10,  3, Jadwal{"2026-10-13", "Minggu", "08:00", 360}, "Persiapan sertifikasi CCIE Cisco", "BGP, MPLS, SDN, Network Automation", "Network Expert"},

		// BLOCKCHAIN 
		{"PLT029", "Pengenalan Blockchain & Crypto",   "Blockchain", "Beginner",     "Regular", 8,  "Mega Puspita, M.Sc",      4.7, 320000, 25, 30, 22, Jadwal{"2026-07-13", "Minggu", "13:00", 120}, "Konsep dasar blockchain & cryptocurrency", "Bitcoin, Ethereum, Wallet, DeFi", "Pemula Crypto"},
		{"PLT030", "Blockchain & Ethereum Dev",        "Blockchain", "Intermediate", "Regular", 8,  "Mega Puspita, M.Sc",      4.7, 520000, 45, 20, 14, Jadwal{"2026-08-13", "Minggu", "10:00", 150}, "Pengembangan smart contract Ethereum", "Solidity, Hardhat, Ethers.js", "Crypto Enthusiast"},
		{"PLT031", "Web3 & DeFi Development",          "Blockchain", "Advanced",     "Regular", 18, "Clara Situmorang",        4.6, 650000, 55, 15, 10, Jadwal{"2026-09-20", "Minggu", "13:00", 180}, "Pengembangan protokol DeFi & NFT", "Web3.js, DeFi, NFT, IPFS", "Blockchain Developer"},
		{"PLT032", "Blockchain Architect Professional","Blockchain", "Professional", "Regular", 18, "Clara Situmorang",        4.6, 900000, 80, 10,  4, Jadwal{"2026-10-20", "Minggu", "08:00", 300}, "Desain sistem blockchain skala enterprise", "Layer2, ZKP, Cross-chain, DAO", "Blockchain Architect"},

		// MOBILE DEVELOPMENT 
		{"PLT033", "Flutter Mobile App Dasar",         "Mobile Development", "Beginner",     "Regular", 9,  "Rizal Firdaus",         4.6, 350000, 30, 40, 36, Jadwal{"2026-07-11", "Jumat",  "19:00", 120}, "Buat aplikasi mobile pertama dengan Flutter", "Flutter, Dart, Firebase, Provider", "Pemula Mobile"},
		{"PLT034", "React Native Cross Platform",      "Mobile Development", "Intermediate", "Regular", 19, "Fajar Nugroho",         4.9, 480000, 40, 30, 26, Jadwal{"2026-08-18", "Jumat",  "18:00", 120}, "Aplikasi mobile cross-platform modern", "React Native, Expo, Redux", "Mobile Developer"},
		{"PLT035", "iOS Swift Development",            "Mobile Development", "Advanced",     "Regular", 19, "Fajar Nugroho",         4.9, 620000, 55, 15, 12, Jadwal{"2026-09-14", "Senin",  "08:00", 180}, "Pengembangan aplikasi iOS native", "Swift, SwiftUI, Xcode, CoreData", "iOS Developer"},
		{"PLT036", "Mobile Architect Professional",    "Mobile Development", "Professional", "Regular", 19, "Fajar Nugroho",         4.9, 940000, 85, 10,  4, Jadwal{"2026-10-18", "Jumat",  "08:00", 360}, "Arsitektur & optimasi aplikasi mobile", "Clean Architecture, CI/CD Mobile, App Store", "Mobile Architect"},

		// DEVOPS 
		{"PLT037", "Linux & Shell Scripting Dasar",    "DevOps", "Beginner",     "Regular", 10, "Hendra Saputra",          4.8, 290000, 25, 40, 32, Jadwal{"2026-07-15", "Selasa", "19:00", 120}, "Dasar Linux & otomasi shell scripting", "Linux, Bash, Cron, SSH", "IT Pemula"},
		{"PLT038", "Docker & Kubernetes",              "DevOps", "Intermediate", "Regular", 10, "Hendra Saputra",          4.8, 500000, 40, 25, 21, Jadwal{"2026-08-15", "Selasa", "20:00", 120}, "Containerisasi & orkestrasi modern", "Docker, Kubernetes, Helm, Registry", "DevOps Engineer"},
		{"PLT039", "CI/CD & GitOps",                  "DevOps", "Advanced",     "Regular", 20, "Bella Anindita",          4.7, 580000, 50, 20, 15, Jadwal{"2026-09-19", "Sabtu",  "13:00", 150}, "Pipeline CI/CD & GitOps workflow", "Jenkins, GitHub Actions, ArgoCD", "DevOps Lead"},
		{"PLT040", "DevOps Engineer Professional",     "DevOps", "Professional", "Regular", 10, "Hendra Saputra",          4.8, 850000, 80, 10,  7, Jadwal{"2026-10-18", "Jumat",  "08:00", 300}, "Sertifikasi DevOps & SRE profesional", "K8s CKA, Terraform, Prometheus, SRE", "Senior DevOps"},
	}	
}



func inputString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(prompt string) int {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(ColorRed + "  [!] Input harus berupa angka." + ColorReset)
			continue
		}
		return val
	}
}

func formatRupiah(jumlah int) string {
	str := strconv.Itoa(jumlah)
	result := ""
	for i, c := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return "Rp " + result
}

func loading(pesan string) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	fmt.Print(ColorCyan)
	for i := 0; i < 15; i++ {
		fmt.Printf("\r  %s %s...", frames[i%len(frames)], pesan)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Println("\r  ✓ " + pesan + " selesai." + ColorReset)
}

func cetakGaris(panjang int, karakter string) {
	fmt.Println(ColorBlue + strings.Repeat(karakter, panjang) + ColorReset)
}

func cetakHeader(judul string) {
	lebar := 70
	cetakGaris(lebar, "═")
	padding := (lebar - len(judul)) / 2
	fmt.Printf("%s%s%s%s%s\n", ColorBold+ColorCyan, strings.Repeat(" ", padding), judul, strings.Repeat(" ", padding), ColorReset)
	cetakGaris(lebar, "═")
}

func generateID(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s-%04d%02d", prefix, rand.Intn(9000)+1000, rand.Intn(99)+1)
}

func bintangRating(rating float64) string {
	full := int(rating)
	stars := strings.Repeat("★", full)
	if rating-float64(full) >= 0.5 {
		stars += "½"
	}
	empty := 5 - full
	if rating-float64(full) >= 0.5 {
		empty--
	}
	stars += strings.Repeat("☆", empty)
	return stars
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}


func tampilkanTabelPelatihan(daftar []Pelatihan) {
	if len(daftar) == 0 {
		fmt.Println(ColorYellow + "  Tidak ada pelatihan yang ditemukan." + ColorReset)
		return
	}
	fmt.Printf("\n%s%-8s %-38s %-16s %-12s %-8s %-10s%s\n",
		ColorBold+ColorCyan, "ID", "Nama Pelatihan", "Bidang", "Level", "Rating", "Harga", ColorReset)
	cetakGaris(100, "─")
	for _, p := range daftar {
		fmt.Printf("%-8s %-38s %-16s %-12s %-8.1f %-10s\n",
			p.ID, p.Nama, p.Bidang[:min(16, len(p.Bidang))], p.Level, p.Rating, formatRupiah(p.Harga))
	}
	cetakGaris(100, "─")
}

func tampilkanDetailPelatihan(p Pelatihan) {
	cetakGaris(100, "─")
	fmt.Printf("  %sID Pelatihan  :%s %s\n", ColorBold, ColorReset, p.ID)
	fmt.Printf("  %sNama         :%s %s\n", ColorBold, ColorReset, p.Nama)
	fmt.Printf("  %sBidang       :%s %s\n", ColorBold, ColorReset, p.Bidang)
	fmt.Printf("  %sLevel        :%s %s\n", ColorBold, ColorReset, p.Level)
	fmt.Printf("  %sMentor       :%s %s\n", ColorBold, ColorReset, p.MentorNama)
	fmt.Printf("  %sRating       :%s %.1f %s\n", ColorBold, ColorReset, p.Rating, bintangRating(p.Rating))
	fmt.Printf("  %sHarga        :%s %s\n", ColorBold, ColorReset, formatRupiah(p.Harga))
	fmt.Printf("  %sKuota        :%s %d (sisa: %d)\n", ColorBold, ColorReset, p.Kuota, p.Kuota-p.Terdaftar)
	fmt.Printf("  %sJadwal       :%s %s, %s pukul %s\n", ColorBold, ColorReset, p.Jadwal.Hari, p.Jadwal.Tanggal, p.Jadwal.Jam)
	cetakGaris(100, "─")
}

func tampilkanTabelMentor(daftar []Mentor) {
	if len(daftar) == 0 {
		fmt.Println(ColorYellow + "  Tidak ada mentor yang ditemukan." + ColorReset)
		return
	}
	fmt.Printf("\n%s%-4s %-28s %-22s %-7s %-5s %-8s%s\n",
		ColorBold+ColorCyan, "ID", "Nama Mentor", "Bidang", "Rating", "Exp", "Siswa", ColorReset)
	cetakGaris(80, "─")
	for _, m := range daftar {
		bidang := m.Bidang
		if len(bidang) > 22 {
			bidang = bidang[:22]
		}
		fmt.Printf("%-4d %-28s %-22s %-7.1f %-5d %-8d\n",
			m.ID, m.Nama, bidang, m.Rating, m.Pengalaman, m.JumlahSiswa)
	}
	cetakGaris(80, "─")
}

func cetakRingkasanPendaftaran(r RiwayatPendaftaran) {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s RINGKASAN PENDAFTARAN%s\n", ColorBold, ColorGreen, ColorReset)
	cetakGaris(70, "─")
	fmt.Printf("  %-20s : %d\n", "No. Pendaftaran", r.ID)
	fmt.Printf("  %-20s : %s\n", "Peserta", r.PesertaNama)
	fmt.Printf("  %-20s : %s\n", "Pelatihan", r.PelatihanNama)
	fmt.Printf("  %-20s : %s\n", "Mentor", r.MentorNama)
	fmt.Printf("  %-20s : %s\n", "Kategori", r.Kategori)
	fmt.Printf("  %-20s : %s\n", "Tanggal Daftar", r.TanggalDaftar)
	fmt.Printf("  %-20s : %s\n", "Total Biaya", formatRupiah(r.TotalBiaya))
	fmt.Printf("  %-20s : %s\n", "Status", r.Status)
	cetakGaris(70, "═")
}

func tampilkanSertifikat(s Sertifikat) {
	fmt.Println()
	cetakGaris(70, "╔")
	fmt.Println("  ╔" + strings.Repeat("═", 68) + "╗")
	fmt.Printf("  ║%s%s%-68s%s║\n", ColorBold+ColorYellow, strings.Repeat(" ", 20), "SERTIFIKAT KELULUSAN", ColorReset)
	fmt.Println("  ╠" + strings.Repeat("═", 68) + "╣")
	fmt.Printf("  ║  %-30s : %-33s║\n", "No. Sertifikat", s.ID)
	fmt.Printf("  ║  %-30s : %-33s║\n", "Nama Peserta", s.PesertaNama)
	fmt.Printf("  ║  %-30s : %-33s║\n", "Program Pelatihan", s.PelatihanNama[:min(33, len(s.PelatihanNama))])
	fmt.Printf("  ║  %-30s : %-33s║\n", "Mentor", s.MentorNama[:min(33, len(s.MentorNama))])
	fmt.Printf("  ║  %-30s : %-33s║\n", "Tanggal Selesai", s.TanggalSelesai)
	fmt.Printf("  ║  %-30s : %-33s║\n", "Grade", s.Grade)
	fmt.Println("  ╚" + strings.Repeat("═", 68) + "╝")
}



func tampilkanMenuUtama() {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s SISTEM MANAJEMEN PELATIHAN TEKNOLOGI %s\n", ColorBold+ColorCyan, "▸▸", ColorReset)
	cetakGaris(70, "─")
	menus := []string{
		" 1. Tambah Data",
		" 2. Edit Data",
		" 3. Hapus Data",
		" 4. Tampilkan Data",
		" 0. Keluar",
	}
	for _, m := range menus {
		fmt.Printf("  %s\n", m)
	}
	cetakGaris(70, "─")
	fmt.Print("  Masukkan pilihan menu (0-4): ")
}

func tampilkanMenuTambahData() {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s TAMBAH DATA %s\n", ColorBold+ColorCyan, "▸▸", ColorReset)
	cetakGaris(70, "─")
	menus := []string{
		" 1. Cari & Daftar Pelatihan",
		" 2. Lihat Semua Pelatihan",
		" 3. Cari Mentor",
		" 4. Statistik Pelatihan",
		" 5. Lihat Semua Mentor",
		" 0. Kembali ke Menu Utama",
	}
	for _, m := range menus {
		fmt.Printf("  %s\n", m)
	}
	cetakGaris(70, "─")
	fmt.Print("  Masukkan pilihan (0-5): ")
}

func tampilkanMenuEditData() {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s EDIT DATA %s\n", ColorBold+ColorCyan, "▸▸", ColorReset)
	cetakGaris(70, "─")
	menus := []string{
		" 1. Edit Nama Peserta",
		" 2. Edit Email Peserta",
		" 3. Edit No. HP Peserta",
		" 4. Edit Kota Peserta",
		" 5. Edit Level Peserta",
		" 6. Edit Bidang Minat Peserta",
		" 0. Kembali ke Menu Utama",
	}
	for _, m := range menus {
		fmt.Printf("  %s\n", m)
	}
	cetakGaris(70, "─")
	fmt.Print("  Masukkan pilihan (0-6): ")
}

func tampilkanMenuHapusData() {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s HAPUS DATA %s\n", ColorBold+ColorCyan, "▸▸", ColorReset)
	cetakGaris(70, "─")
	menus := []string{
		" 1. Hapus Data Peserta",
		" 2. Hapus Pendaftaran Pelatihan",
		" 0. Kembali ke Menu Utama",
	}
	for _, m := range menus {
		fmt.Printf("  %s\n", m)
	}
	cetakGaris(70, "─")
	fmt.Print("  Masukkan pilihan (0-2): ")
}

func tampilkanMenuTampilkanData() {
	fmt.Println()
	cetakGaris(70, "═")
	fmt.Printf("  %s%s TAMPILKAN DATA %s\n", ColorBold+ColorCyan, "▸▸", ColorReset)
	cetakGaris(70, "─")
	menus := []string{
		" 1. Data Peserta",
		" 2. Riwayat Pendaftaran",
		" 3. Sertifikat Peserta",
		" 0. Kembali ke Menu Utama",
	}
	for _, m := range menus {
		fmt.Printf("  %s\n", m)
	}
	cetakGaris(70, "─")
	fmt.Print("  Masukkan pilihan (0-3): ")
}


func menuCariPelatihan() {
	cetakHeader(" CARI & DAFTAR PELATIHAN")
	
	fmt.Println("\n  Langkah 1: Pilih Bidang Teknologi")
	for _, b := range daftarBidang {
		fmt.Printf("  %s%d.%s %s %s\n", ColorCyan, b.ID, ColorReset, b.IconASCII, b.Nama)
	}
	
	pilihanBidang := inputInt("\n  Pilih nomor bidang (1-10): ")
	if pilihanBidang < 1 || pilihanBidang > len(daftarBidang) {
		fmt.Println(ColorRed + "  [!] Pilihan bidang tidak tersedia." + ColorReset)
		return
	}
	bidangDipilih := daftarBidang[pilihanBidang-1].Nama
	
	fmt.Println("\n  Langkah 2: Pilih Level Pelatihan")
	levels := []string{"Beginner", "Intermediate", "Advanced", "Professional"}
	for i, l := range levels {
		fmt.Printf("  %s%d.%s %s\n", ColorCyan, i+1, ColorReset, l)
	}
	pilihanLevel := inputInt("\n  Pilih level (1-4): ")
	if pilihanLevel < 1 || pilihanLevel > 4 {
		fmt.Println(ColorRed + "  [!] Pilihan level tidak valid." + ColorReset)
		return
	}
	levelDipilih := levels[pilihanLevel-1]
	
	// Filter pelatihan
	var hasilFilter []Pelatihan
	for _, p := range daftarPelatihan {
		if p.Bidang == bidangDipilih && p.Level == levelDipilih {
			hasilFilter = append(hasilFilter, p)
		}
	}
	
	if len(hasilFilter) == 0 {
		fmt.Println(ColorYellow + "\n  Tidak ada kelas tersedia untuk bidang dan level ini." + ColorReset)
		return
	}
	
	tampilkanTabelPelatihan(hasilFilter)
	
	pilihanKelas := inputInt("\n  Masukkan nomor urut kelas yang dipilih: ")
	if pilihanKelas < 1 || pilihanKelas > len(hasilFilter) {
		fmt.Println(ColorRed + "  [!] Pilihan tidak valid." + ColorReset)
		return
	}
	kelasDipilih := hasilFilter[pilihanKelas-1]
	
	tampilkanDetailPelatihan(kelasDipilih)
	
	// Proses pendaftaran
	fmt.Println("\n  Langkah 3: Masukkan Data Peserta")
	
	// Cari peserta berdasarkan nama
	namaPeserta := inputString("  Nama lengkap Anda: ")
	if namaPeserta == "" {
		fmt.Println(ColorRed + "   Nama tidak boleh kosong!" + ColorReset)
		return
	}
	
	// Cek apakah peserta sudah terdaftar
	var pesertaID int
	var pesertaExist bool = false
	
	for _, p := range daftarPeserta {
		if strings.EqualFold(p.Nama, namaPeserta) {
			pesertaID = p.ID
			pesertaExist = true
			break
		}
	}
	
	// Jika belum terdaftar, buat peserta baru
	if !pesertaExist {
		fmt.Println("\n  " + ColorYellow + " Data peserta belum terdaftar. Silakan lengkapi:" + ColorReset)
		email := inputString("  Email : ")
		noHP := inputString("  No. HP: ")
		kota := inputString("  Kota  : ")
		
		fmt.Println("\n  Pilih Level:")
		fmt.Println("  1. Beginner\n  2. Intermediate\n  3. Advanced\n  4. Professional")
		levelPilihan := inputInt("  Level (1-4): ")
		levelMap := map[int]string{1: "Beginner", 2: "Intermediate", 3: "Advanced", 4: "Professional"}
		level := levelMap[levelPilihan]
		if level == "" {
			level = "Beginner"
		}
		
		pesertaBaru := Peserta{
			ID:     pesertaIDCounter,
			Nama:   namaPeserta,
			Email:  email,
			NoHP:   noHP,
			Kota:   kota,
			Level:  level,
			Bidang: bidangDipilih,
		}
		daftarPeserta = append(daftarPeserta, pesertaBaru)
		pesertaID = pesertaIDCounter
		pesertaIDCounter++
		fmt.Printf("\n  %s Data peserta '%s' berhasil didaftarkan! (ID: %d)%s\n", 
			ColorGreen, namaPeserta, pesertaID, ColorReset)
	} else {
		fmt.Printf("\n  %s Selamat datang kembali, %s! (ID: %d)%s\n", 
			ColorGreen, namaPeserta, pesertaID, ColorReset)
	}
	
	konfirmasi := inputString("\n  Konfirmasi pendaftaran? (y/t): ")
	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println("  Pendaftaran dibatalkan.")
		return
	}
	
	// Simpan riwayat pendaftaran
	riwayat := RiwayatPendaftaran{
		ID:             riwayatIDCounter,
		PesertaID:      pesertaID,
		PesertaNama:    namaPeserta,
		PelatihanID:    kelasDipilih.ID,
		PelatihanNama:  kelasDipilih.Nama,
		MentorNama:     kelasDipilih.MentorNama,
		Kategori:       "Regular",
		TanggalDaftar:  time.Now().Format("2006-01-02"),
		TanggalMulai:   kelasDipilih.Jadwal.Tanggal,
		TanggalSelesai: kelasDipilih.Jadwal.Tanggal,
		TotalBiaya:     kelasDipilih.Harga,
		Status:         "Aktif",
	}
	daftarRiwayat = append(daftarRiwayat, riwayat)
	riwayatIDCounter++
	
	// Buat sertifikat
	sertifikat := Sertifikat{
		ID:             generateID("CERT"),
		PesertaID:      pesertaID,
		PesertaNama:    namaPeserta,
		PelatihanID:    kelasDipilih.ID,
		PelatihanNama:  kelasDipilih.Nama,
		MentorNama:     kelasDipilih.MentorNama,
		TanggalSelesai: time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
		Grade:          "Pending",
		Valid:          true,
	}
	daftarSertifikat = append(daftarSertifikat, sertifikat)
	
	loading("Memproses pendaftaran")
	cetakRingkasanPendaftaran(riwayat)
	fmt.Printf("\n  %sPENDAFTARAN BERHASIL!%s\n", ColorGreen, ColorReset)
	fmt.Printf("  Nomor Sertifikat: %s%s%s\n", ColorYellow, sertifikat.ID, ColorReset)
}

func menuLihatSemuaPelatihan() {
	cetakHeader("DAFTAR SEMUA PELATIHAN")
	tampilkanTabelPelatihan(daftarPelatihan)
}

func menuCariMentor() {
	cetakHeader("CARI MENTOR")
	keyword := inputString("  Masukkan nama atau bidang mentor: ")
	
	var hasil []Mentor
	for _, m := range daftarMentor {
		if strings.Contains(strings.ToLower(m.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(m.Bidang), strings.ToLower(keyword)) {
			hasil = append(hasil, m)
		}
	}
	
	if len(hasil) == 0 {
		fmt.Println(ColorYellow + "  Tidak ada mentor yang ditemukan." + ColorReset)
		return
	}
	
	tampilkanTabelMentor(hasil)
}

func menuLihatSemuaMentor() {
	cetakHeader("DAFTAR SEMUA MENTOR")
	tampilkanTabelMentor(daftarMentor)
}

func menuStatistik() {
	cetakHeader("STATISTIK")
	fmt.Printf("\n  Total Peserta    : %d\n", len(daftarPeserta))
	fmt.Printf("  Total Pelatihan  : %d\n", len(daftarPelatihan))
	fmt.Printf("  Total Mentor     : %d\n", len(daftarMentor))
	fmt.Printf("  Total Pendaftaran: %d\n", len(daftarRiwayat))
	fmt.Printf("  Total Sertifikat : %d\n", len(daftarSertifikat))
}


func pilihPesertaByID() (int, *Peserta) {
	if len(daftarPeserta) == 0 {
		fmt.Println(ColorYellow + "Belum ada data peserta!" + ColorReset)
		fmt.Println("Silakan daftar pelatihan terlebih dahulu.")
		return -1, nil
	}
	
	menuDataPeserta()
	id := inputInt("\n  Masukkan ID peserta: ")
	
	for i := range daftarPeserta {
		if daftarPeserta[i].ID == id {
			return i, &daftarPeserta[i]
		}
	}
	fmt.Println(ColorRed + "   ID peserta tidak ditemukan." + ColorReset)
	return -1, nil
}

func menuEditNamaPeserta() {
	cetakHeader("EDIT NAMA PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  Nama saat ini: %s%s%s\n", ColorCyan, peserta.Nama, ColorReset)
	namaBaru := inputString("  Nama baru: ")
	if namaBaru == "" {
		fmt.Println(ColorRed + "   Nama tidak boleh kosong!" + ColorReset)
		return
	}
	
	namaLama := daftarPeserta[idx].Nama
	daftarPeserta[idx].Nama = namaBaru
	
	// Update di riwayat
	for i := range daftarRiwayat {
		if daftarRiwayat[i].PesertaNama == namaLama {
			daftarRiwayat[i].PesertaNama = namaBaru
		}
	}
	
	// Update di sertifikat
	for i := range daftarSertifikat {
		if daftarSertifikat[i].PesertaNama == namaLama {
			daftarSertifikat[i].PesertaNama = namaBaru
		}
	}
	
	loading("Menyimpan perubahan")
	fmt.Printf("  %s Nama berhasil diubah menjadi '%s'%s\n", ColorGreen, namaBaru, ColorReset)
}

func menuEditEmailPeserta() {
	cetakHeader("EDIT EMAIL PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  Email saat ini: %s%s%s\n", ColorCyan, peserta.Email, ColorReset)
	emailBaru := inputString("  Email baru: ")
	if emailBaru == "" {
		fmt.Println(ColorRed + "   Email tidak boleh kosong!" + ColorReset)
		return
	}
	
	daftarPeserta[idx].Email = emailBaru
	loading("Menyimpan perubahan")
	fmt.Printf("  %s Email berhasil diubah menjadi '%s'%s\n", ColorGreen, emailBaru, ColorReset)
}

func menuEditNoHPPeserta() {
	cetakHeader(" EDIT NO. HP PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  No. HP saat ini: %s%s%s\n", ColorCyan, peserta.NoHP, ColorReset)
	hpBaru := inputString("  No. HP baru: ")
	if hpBaru == "" {
		fmt.Println(ColorRed + "   No. HP tidak boleh kosong!" + ColorReset)
		return
	}
	
	daftarPeserta[idx].NoHP = hpBaru
	loading("Menyimpan perubahan")
	fmt.Printf("  %s No. HP berhasil diubah menjadi '%s'%s\n", ColorGreen, hpBaru, ColorReset)
}

func menuEditKotaPeserta() {
	cetakHeader(" EDIT KOTA PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  Kota saat ini: %s%s%s\n", ColorCyan, peserta.Kota, ColorReset)
	kotaBaru := inputString("  Kota baru: ")
	if kotaBaru == "" {
		fmt.Println(ColorRed + "   Kota tidak boleh kosong!" + ColorReset)
		return
	}
	
	daftarPeserta[idx].Kota = kotaBaru
	loading("Menyimpan perubahan")
	fmt.Printf("  %s Kota berhasil diubah menjadi '%s'%s\n", ColorGreen, kotaBaru, ColorReset)
}

func menuEditLevelPeserta() {
	cetakHeader(" EDIT LEVEL PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  Level saat ini: %s%s%s\n", ColorCyan, peserta.Level, ColorReset)
	fmt.Println("\n  Pilih Level baru:")
	fmt.Println("  1. Beginner\n  2. Intermediate\n  3. Advanced\n  4. Professional")
	
	levelPilihan := inputInt("  Level (1-4): ")
	levelMap := map[int]string{1: "Beginner", 2: "Intermediate", 3: "Advanced", 4: "Professional"}
	levelBaru := levelMap[levelPilihan]
	
	if levelBaru == "" {
		fmt.Println(ColorRed + "   Pilihan tidak valid!" + ColorReset)
		return
	}
	
	daftarPeserta[idx].Level = levelBaru
	loading("Menyimpan perubahan")
	fmt.Printf("  %s Level berhasil diubah menjadi '%s'%s\n", ColorGreen, levelBaru, ColorReset)
}

func menuEditBidangPeserta() {
	cetakHeader(" EDIT BIDANG MINAT PESERTA")
	idx, peserta := pilihPesertaByID()
	if idx == -1 {
		return
	}
	
	fmt.Printf("\n  Bidang minat saat ini: %s%s%s\n", ColorCyan, peserta.Bidang, ColorReset)
	fmt.Println("\n  Pilih Bidang Minat baru:")
	for _, b := range daftarBidang {
		fmt.Printf("  %d. %s %s\n", b.ID, b.IconASCII, b.Nama)
	}
	
	bidangPilihan := inputInt("  Bidang (1-10): ")
	if bidangPilihan < 1 || bidangPilihan > len(daftarBidang) {
		fmt.Println(ColorRed + "   Pilihan tidak valid!" + ColorReset)
		return
	}
	
	bidangBaru := daftarBidang[bidangPilihan-1].Nama
	daftarPeserta[idx].Bidang = bidangBaru
	loading("Menyimpan perubahan")
	fmt.Printf("  %s Bidang minat berhasil diubah menjadi '%s'%s\n", ColorGreen, bidangBaru, ColorReset)
}


func menuHapusDataPeserta() {
	cetakHeader("HAPUS DATA PESERTA")
	
	if len(daftarPeserta) == 0 {
		fmt.Println(ColorYellow + " Belum ada data peserta!" + ColorReset)
		return
	}
	
	menuDataPeserta()
	id := inputInt("\n  Masukkan ID peserta yang akan dihapus: ")
	
	idx := -1
	for i, p := range daftarPeserta {
		if p.ID == id {
			idx = i
			break
		}
	}
	
	if idx == -1 {
		fmt.Println(ColorRed + "   ID peserta tidak ditemukan." + ColorReset)
		return
	}
	
	peserta := daftarPeserta[idx]
	fmt.Printf("\n  %s PERINGATAN! Data berikut akan dihapus:%s\n", ColorRed, ColorReset)
	fmt.Printf("  ┌─────────────────────────────────────┐\n")
	fmt.Printf("  │ ID      : %d\n", peserta.ID)
	fmt.Printf("  │ Nama    : %s\n", peserta.Nama)
	fmt.Printf("  │ Email   : %s\n", peserta.Email)
	fmt.Printf("  │ Kota    : %s\n", peserta.Kota)
	fmt.Printf("  └─────────────────────────────────────┘\n")
	
	konfirmasi := inputString("\n  Yakin ingin menghapus? (ya/tidak): ")
	if strings.ToLower(konfirmasi) != "ya" {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}
	
	namaHapus := daftarPeserta[idx].Nama
	daftarPeserta = append(daftarPeserta[:idx], daftarPeserta[idx+1:]...)
	
	loading("Menghapus data peserta")
	fmt.Printf("  %s Data peserta '%s' berhasil dihapus!%s\n", ColorGreen, namaHapus, ColorReset)
}

func menuHapusPendaftaran() {
	cetakHeader("HAPUS PENDAFTARAN")
	
	if len(daftarRiwayat) == 0 {
		fmt.Println(ColorYellow + "   Belum ada riwayat pendaftaran!" + ColorReset)
		return
	}
	
	menuRiwayatPendaftaran()
	id := inputInt("\n  Masukkan ID pendaftaran yang akan dihapus: ")
	
	idx := -1
	for i, r := range daftarRiwayat {
		if r.ID == id {
			idx = i
			break
		}
	}
	
	if idx == -1 {
		fmt.Println(ColorRed + "   ID pendaftaran tidak ditemukan." + ColorReset)
		return
	}
	
	r := daftarRiwayat[idx]
	fmt.Printf("\n  Pendaftaran yang akan dihapus:\n")
	fmt.Printf("  Peserta   : %s\n", r.PesertaNama)
	fmt.Printf("  Pelatihan : %s\n", r.PelatihanNama)
	
	konfirmasi := inputString("\n  Yakin ingin menghapus? (ya/tidak): ")
	if strings.ToLower(konfirmasi) != "ya" {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}
	
	daftarRiwayat = append(daftarRiwayat[:idx], daftarRiwayat[idx+1:]...)
	loading("Menghapus pendaftaran")
	fmt.Printf("  %s Pendaftaran berhasil dihapus!%s\n", ColorGreen, ColorReset)
}


func menuDataPeserta() {
	cetakHeader(" DATA PESERTA")
	
	if len(daftarPeserta) == 0 {
		fmt.Println(ColorYellow + "  Belum ada data peserta." + ColorReset)
		fmt.Println("   Silakan daftar pelatihan terlebih dahulu.")
		return
	}
	
	fmt.Printf("\n  %s%-5s %-22s %-25s %-12s %-14s %-20s%s\n",
		ColorBold+ColorCyan, "ID", "Nama", "Email", "Kota", "Level", "Bidang Minat", ColorReset)
	cetakGaris(100, "─")
	
	for _, p := range daftarPeserta {
		fmt.Printf("  %-5d %-22s %-25s %-12s %-14s %-20s\n",
			p.ID,
			truncate(p.Nama, 22),
			truncate(p.Email, 25),
			truncate(p.Kota, 12),
			truncate(p.Level, 14),
			truncate(p.Bidang, 20))
	}
	cetakGaris(100, "─")
	fmt.Printf("  Total: %d peserta\n", len(daftarPeserta))
}

func menuRiwayatPendaftaran() {
	cetakHeader(" RIWAYAT PENDAFTARAN")
	
	if len(daftarRiwayat) == 0 {
		fmt.Println(ColorYellow + "  Belum ada riwayat pendaftaran." + ColorReset)
		return
	}
	
	fmt.Printf("\n  %s%-5s %-20s %-30s %-10s %-12s %-10s%s\n",
		ColorBold+ColorCyan, "ID", "Peserta", "Pelatihan", "Kategori", "Tanggal", "Status", ColorReset)
	cetakGaris(93, "─")
	
	for _, r := range daftarRiwayat {
		fmt.Printf("  %-5d %-20s %-30s %-10s %-12s %-10s\n",
			r.ID,
			truncate(r.PesertaNama, 20),
			truncate(r.PelatihanNama, 30),
			r.Kategori,
			r.TanggalDaftar,
			r.Status)
	}
	cetakGaris(93, "─")
}

func menuSertifikat() {
	cetakHeader(" SERTIFIKAT PESERTA")
	
	if len(daftarSertifikat) == 0 {
		fmt.Println(ColorYellow + "  Belum ada sertifikat." + ColorReset)
		return
	}
	
	fmt.Printf("\n  %s%-12s %-20s %-30s %-10s%s\n",
		ColorBold+ColorCyan, "No. Sertifikat", "Peserta", "Pelatihan", "Grade", ColorReset)
	cetakGaris(80, "─")
	
	for _, s := range daftarSertifikat {
		fmt.Printf("  %-12s %-20s %-30s %-10s\n",
			s.ID,
			truncate(s.PesertaNama, 20),
			truncate(s.PelatihanNama, 30),
			s.Grade)
	}
	cetakGaris(80, "─")
}


func main() {
	// Inisialisasi data dummy (hanya mentor dan pelatihan)
	inisialisasiMentor()
	inisialisasiPelatihan()
	
	// Data peserta dimulai KOSONG
	
	fmt.Println(ColorBold + ColorCyan)
	fmt.Println("  ╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║             Sistem Manajemen Pelatihan Teknologi                 ║")
	fmt.Println("  ║        Platform Pelatihan Online Terdepan di Indonesia           ║")
	fmt.Println("  ╚══════════════════════════════════════════════════════════════════╝")
	fmt.Println(ColorReset)
	time.Sleep(500 * time.Millisecond)
	
	for {
		tampilkanMenuUtama()
		pilihan := inputInt("")
		
		switch pilihan {
		case 1: // TAMBAH DATA
			for {
				tampilkanMenuTambahData()
				subPilihan := inputInt("")
				
				if subPilihan == 0 {
					break
				}
				
				switch subPilihan {
				case 1:
					menuCariPelatihan()
				case 2:
					menuLihatSemuaPelatihan()
				case 3:
					menuCariMentor()
				case 4:
					menuStatistik()
				case 5:
					menuLihatSemuaMentor()
				default:
					fmt.Println(ColorRed + "\n  [!] Pilihan tidak tersedia.\n" + ColorReset)
				}
				fmt.Print(ColorYellow + "\n  [Tekan ENTER untuk kembali...]" + ColorReset)
				reader.ReadString('\n')
			}
			
		case 2: // EDIT DATA
			for {
				tampilkanMenuEditData()
				subPilihan := inputInt("")
				
				if subPilihan == 0 {
					break
				}
				
				switch subPilihan {
				case 1:
					menuEditNamaPeserta()
				case 2:
					menuEditEmailPeserta()
				case 3:
					menuEditNoHPPeserta()
				case 4:
					menuEditKotaPeserta()
				case 5:
					menuEditLevelPeserta()
				case 6:
					menuEditBidangPeserta()
				default:
					fmt.Println(ColorRed + "\n  [!] Pilihan tidak tersedia.\n" + ColorReset)
				}
				fmt.Print(ColorYellow + "\n  [Tekan ENTER untuk kembali...]" + ColorReset)
				reader.ReadString('\n')
			}
			
		case 3: // HAPUS DATA
			for {
				tampilkanMenuHapusData()
				subPilihan := inputInt("")
				
				if subPilihan == 0 {
					break
				}
				
				switch subPilihan {
				case 1:
					menuHapusDataPeserta()
				case 2:
					menuHapusPendaftaran()
				default:
					fmt.Println(ColorRed + "\n  [!] Pilihan tidak tersedia.\n" + ColorReset)
				}
				fmt.Print(ColorYellow + "\n  [Tekan ENTER untuk kembali...]" + ColorReset)
				reader.ReadString('\n')
			}
			
		case 4: // TAMPILKAN DATA
			for {
				tampilkanMenuTampilkanData()
				subPilihan := inputInt("")
				
				if subPilihan == 0 {
					break
				}
				
				switch subPilihan {
				case 1:
					menuDataPeserta()
				case 2:
					menuRiwayatPendaftaran()
				case 3:
					menuSertifikat()
				default:
					fmt.Println(ColorRed + "\n  [!] Pilihan tidak tersedia.\n" + ColorReset)
				}
				fmt.Print(ColorYellow + "\n  [Tekan ENTER untuk kembali...]" + ColorReset)
				reader.ReadString('\n')
			}
			
		case 0:
			fmt.Println(ColorGreen + "\n  Terima kasih telah menggunakan sistem kami! " + ColorReset)
			fmt.Println("  Selamat belajar dan terus tingkatkan skill Anda!" + ColorReset)
			return
			
		default:
			fmt.Println(ColorRed + "\n  [!] Pilihan tidak tersedia. Masukkan 0-4.\n" + ColorReset)
		}
	}
}