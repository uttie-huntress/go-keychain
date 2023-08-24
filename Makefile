all:

cscope:
	@echo "Creating cscope database"
	-@rm -f ./cscope.files ./cscope.out
	@find . -iname "*go" > cscope.files
	@cscope -q -b -k -i ./cscope.files

