FROM golang:1.11-stretch as builder
ARG PROJECTNAME=github.com/pollosp/terraform_plugin_poc2
WORKDIR /go/src/${PROJECTNAME}
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -v && go build -o terraform/plugins/terraform-provider-artifact


FROM alpine as release
ARG PROJECTNAME=github.com/pollosp/terraform_plugin_poc2
COPY --from=hashicorp/terraform:0.11.11 /bin/terraform /bin/terraform
RUN mkdir -p /tmp/terraform/plugins
COPY ./terraform/ /tmp/terraform/
COPY --from=builder /go/src/${PROJECTNAME}/terraform/plugins/terraform-provider-artifact /tmp/terraform/plugins/terraform-provider-artifact
