#!/usr/bin/env bash
set -euox pipefail

read -r -p "Lower case provider name (ex. github): " PROVIDER_NAME_LOWER
read -r -p "Normal case provider name (ex. GitHub): " PROVIDER_NAME_NORMAL
read -r -p "Organization (ex. devops-chris, my-org-name): " ORGANIZATION_NAME

REPLACE_FILES='./* ./.github :!build/** :!go.* :!hack/prepare.sh'
# shellcheck disable=SC2086
git grep -l 'twingate' -- ${REPLACE_FILES} | xargs sed -i.bak "s/upjet-provider-twingate/provider-${PROVIDER_NAME_LOWER}/g"
# shellcheck disable=SC2086
git grep -l 'twingate' -- ${REPLACE_FILES} | xargs sed -i.bak "s/twingate/${PROVIDER_NAME_LOWER}/g"
# shellcheck disable=SC2086
git grep -l "devops-chris/provider-${PROVIDER_NAME_LOWER}" -- ${REPLACE_FILES} | xargs sed -i.bak "s|devops-chris/provider-${PROVIDER_NAME_LOWER}|${ORGANIZATION_NAME}/provider-${PROVIDER_NAME_LOWER}|g"
# shellcheck disable=SC2086
git grep -l 'Twingate' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Twingate/${PROVIDER_NAME_NORMAL}/g"
# We need to be careful while replacing "twingate" keyword in go.mod as it could tamper
# some imported packages under require section.
sed -i.bak "s|devops-chris/upjet-provider-twingate|${ORGANIZATION_NAME}/provider-${PROVIDER_NAME_LOWER}|g" go.mod
sed -i.bak "s|PROJECT_REPO := github.com/devops-chris/|PROJECT_REPO := github.com/${ORGANIZATION_NAME}/|g" Makefile

# Clean up the .bak files created by sed
git clean -fd

git mv "internal/clients/twingate.go" "internal/clients/${PROVIDER_NAME_LOWER}.go"
git mv "cluster/images/upjet-provider-twingate" "cluster/images/provider-${PROVIDER_NAME_LOWER}"

# We need to remove this api folder otherwise first `make generate` fails with
# the following error probably due to some optimizations in go generate with v1.17:
# generate: open /Users/hasanturken/Workspace/crossplane-contrib/upjet-provider-twingate/apis/null/v1alpha1/zz_generated.deepcopy.go: no such file or directory
rm -rf apis/null