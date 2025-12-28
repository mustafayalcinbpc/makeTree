# makeTree

`makeTree`, bulunduğunuz ya da verdiğiniz bir dizinin **klasör–dosya ağacını**
`tree` komutuna benzer şekilde çıkaran ve sonucu **tree.md** dosyasına yazan
hafif (single-file) bir Go CLI uygulamasıdır.

## Özellikler

- 📁 Root dizin adı en üstte gösterilir
- 📂 Klasörler, dosyaların üstünde listelenir
- 📄 Normal dosyalar klasörlerin altında yer alır
- ⚙️ `.env`, `.gitignore` gibi **dot dosyalar dahil edilir**
- 🚫 `.git`, `node_modules`, `vendor` otomatik hariç tutulur
- 🔢 Derinlik sınırı (`--depth`)
- 🧱 Tek dosya (`main.go`), zero dependency

---

## --help Çıktısı

```text
Usage:
  tree-md [path] [flags]

Arguments:
  path                Taranacak dizin yolu.
                      Verilmezse mevcut dizin (.) kullanılır.

Flags:
  --depth int         Dizin ağacının maksimum derinliği (varsayılan: sınırsız)
  -h, --help          Bu yardım mesajını gösterir

Examples:
  tree-md
  tree-md .
  tree-md /path/to/project
  tree-md ../another-project
  tree-md /path/to/project --depth=2

Description:
  tree-md, verilen dizinin klasör ve dosya yapısını tree benzeri bir formatta
  çıkarır ve sonucu ilgili dizinin içine tree.md dosyası olarak yazar.

  Varsayılan olarak aşağıdaki dizinler hariç tutulur:
    - .git
    - node_modules
    - vendor

  Dot dosyalar (.env, .gitignore vb.) listeye dahildir.



## Örnek Çıktı

```text
makeTree/
├── cmd/
├── internal/
├── scripts/
├── LICENSE
├── README.md
├── main.go
├── .env
└── .gitignore


