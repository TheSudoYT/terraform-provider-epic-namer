terraform {
    required_providers {
        epic-namer = {
            version = "0.1.0"
            source = "localhost/providers/epic-namer"
        }
    }
}

provider "epic-namer" {}

resource "epic_random_name" "example" {}

output "random_name" {
    value = epic_random_name.example.name
}