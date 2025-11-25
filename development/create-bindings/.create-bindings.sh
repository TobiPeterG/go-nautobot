#!/bin/bash
set -euxo pipefail

echo "Creating GO bindings"

VERSION_FILE="/client/api/nautobot_version"
CURRENT_VERSION=$(head -n 1 $VERSION_FILE)
CURRENT_MAJOR_MINOR_VER=${CURRENT_VERSION%.*}

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}

# Remove generated files
for F in $(cat /client/.openapi-generator/FILES) ; do
    rm -f /client/"${F}"
done

cp /client/api/openapi-original.yaml /client/api/openapi.yaml


if [ "$CURRENT_MAJOR_MINOR_VER" = "$MAJOR_MINOR_VER" ]; then
    # Get the Patch version string
    NEW_PATCH_VERSION=$(echo $CURRENT_VERSION | awk -F '.' '{ print $3;}')
    # Remove suffixes in version
    NEW_PATCH_VERSION=${NEW_PATCH_VERSION%-*}
    # Increment Patch version string with 1
    NEW_PATCH_VERSION=$((${NEW_PATCH_VERSION} + 1))
    NEW_TAG=${CURRENT_MAJOR_MINOR_VER}.$NEW_PATCH_VERSION
else
    NEW_TAG=${MAJOR_MINOR_VER}.0
fi

# TODO: remove beta when it's in production
FINAL_NEW_TAG=${NEW_TAG}-beta

echo $FINAL_NEW_TAG > /client/api/nautobot_version

#Fix openapi spec file
/client/development/create-bindings/scripts/fix-spec.py

#yaml file is too long
export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
openapi-generator-cli generate --config /client/development/create-bindings/oapi-config.yaml \
    --input-spec /client/api/openapi.yaml \
    --output /client/ \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --http-user-agent go-nautobot/$(cat /client/api/nautobot_version)

rm /client/.travis.yml
/client/development/create-bindings/scripts/add-missing-imports.sh

echo "Copying READMEs"
mv /client/README.md /client/docs/README.md
sed -i 's|docs/||g' /client/docs/README.md
echo "docs/README.md" >> /client/.openapi-generator/FILES
cp /client/.README.md /client/README.md

echo "Starting Nautobot client tests..."

#export NAUTOBOT_URL=http://nautobot:8080/api/
#export NAUTOBOT_TOKEN=0123456789abcdef0123456789abcdef01234567

cd /client
go mod tidy
go test -v -gcflags="-e" ./...


echo "Nautobot client tests completed"
