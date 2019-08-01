read -p "Arguments to pass: "  args

cd ./cmd/snip && go run main.go "$args"
