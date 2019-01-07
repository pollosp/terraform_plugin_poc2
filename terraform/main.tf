provider "google" {
  project     = "XXXX"
}


data "artifact" "test1" {
}

output "artifacttype" {
  value = "${data.artifact.test1.artifacttype}"
}

output "artifactlist" {
  value = "${data.artifact.test1.artifactlist}"
}


locals {
  artifactlist = "${data.artifact.test1.artifactlist}"
}


resource "google_storage_bucket" "from_list" {
  count    = "${length(local.artifactlist)}"
  name     = "${element(local.artifactlist, count.index)}"
  location = "EU"
}
