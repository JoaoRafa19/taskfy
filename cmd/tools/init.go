package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initStore = &cobra.Command{
	Use:     `init [opções]`,
	Aliases: []string{"i"},
	Short:   "Inicializa uma nova estrutura para o tern e sqlc ",
	Long:    `Inicializa a estrutura de pastas para a store do projeto, cria os arquivos do tern e do sqlc, a pasta de queries e migrations`,
	Run: func(cmd *cobra.Command, args []string) {

		dir := os.DirFS(".")

		_, err := dir.Open("go.mod")
		if err != nil {
			fmt.Println("Erro ao inicializar\nComando deve ser executado na raiz do projeto")
			os.Exit(1)
		}
		err = os.MkdirAll(migrationsPath, os.ModePerm)
		if err != nil {
			fmt.Println("erro ao criar migrations")
			os.Exit(1)
		}
		err = os.MkdirAll(queriesPath, os.ModePerm)
		if err != nil {
			fmt.Println("erro ao criar queries")
			os.Exit(1)
		}

		f, err := os.Create(configFile)
		if err != nil {
			fmt.Println("erro ao criar configurações")
			os.Exit(1)
		}
		n, err := f.WriteString(`[database]
port = {{env "DATABASE_PORT"}}
database = {{env "DATABASE_NAME"}}
user = {{env "DATABASE_USER"}}
password = {{env "DATABASE_PASSWORD"}}
host = {{env "DATABASE_HOST"}}
`)

		if err != nil || n == 0 {
			fmt.Println("erro ao escrever configuralções")
			os.Exit(1)
		}

		sf, err := os.Create(storePath + "/sqlc.yaml")
		if err != nil {
			fmt.Println("erro ao criar configurações")
			os.Exit(1)
		}
		n, err = sf.WriteString(`version: 2
sql:
  - engine: postgresql
    queries: "./queries"
    schema: "./migrations"
    gen:
      go:
        emit_json_tags: true
        out: "."
        package: "pgstore"
        sql_package: "pgx/v5"
			`)

		if err != nil || n == 0 {
			fmt.Println("erro ao escrever configuralções sqlc")
			os.Exit(1)
		}

		mf, err := os.Create(migrationsPath + "/001_exemple.sql")
		if err != nil {
			fmt.Println("erro ao criar arquivo exemplo")
			os.Exit(1)
		}
		nmf, err := mf.WriteString(`-- Write your migrate up statements here

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
`)

		if err != nil || nmf == 0 {
			fmt.Println("erro ao escrever configuralções")
			os.Exit(1)
		}

		fmt.Println("Inicializado com sucesso !")
	},
}
