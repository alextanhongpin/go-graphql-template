#!/bin/bash

if [ -z "$GOPATH" ]; then
	echo "ERROR: pre-commit hook: \$GOPATH is empty"
	exit 1
fi

if [ -z "$(which golint)" ]; then
	echo "ERROR: golint not found, please run: go get -u golang.org/x/lint/golint"
	exit 1
fi

if [ -z "$(which errcheck)" ]; then
	echo "ERROR: errcheck not found, please run: go get -u github.com/kisielk/errcheck
"
fi

# diff-filter: ACM (Add, Copied, Modified)
STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

if [ -z "$STAGED_GO_FILES" ]; then
	echo "No staged go files"
	exit 0
fi

# Find all unique directories - since the tests won't run per file.
#PKGS=$(find $STAGED_GO_FILES -exec dirname "./{}" \; | sort | uniq)

# Check if errors are checked.
errcheck -ignoretests -ignoregenerated   ./...
if [[ $? -ne 0 ]]; then
	echo "ERROR: errcheck failed"
	exit 1
fi

errors=
for file in $STAGED_GO_FILES
do
	go fmt $file
	if [ $? -ne 0 ]; then
		errors=YES
	fi
done

if [ ! -z "$errors" ]; then
	echo "ERROR: go fmt failed"
	exit 1
fi

# Go vet does not work on partial imports.
go vet ./...

errors=
for file in $PKGS
do
	golint -set_exit_status $file
	if [ $? -ne 0 ]; then
		errors=YES
	fi
done

if [ ! -z "$errors" ]; then
	echo "ERROR: golint failed"
	exit 1
fi

STAGED_GO_TEST_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\_test.go$')
if [ -z "$STAGED_GO_TEST_FILES" ]; then
	echo "No staged go files"
	exit 0
fi

# Find all unique directories - since the tests won't run per file.
# NOTE: We append ./ to make the path relative to root dir.
TESTS=$(find $STAGED_GO_TEST_FILES -exec dirname "./{}" \; | sort | uniq)

# Named files must all be in one directory;
for pkg in $TESTS
do
	go test -v -failfast $pkg
	if [ $? -ne 0 ]; then
		exit 1
	fi
done
