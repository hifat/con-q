data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./cmd/migration",
  ]
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = "postgresql://postgres:postgres@localhost:5432/playgroud?sslmode=disable"
  migration {
    dir = "file://internal/app/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}