#!/bin/bash

SCRIPT="$0"
echo "# START SCRIPT: $SCRIPT"

EXECUTABLE='/d/AppData/bin/openapi-generator-cli.jar'
PACKAGE_NAME="blackapi"
SPEC="spec.yaml"
OUT_DIR="$PACKAGE_NAME"
GIT_USER_ID=aibotsoft
GIT_REPO_ID=blackapi

#echo "Removing files and folders under $STUB_DIR"
#rm -rf $STUB_DIR/{*.go,*.sh,*.md,*.mod,*.sum}

java -jar $EXECUTABLE generate -i $SPEC -g go -o $OUT_DIR --package-name $PACKAGE_NAME --type-mappings integer=int64\
 --git-user-id $GIT_USER_ID --git-repo-id $GIT_REPO_ID\
  --additional-properties=disallowAdditionalPropertiesIfNotPresent=true,isGoSubmodule=true,packageVersion=2.0.0

# --additional-properties --generate-alias-as-model true
#echo "Removing mod files"
#rm -rf $STUB_DIR/api
#rm -rf $STUB_DIR/go.mod
#rm -rf $STUB_DIR/go.sum
#rm -rf $STUB_DIR/git_push.sh
#rm -rf $STUB_DIR/readme.md
#rm -rf $STUB_DIR/.travis.yml
#rm -rf $STUB_DIR/.gitignore
#rm -rf $STUB_DIR/docs
#rm -rf $STUB_DIR/.openapi-generator
