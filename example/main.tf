terraform {
  required_providers {
    importtest = {
      versions = ["0.1.0"]
      source   = "hashicorp.dev/paddy/importtest"
    }
  }
}

provider "importtest" {
}

resource "importtest_resource" "test" {
  name = "importtest"
  sample_attribute = {
    "test" : "foo",
  }
}
