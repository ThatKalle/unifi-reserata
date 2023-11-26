# UniFi-Reserata

Command-line tool to decrypt UniFi Controller backup `.unf` and support `.supp` files.

> Reserata - unlocked, open up, unseal. -- Latin

## 🔓 Use the application

Place the application in a directory containing your UniFi Controller backup `.unf` and support `.supp` files.
```
directory
├── network_backup_21.11.2023_15-57_v7.4.162.unf
├── network_support_20-11-2023.supp
└── unifireserata.exe
```

Run the application:
```powershell
PS directory> .\unifireserata.exe
1) network_backup_21.11.2023_15-57_v7.4.162.unf
2) network_support_20-11-2023.supp
Select a File
1
PS directory .\unifireserata.exe
1) network_backup_21.11.2023_15-57_v7.4.162.unf
2) network_support_20-11-2023.supp
Select a File
2
```

A `.zip` file with the same name as the file selected will be created in the same directory.\
This .zip file can be browsed using [7ip](https://www.7-zip.org/download.html).

## ⚡️ Quick start

Head over to the repository's [Release page](https://github.com/ThatKalle/unifi-reserata/releases), if you want to download a ready-made Windows or Linux packages build by [GoReleaser](https://goreleaser.com/).

### Build from Source

[Download](https://go.dev/dl/) and install **Go**. Version `1.21.4` or higher is required.

Install dependencies and tools:
```go
go mod download
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
```

Run the [`go generate`](https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source) command to build the `resource.syso` needed to set Windows `.exe` versioninfo using [goversioninfo](github.com/josephspurrier/goversioninfo).

Build the package using the [`go build`](https://pkg.go.dev/cmd/go#hdr-Build_and_test_caching) command:
```go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/unifireserata
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/unifireserata.exe
```

There is also a makefile:
```bash
make build # if on Linux
make buildwin # if on Windows
```
