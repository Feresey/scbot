data "template_dir" "migrations" {
  path = "migrations"
  vars = {}
}

// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://migrations/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = data.template_dir.migrations.url
  }
}

env "dev" {
  // ... a different env
}