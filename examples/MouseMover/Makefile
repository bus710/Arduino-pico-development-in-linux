clean: 
	rm -rf ./build

# Build compiles/verifies/uploads the sketch.
build:
	arduino-cli compile \
		-b rp2040:rp2040:rpipicow \
		--output-dir ./build \
		--board-options ipbtstack=ipv4btcble \
		--upload \
		--verify \
		-p /dev/ttyACM0 \
		-l serial \

