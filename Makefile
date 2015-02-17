parser: y.tab.c lex.yy.c 
	gcc y.tab.c lex.yy.c -o parser -lfl -w

y.tab.c: 
	bison -y -d -g -t --verbose yacc.y

lex.yy.c: 
	lex lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output parser y.vcg test.py
	rm -rf tests

testgh: parser
	wget https://raw.githubusercontent.com/mhoc/cs352-test-cases/master/run.sh
	chmod +x run.sh
	./run.sh
	@$(MAKE) clean

testlocal: parser
	@cp ~/src/cs352-test-cases/test.py .
	@cp -R ~/src/cs352-test-cases/tests/ tests/
	python test.py
	@$(MAKE) clean
