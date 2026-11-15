@echo off

:: go run .

:: go CGO_CFLAGS="-Iexternal/Raven/include" CGO_LDFLAGS="-Lexternal/Raven/build/Release" go run .
go run . CGO_CFLAGS="-Iexternal/Raven/include" CGO_LDFLAGS="-Lexternal/Raven/build/Release"
