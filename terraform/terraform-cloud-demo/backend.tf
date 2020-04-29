terraform {
  backend "terraform_cloud" {
    organization = "mfykmn-example"

    workspaces {
      name = "terraform-cloud-demo"
    }
  }
}
