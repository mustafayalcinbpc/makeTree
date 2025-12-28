# makeTree

`makeTree` is a lightweight (single-file) Go CLI application that generates a
directory–file tree similar to the tree command for the current or a given
directory and writes the output to a tree.md file.

## Features

- 📁 Root directory name is displayed at the top
- 📂 Directories are listed above files
- 📄 Regular files are listed under directories
- ⚙️ Dot files such as .env, .gitignore are included
- 🚫 .git, node_modules, vendor are automatically excluded
- 🔢 Depth limit support (--depth)
- 🧱 Single file (main.go), zero dependencies

## --help Output
```
Usage:
tree-md [path] [flags]

Arguments:
path                Directory path to scan.
If not provided, the current directory (.) is used.

Flags:
--depth int         Maximum depth of the directory tree (default: unlimited)
-h, --help          Show this help message

Examples:
tree-md
tree-md .
tree-md /path/to/project
tree-md ../another-project
tree-md /path/to/project --depth=2

Description:
tree-md generates a tree-like representation of the directory and file
structure of the given path and writes the result to a tree.md file
inside the target directory.

The following directories are excluded by default:
- .git
- node_modules
- vendor

Dot files (.env, .gitignore, etc.) are included.
```

## Example Output
```
makeTree/
├── cmd/
├── internal/
├── scripts/
├── LICENSE
├── README.md
├── main.go
├── .env
└── .gitignore
```


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
```


## Örnek Çıktı

```
makeTree/
├── cmd/
├── internal/
├── scripts/
├── LICENSE
├── README.md
├── main.go
├── .env
└── .gitignore
```


