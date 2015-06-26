deb:
	GOOS=linux GOARCH=amd64 go build -o build/stash-hook-router
	fpm -f -n stash-hook-router -s dir -t deb \
	    --prefix /usr/bin \
	    --version `git describe --tags --long` \
	    --license "BSD License" \
	    --description "Stash hook router" \
	    --url https://github.com/jhaals/stash-hook-router \
	    -C build \
	    .
	rm -r build
