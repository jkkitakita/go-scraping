# go-scraping

## Requirement

- golang (=> 1.12.5)

## Getting Started

1. Copy `category.sample.yaml` to `category.yaml`
```bash
cp category.sample.yaml category.yaml
```
2. Edit `category.yaml`
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
