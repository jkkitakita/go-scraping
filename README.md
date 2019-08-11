# go-scraping

## Requirement

- golang (=> 1.12.5)

## Getting Started

1. Copy `./config/config.sample.yaml` to `./config/config.yaml`
```bash
cp ./config/config.sample.yaml ./config/config.yaml
```
2. Edit `./config/config.yaml`
3. Setup golang
```
make deps
```
4. Let's scraping!
```
make run

or

make build && ./go-scraping
```
