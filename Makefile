# Variáveis
BINARY_NAME=diino
MAIN_PATH=./cmd/server/main.go

# Comandos principais
.PHONY: help
help: ## Mostra os comandos disponíveis
	@echo "Comandos disponíveis:"
	@echo "  make install    - Instala as dependências"
	@echo "  make dev        - Inicia o servidor com hot reload"
	@echo "  make build      - Compila a aplicação"
	@echo "  make run        - Executa a aplicação"
	@echo "  make css        - Compila o Tailwind CSS"
	@echo "  make css-watch  - Compila o Tailwind CSS com watch"
	@echo "  make templ      - Gera os arquivos Templ"
	@echo "  make clean      - Limpa arquivos gerados"

.PHONY: install
install: ## Instala todas as dependências
	@echo "Instalando dependências Go..."
	go mod download
	go mod tidy
	@echo "Instalando Air..."
	go install github.com/air-verse/air@latest
	@echo "Instalando Templ..."
	go install github.com/a-h/templ/cmd/templ@latest
	@echo "Instalando Tailwind (necessário Node.js)..."
	npm install -D tailwindcss
	@echo "Dependências instaladas com sucesso!"

.PHONY: templ
templ: ## Gera os arquivos Templ
	@echo "Gerando arquivos Templ..."
	templ generate

.PHONY: css
css: ## Compila o Tailwind CSS
	@echo "Compilando Tailwind CSS..."
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

.PHONY: css-watch
css-watch: ## Compila o Tailwind CSS com watch mode
	@echo "Iniciando Tailwind CSS em modo watch..."
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

.PHONY: build
build: templ css ## Compila a aplicação
	@echo "Compilando aplicação..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build concluído: $(BINARY_NAME)"

.PHONY: run
run: build ## Executa a aplicação
	@echo "Executando aplicação..."
	./$(BINARY_NAME)

.PHONY: dev
dev: ## Inicia o servidor com hot reload
	@echo "Iniciando ambiente de desenvolvimento..."
	@echo "DICA: Execute 'make css-watch' em outro terminal para Tailwind hot reload"
	@make css &
	@air

.PHONY: dev-full
dev-full: ## Inicia dev com Tailwind watch em background
	@echo "Iniciando ambiente de desenvolvimento completo..."
	@make css-watch & air

.PHONY: clean
clean: ## Limpa arquivos gerados
	@echo "Limpando arquivos gerados..."
	@rm -f $(BINARY_NAME)
	@rm -rf tmp
	@echo "Limpeza concluída!"

.DEFAULT_GOAL := help
