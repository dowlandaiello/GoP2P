# GoP2P

A streamlined shard-based P2P networking stack built in Go.

[![Author](https://img.shields.io/badge/made%20by-Dowland%20Aiello-purple.svg?style=flat-round)](https://github.com/dowlandaiello)
[![Project](https://img.shields.io/badge/project-GoP2P-blue.svg?style=flat-round)](https://github.com/dowlandaiello/gop2p)
[![GoDoc](https://godoc.org/github.com/dowlandaiello/gop2p?status.svg)](https://godoc.org/github.com/dowlandaiello/gop2p)
[![Author](https://godoc.org/github.com/dowlandaiello/gop2p?status.svg)](https://godoc.org/github.com/dowlandaiello/gop2p)
[![Build Status](https://travis-ci.com/dowlandaiello/GoP2P.svg?branch=master)](https://travis-ci.com/dowlandaiello/GoP2P)
[![CircleCI](https://circleci.com/gh/dowlandaiello/GoP2P.svg?style=svg)](https://circleci.com/gh/dowlandaiello/GoP2P)
[![Go Report Card](https://goreportcard.com/badge/github.com/dowlandaiello/gop2p)](https://goreportcard.com/report/github.com/dowlandaiello/gop2p)
[![codecov](https://codecov.io/gh/dowlandaiello/GoP2P/branch/master/graph/badge.svg)](https://codecov.io/gh/dowlandaiello/GoP2P)

## Dependencies

All critical dependencies have been included in .vendor/. As of now, protoc (compile .proto files) is the only GoP2P dependency not shipped in .vendor or .vendor-new. Fortunately, protoc is only necessary for the proper use of GoP2P in the case that .pb.go files need to be compiled on runtime. In production environments where users cannot be prompted to install protoc, it is recommended that the necessary .pb.go files are compiled and provided by the developers of software using GoP2P before runtime.
