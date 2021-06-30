AR        := ar
CC        := gcc
INCLUDES  := -Isdk/C_SDK/crypt -Isdk/C_SDK/zlib -Isdk/C_SDK/
OPTIONS   := -Wno-macro-redefined -Wno-shift-negative-value -Os

all:  libcubesql.a
	rm -rf *.o
	
sdk:	
ifeq ($(wildcard sdk),)
	git clone https://github.com/cubesql/sdk.git
endif

zlib.o: sdk sdk/C_SDK/zlib/zlib.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/zlib/zlib.c -c -o zlib.o
    
aescrypt.o: sdk sdk/C_SDK/crypt/aescrypt.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/aescrypt.c -c -o aescrypt.o
	
aeskey.o: sdk sdk/C_SDK/crypt/aeskey.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/aeskey.c -c -o aeskey.o

aestab.o: sdk sdk/C_SDK/crypt/aestab.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/aestab.c -c -o aestab.o

base64.o: sdk sdk/C_SDK/crypt/base64.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/base64.c -c -o base64.o

pseudorandom.o: sdk sdk/C_SDK/crypt/pseudorandom.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/pseudorandom.c -c -o pseudorandom.o
  
sha1.o: sdk sdk/C_SDK/crypt/sha1.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/crypt/sha1.c -c -o sha1.o

cubesql.o: sdk sdk/C_SDK/cubesql.c
	$(CC) $(OPTIONS) $(INCLUDES) sdk/C_SDK/cubesql.c -c -o cubesql.o

libcubesql.a: zlib.o aescrypt.o aeskey.o aestab.o base64.o pseudorandom.o sha1.o cubesql.o
	$(AR) rcs libcubesql.a *.o

libcubesql.so: zlib.o aescrypt.o aeskey.o aestab.o base64.o pseudorandom.o sha1.o cubesql.o
	$(CC) -fPIC -shared *.o -o libcubesql.so

clean:
	rm -rf *.o *.a *.so
