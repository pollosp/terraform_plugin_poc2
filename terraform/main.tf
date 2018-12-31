data "artifact" "test1" {
}

output "artifacttype" {
  value = "${data.artifact.test1.artifacttype}"
}
output "artifactlist" {
  value = "${data.artifact.test1.artifactlist}"
}

